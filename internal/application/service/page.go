package application

import (
	"context"

	"github.com/google/uuid"
	me "github.com/octoposprime/op-be-book/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-book/internal/domain/model/object"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// GetPagesByFilter returns the pages that match the given filter.
func (a *Service) GetPagesByFilter(ctx context.Context, pageFilter me.PageFilter) (me.Pages, error) {
	return a.DbPort.GetPagesByFilter(ctx, pageFilter)
}

// CreatePage sends the given page to the repository of the infrastructure layer for creating a new page.
func (a *Service) CreatePage(ctx context.Context, page me.Page) (me.Page, error) {
	page.Id = uuid.UUID{}
	if err := a.ValidatePage(&page); err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreatePage", userId, err.Error()))
		return me.Page{}, err
	}
	if page.PageStatus == mo.PageStatusNONE {
		page.PageStatus = mo.PageStatusACTIVE
	}
	return a.DbPort.SavePage(ctx, page)
}

// UpdatePageBase sends the given base values of the page to the repository of the infrastructure layer for updating base values of page data.
func (a *Service) UpdatePageBase(ctx context.Context, page me.Page) (me.Page, error) {
	if user.Id.String() == "" || user.Id == (uuid.UUID{}) {
		err := mo.ErrorUserIdIsEmpty
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	var userFilter me.UserFilter
	userFilter.Id = user.Id
	users, err := a.GetUsersByFilter(ctx, userFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	if users.TotalRows > 0 {
		dbUser := users.Users[0]
		dbUser.Tags = user.Tags
		dbUser.PageType = user.PageType
		if err := a.ValidateUser(&dbUser); err != nil {
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
			return me.User{}, err
		}
		return a.DbPort.SaveUser(ctx, dbUser)
	} else {
		return user, mo.ErrorUserNotFound
	}
}

// UpdatePageCore sends the given core values of the page to the repository of the infrastructure layer for updating core values of page data.
func (a *Service) UpdatePageCore(ctx context.Context, page me.Page) (me.Page, error) {
	if user.Id.String() == "" || user.Id == (uuid.UUID{}) {
		err := mo.ErrorUserIdIsEmpty
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	var userFilter me.UserFilter
	userFilter.Id = user.Id
	users, err := a.GetUsersByFilter(ctx, userFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	if users.TotalRows > 0 {
		dbUser := users.Users[0]
		dbUser.PageData = user.PageData
		if err := a.ValidateUser(&dbUser); err != nil {
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
			return me.User{}, err
		}
		return a.DbPort.SaveUser(ctx, dbUser)
	} else {
		return user, mo.ErrorUserNotFound
	}
}

// UpdatePageStatus sends the given status value of the page to the repository of the infrastructure layer for updating status of page data.
func (a *Service) UpdatePageStatus(ctx context.Context, page me.Page) (me.Page, error) {
	if user.Id.String() == "" || user.Id == (uuid.UUID{}) {
		err := mo.ErrorUserIdIsEmpty
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	var userFilter me.UserFilter
	userFilter.Id = user.Id
	users, err := a.GetUsersByFilter(ctx, userFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	if users.TotalRows > 0 {
		dbUser := users.Users[0]
		dbUser.UserStatus = user.UserStatus
		if err := a.ValidateUser(&dbUser); err != nil {
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
			return me.User{}, err
		}
		return a.DbPort.SaveUser(ctx, dbUser)
	} else {
		return user, mo.ErrorUserNotFound
	}
}

// DeletePage sends the given page to the repository of the infrastructure layer for deleting data.
func (a *Service) DeletePage(ctx context.Context, page me.Page) (me.Page, error) {
	var err error
	page, err = a.DbPort.DeletePage(ctx, page)
	if err != nil {
		pageId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeletePage", pageId, err.Error()))
		return me.Page{}, err
	}

	return page, err
}
