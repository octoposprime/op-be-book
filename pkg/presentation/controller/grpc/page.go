package presentation

import (
	"context"

	dto "github.com/octoposprime/op-be-book/pkg/presentation/dto"
	pb_page "github.com/octoposprime/op-be-shared/pkg/proto/pb/page"
)

// GetPagesByFilter returns the pages that match the given filter.
func (a *Grpc) GetPagesByFilter(ctx context.Context, filter *pb_page.PageFilter) (*pb_page.Pages, error) {
	pages, err := a.queryHandler.GetPagesByFilter(ctx, *dto.NewPageFilter(filter).ToEntity())
	return dto.NewPageFromEntities(pages).ToPbs(), err
}

// CreatePage sends the given page to the application layer for creating new page.
func (a *Grpc) CreatePage(ctx context.Context, page *pb_page.Page) (*pb_page.Page, error) {
	data, err := a.commandHandler.CreatePage(ctx, *dto.NewPage(page).ToEntity())
	return dto.NewPageFromEntity(data).ToPb(), err
}

// UpdatePageBase sends the given page to the application layer for updating page's base values.
func (a *Grpc) UpdatePageBase(ctx context.Context, page *pb_page.Page) (*pb_page.Page, error) {
	data, err := a.commandHandler.UpdatePageBase(ctx, *dto.NewPage(page).ToEntity())
	return dto.NewPageFromEntity(data).ToPb(), err
}

// UpdatePageCore sends the given page to the application layer for updating page's core values.
func (a *Grpc) UpdatePageCore(ctx context.Context, page *pb_page.Page) (*pb_page.Page, error) {
	data, err := a.commandHandler.UpdatePageCore(ctx, *dto.NewPage(page).ToEntity())
	return dto.NewPageFromEntity(data).ToPb(), err
}

// UpdatePageStatus sends the given page to the application layer for updating page status.
func (a *Grpc) UpdatePageStatus(ctx context.Context, page *pb_page.Page) (*pb_page.Page, error) {
	data, err := a.commandHandler.UpdatePageStatus(ctx, *dto.NewPage(page).ToEntity())
	return dto.NewPageFromEntity(data).ToPb(), err
}

// DeletePage sends the given page to the application layer for deleting data.
func (a *Grpc) DeletePage(ctx context.Context, page *pb_page.Page) (*pb_page.Page, error) {
	data, err := a.commandHandler.DeletePage(ctx, *dto.NewPage(page).ToEntity())
	return dto.NewPageFromEntity(data).ToPb(), err
}
