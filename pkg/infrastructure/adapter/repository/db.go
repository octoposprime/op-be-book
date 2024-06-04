package infrastructure

import (
	"context"

	"github.com/google/uuid"
	me "github.com/octoposprime/op-be-book/internal/domain/model/entity"
	map_repo "github.com/octoposprime/op-be-book/pkg/infrastructure/mapper/repository"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tgorm "github.com/octoposprime/op-be-shared/tool/gorm"
)

type DbAdapter struct {
	*tgorm.GormClient
	Log func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)
}

func NewDbAdapter(dbClient *tgorm.GormClient) DbAdapter {
	adapter := DbAdapter{
		dbClient,
		Log,
	}

	err := dbClient.DbClient.AutoMigrate(&map_repo.Page{})
	if err != nil {
		panic(err)
	}

	return adapter
}

// SetLogger sets logging call-back function
func (a *DbAdapter) SetLogger(LoggerFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)) {
	a.Log = LoggerFunc
}

// GetPagesByFilter returns the pages that match the given filter.
func (a DbAdapter) GetPagesByFilter(ctx context.Context, pageFilter me.PageFilter) (me.Pages, error) {
	var pagesDbMapper map_repo.Pages
	var filter map_repo.Page
	qry := a.DbClient
	if pageFilter.Id.String() != "" && pageFilter.Id != (uuid.UUID{}) {
		filter.ID = pageFilter.Id
	}
	if pageFilter.PageData != "" {
		filter.PageData = pageFilter.PageData
	}
	if pageFilter.PageType != 0 {
		filter.PageType = int(pageFilter.PageType)
	}
	if pageFilter.PageStatus != 0 {
		filter.PageStatus = int(pageFilter.PageStatus)
	}
	if len(pageFilter.Tags) > 0 {
		filter.Tags = pageFilter.Tags
	}
	if !pageFilter.CreatedAtFrom.IsZero() && !pageFilter.CreatedAtTo.IsZero() {
		qry = qry.Where("created_at between ? and ?", pageFilter.CreatedAtFrom, pageFilter.CreatedAtTo)
	}
	if !pageFilter.UpdatedAtFrom.IsZero() && !pageFilter.UpdatedAtTo.IsZero() {
		qry = qry.Where("updated_at between ? and ?", pageFilter.UpdatedAtFrom, pageFilter.UpdatedAtTo)
	}
	if pageFilter.SearchText != "" {
		qry = qry.Where(
			qry.Where("UPPER(page_name) LIKE UPPER(?)", "%"+pageFilter.SearchText+"%").
				Or("UPPER(email) LIKE UPPER(?)", "%"+pageFilter.SearchText+"%").
				Or("UPPER(array_to_string(tags, ',')) LIKE UPPER(?)", "%"+pageFilter.SearchText+"%"),
		)
	}
	qry = qry.Where(filter)
	var totalRows int64
	result := qry.Model(&map_repo.Page{}).Where(filter).Count(&totalRows)
	if result.Error != nil {
		pageId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetPagesByFilter", pageId, result.Error.Error()))
		totalRows = 0
	}
	if pageFilter.Limit != 0 {
		qry = qry.Limit(pageFilter.Limit)
	}
	if pageFilter.Offset != 0 {
		qry = qry.Offset(pageFilter.Offset)
	}
	if pageFilter.SortType != "" && pageFilter.SortField != 0 {
		sortStr := map_repo.PageSortMap[pageFilter.SortField]
		if pageFilter.SortType == "desc" || pageFilter.SortType == "DESC" {
			sortStr += " desc"
		} else {
			sortStr += " asc"
		}
		qry = qry.Order(sortStr)
	}
	result = qry.Find(&pagesDbMapper)
	if result.Error != nil {
		pageId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetPagesByFilter", pageId, result.Error.Error()))
		return me.Pages{}, result.Error
	}
	return me.Pages{
		Pages:     pagesDbMapper.ToEntities(),
		TotalRows: totalRows,
	}, nil
}

// SavePage insert a new page or update the existing one in the database.
func (a DbAdapter) SavePage(ctx context.Context, page me.Page) (me.Page, error) {
	pageDbMapper := map_repo.NewPageFromEntity(page)
	qry := a.DbClient
	if page.Id.String() != "" && page.Id != (uuid.UUID{}) {
		qry = qry.Omit("created_at")
	}
	pageId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	if pageDbMapper.ID != (uuid.UUID{}) {
		pageDbMapper.UpdatedBy, _ = uuid.Parse(pageId)
	} else {
		pageDbMapper.CreatedBy, _ = uuid.Parse(pageId)
	}
	result := qry.Save(&pageDbMapper)
	if result.Error != nil {
		pageId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "SavePage", pageId, result.Error.Error()))
		return me.Page{}, result.Error
	}
	return *pageDbMapper.ToEntity(), nil
}

// DeletePage soft-deletes the given page in the database.
func (a DbAdapter) DeletePage(ctx context.Context, page me.Page) (me.Page, error) {
	pageDbMapper := map_repo.NewPageFromEntity(page)
	pageId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	pageDbMapper.DeletedBy, _ = uuid.Parse(pageId)
	result := a.DbClient.Delete(&pageDbMapper)
	if result.Error != nil {
		pageId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeletePage", pageId, result.Error.Error()))
		return me.Page{}, result.Error
	}
	return *pageDbMapper.ToEntity(), nil
}
