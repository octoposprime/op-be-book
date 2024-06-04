package application

import (
	"context"

	me "github.com/octoposprime/op-be-book/internal/domain/model/entity"
)

// CreatePage sends the given page to the application layer for creating a new page.
func (a CommandAdapter) CreatePage(ctx context.Context, page me.Page) (me.Page, error) {
	return a.Service.CreatePage(ctx, page)
}

// UpdatePageBase sends the given base values of the page to the repository of the infrastructure layer for updating base values of page data.
func (a CommandAdapter) UpdatePageBase(ctx context.Context, page me.Page) (me.Page, error) {
	return a.Service.UpdatePageBase(ctx, page)
}

// UpdatePageCore sends the given core values of the page to the repository of the infrastructure layer for updating core values of page data.
func (a CommandAdapter) UpdatePageCore(ctx context.Context, page me.Page) (me.Page, error) {
	return a.Service.UpdatePageCore(ctx, page)
}

// UpdatePageStatus sends the given status value of the page to the repository of the infrastructure layer for updating status of page data.
func (a CommandAdapter) UpdatePageStatus(ctx context.Context, page me.Page) (me.Page, error) {
	return a.Service.UpdatePageStatus(ctx, page)
}

// DeletePage sends the given page to the application layer for deleting data.
func (a CommandAdapter) DeletePage(ctx context.Context, page me.Page) (me.Page, error) {
	return a.Service.DeletePage(ctx, page)
}
