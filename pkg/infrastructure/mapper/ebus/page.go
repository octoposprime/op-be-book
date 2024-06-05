package infrastructure

import (
	"fmt"

	me "github.com/octoposprime/op-be-book/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-book/internal/domain/model/object"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/page"
	tuuid "github.com/octoposprime/op-be-shared/tool/uuid"
)

// Page is a struct that represents the ebus mapper of a page basic values.
type Page struct {
	proto *pb.Page
}

// NewPage creates a new *Page.
func NewPage(pb *pb.Page) *Page {
	return &Page{
		proto: pb,
	}
}

// String returns a string representation of the Page.
func (s *Page) String() string {
	return fmt.Sprintf("Id: %v, "+
		"PageData: %v, "+
		"PageType: %v, "+
		"PageStatus: %v, "+
		"Tags: %v",
		s.proto.Id,
		s.proto.PageData,
		s.proto.PageType,
		s.proto.PageStatus,
		s.proto.Tags)
}

// NewPageFromEntity creates a new *Page from entity.
func NewPageFromEntity(entity me.Page) *Page {
	return &Page{
		&pb.Page{
			Id:         entity.Id.String(),
			PageData:   entity.PageData,
			PageType:   pb.PageType(entity.PageType),
			PageStatus: pb.PageStatus(entity.PageStatus),
			Tags:       entity.Tags,
		},
	}
}

// ToPb returns a protobuf representation of the Page.
func (s *Page) ToPb() *pb.Page {
	return s.proto
}

// ToEntity returns a entity representation of the Page.
func (s *Page) ToEntity() *me.Page {
	return &me.Page{
		Id: tuuid.FromString(s.proto.Id),
		Page: mo.Page{
			PageData:   s.proto.PageData,
			PageType:   mo.PageType(s.proto.PageType),
			PageStatus: mo.PageStatus(s.proto.PageStatus),
			Tags:       s.proto.Tags,
		},
	}
}

type Pages []*Page

// NewPagesFromEntities creates a new []*Page from entities.
func NewPageFromEntities(entities []me.Page) Pages {
	pages := make([]*Page, len(entities))
	for i, entity := range entities {
		pages[i] = NewPageFromEntity(entity)
	}
	return pages
}

// ToPbs returns a protobuf representation of the Pages.
func (s Pages) ToPbs() *pb.Pages {
	pages := make([]*pb.Page, len(s))
	for i, page := range s {
		pages[i] = page.ToPb()
	}
	return &pb.Pages{
		Pages: pages,
	}
}
