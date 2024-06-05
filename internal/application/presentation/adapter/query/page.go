package application

import (
	"context"

	me "github.com/octoposprime/op-be-book/internal/domain/model/entity"
)

// GetPagesByFilter returns the pages that match the given filter.
func (a QueryAdapter) GetPagesByFilter(ctx context.Context, pageFilter me.PageFilter) (me.Pages, error) {
	return a.Service.GetPagesByFilter(ctx, pageFilter)
}
