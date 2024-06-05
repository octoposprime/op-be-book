package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	mo "github.com/octoposprime/op-be-book/internal/domain/model/object"
)

// PageFilter is a struct that represents the filter of a page.
type PageFilter struct {
	Id         uuid.UUID     `json:"id"`          // Id is the id of the page.
	PageData   string        `json:"page_name"`   // PageData is the page name of the page.
	PageType   mo.PageType   `json:"page_type"`   // PageType is the type of the page.
	PageStatus mo.PageStatus `json:"page_status"` // PageStatus is the status of the page.
	Tags       []string      `json:"tags"`        // Tags is the tags of the page.

	CreatedAtFrom time.Time `json:"created_at_from"` // CreatedAt is in the between of CreatedAtFrom and CreatedAtTo.
	CreatedAtTo   time.Time `json:"created_at_to"`   // CreatedAt is in the between of CreatedAtFrom and CreatedAtTo.
	UpdatedAtFrom time.Time `json:"updated_at_from"` // UpdatedAt is in the between of UpdatedAtFrom and UpdatedAtTo.
	UpdatedAtTo   time.Time `json:"updated_at_to"`   // UpdatedAt is in the between of UpdatedAtFrom and UpdatedAtTo.

	SearchText string           `json:"search_text"` // SearchText is the full-text search value.
	SortType   string           `json:"sort_type"`   // SortType is the sorting type (ASC,DESC).
	SortField  mo.PageSortField `json:"sort_field"`  // SortField is the sorting field of the page.

	Limit  int `json:"limit"`  // Limit provides to limitation row size.
	Offset int `json:"offset"` // Offset provides a starting row number of the limitation.
}

// NewPageFilter creates a new *PageFilter.
func NewPageFilter(id uuid.UUID,
	pageData string,
	pageType mo.PageType,
	pageStatus mo.PageStatus,
	tags []string,
	createdAtFrom time.Time,
	createdAtTo time.Time,
	updatedAtFrom time.Time,
	updatedAtTo time.Time,
	searchText string,
	sortType string,
	sortField mo.PageSortField,
	limit int,
	offset int) *PageFilter {
	return &PageFilter{
		Id:            id,
		PageData:      pageData,
		PageType:      pageType,
		PageStatus:    pageStatus,
		Tags:          tags,
		CreatedAtFrom: createdAtFrom,
		CreatedAtTo:   createdAtTo,
		UpdatedAtFrom: updatedAtFrom,
		UpdatedAtTo:   updatedAtTo,
		SearchText:    searchText,
		SortType:      sortType,
		SortField:     sortField,
		Limit:         limit,
		Offset:        offset,
	}
}

// NewEmptyPageFilter creates a new *PageFilter with empty values.
func NewEmptyPageFilter() *PageFilter {
	return &PageFilter{
		Id:            uuid.UUID{},
		PageData:      "",
		PageType:      mo.PageTypeNONE,
		PageStatus:    mo.PageStatusNONE,
		Tags:          []string{},
		CreatedAtFrom: time.Time{},
		CreatedAtTo:   time.Time{},
		UpdatedAtFrom: time.Time{},
		UpdatedAtTo:   time.Time{},
		SearchText:    "",
		SortType:      "",
		SortField:     mo.PageSortFieldNONE,
		Limit:         0,
		Offset:        0,
	}
}

// String returns a string representation of the PageFilter.
func (s *PageFilter) String() string {
	return fmt.Sprintf("Id: %v, "+
		"PageData: %v, "+
		"PageType: %v, "+
		"PageStatus: %v, "+
		"Tags: %v, "+
		"CreatedAtFrom: %v, "+
		"CreatedAtTo: %v, "+
		"UpdatedAtFrom: %v, "+
		"UpdatedAtTo: %v, "+
		"SearchText: %v, "+
		"SortType: %v, "+
		"SortField: %v, "+
		"Limit: %v, "+
		"Offset: %v",
		s.Id,
		s.PageData,
		s.PageType,
		s.PageStatus,
		s.Tags,
		s.CreatedAtFrom,
		s.CreatedAtTo,
		s.UpdatedAtFrom,
		s.UpdatedAtTo,
		s.SearchText,
		s.SortType,
		s.SortField,
		s.Limit,
		s.Offset)
}
