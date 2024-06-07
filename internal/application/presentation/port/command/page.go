package application

import (
	"context"

	me "github.com/octoposprime/op-be-book/internal/domain/model/entity"
)

// CommandPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type PageCommandPort interface {
	// CreatePage sends the given page to the application layer for creating a new page.
	CreatePage(ctx context.Context, page me.Page) (me.Page, error)

	// UpdatePageBase sends the given base values of the page to the repository of the infrastructure layer for updating base values of page data.
	UpdatePageBase(ctx context.Context, page me.Page) (me.Page, error)

	// UpdatePageCore sends the given core values of the page to the repository of the infrastructure layer for updating core values of page data.
	UpdatePageCore(ctx context.Context, page me.Page) (me.Page, error)

	// UpdatePageStatus sends the given status value of the page to the repository of the infrastructure layer for updating status of page data.
	UpdatePageStatus(ctx context.Context, page me.Page) (me.Page, error)

	// DeletePage sends the given page to the application layer for deleting data.
	DeletePage(ctx context.Context, page me.Page) (me.Page, error)
}
