package application

import (
	"context"

	me "github.com/octoposprime/op-be-book/internal/domain/model/entity"
)

// QueryPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type PageQueryPort interface {
	// GetPagesByFilter returns the pages that match the given filter.
	GetPagesByFilter(ctx context.Context, pageFilter me.PageFilter) (me.Pages, error)
}
