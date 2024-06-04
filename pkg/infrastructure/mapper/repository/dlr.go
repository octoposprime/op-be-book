package infrastructure

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	me "github.com/octoposprime/op-be-book/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-book/internal/domain/model/object"
	tgorm "github.com/octoposprime/op-be-shared/tool/gorm"
)

// Page is a struct that represents the db mapper of a page basic values.
type Page struct {
	tgorm.Model

	PageData   string         `json:"page_data" gorm:"not null;default:''"`  // PageData is the page data of the page.
	PageType   int            `json:"page_type" gorm:"not null;default:0"`   // PageType is the type of the page.
	PageStatus int            `json:"page_status" gorm:"not null;default:0"` // PageStatus is the status of the page.
	Tags       pq.StringArray `json:"tags" gorm:"type:text[]"`               // Tags is the tags of the page.
}

// NewPage creates a new *Page.
func NewPage(id uuid.UUID,
	pageData string,
	pageType int,
	pageStatus int,
	tags pq.StringArray) *Page {
	return &Page{
		Model:      tgorm.Model{ID: id},
		PageData:   pageData,
		PageType:   pageType,
		PageStatus: pageStatus,
		Tags:       tags,
	}
}

// String returns a string representation of the Page.
func (s *Page) String() string {
	return fmt.Sprintf("Id: %v, "+
		"PageData: %v, "+
		"PageType: %v, "+
		"PageStatus: %v, "+
		"Tags: %v",
		s.ID,
		s.PageData,
		s.PageType,
		s.PageStatus,
		s.Tags)
}

// NewPageFromEntity creates a new *Page from entity.
func NewPageFromEntity(entity me.Page) *Page {
	return &Page{
		Model:      tgorm.Model{ID: entity.Id},
		PageData:   entity.PageData,
		PageType:   int(entity.PageType),
		PageStatus: int(entity.PageStatus),
		Tags:       entity.Tags,
	}
}

// ToEntity returns a entity representation of the Page.
func (s *Page) ToEntity() *me.Page {
	return &me.Page{
		Id: s.ID,
		Page: mo.Page{
			PageData:   s.PageData,
			PageType:   mo.PageType(s.PageType),
			PageStatus: mo.PageStatus(s.PageStatus),
			Tags:       s.Tags,
		},
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
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

// ToEntities creates a new []me.Page entity.
func (s Pages) ToEntities() []me.Page {
	pages := make([]me.Page, len(s))
	for i, page := range s {
		pages[i] = *page.ToEntity()
	}
	return pages
}
