package application

import (
	"context"

	me "github.com/octoposprime/op-be-book/internal/domain/model/entity"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// DbPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the database.
type DbPort interface {
	// SetLogger sets logging call-back function
	SetLogger(LogFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error))

	// GetPagesByFilter returns the pages that match the given filter.
	GetPagesByFilter(ctx context.Context, pageFilter me.PageFilter) (me.Pages, error)

	// SavePage insert a new page or update the existing one in the database.
	SavePage(ctx context.Context, page me.Page) (me.Page, error)

	// DeletePage soft-deletes the given page in the database.
	DeletePage(ctx context.Context, page me.Page) (me.Page, error)
}
