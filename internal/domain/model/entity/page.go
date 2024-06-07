package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	mo "github.com/octoposprime/op-be-book/internal/domain/model/object"
)

// Page is a struct that represents the entity of a page basic values.
type Page struct {
	Id      uuid.UUID `json:"id"` // Id is the id of the page.
	mo.Page           // Page is the basic values of the page.

	// Only for view
	CreatedAt time.Time `json:"created_at"` // CreatedAt is the create time.
	UpdatedAt time.Time `json:"updated_at"` // UpdatedAt is the update time.
}

// NewPage creates a new *Page.
func NewPage(id uuid.UUID,
	page mo.Page) *Page {
	return &Page{
		Id:   id,
		Page: page,
	}
}

// NewEmptyPage creates a new *Page with empty values.
func NewEmptyPage() *Page {
	return &Page{
		Id:   uuid.UUID{},
		Page: *mo.NewEmptyPage(),
	}
}

// String returns a string representation of the Page.
func (s *Page) String() string {
	return fmt.Sprintf("Id: %v, "+
		"Page: %v",
		s.Id,
		s.Page)
}

// Equals returns true if the Page is equal to the other Page.
func (s *Page) Equals(other *Page) bool {
	if s.Id != other.Id {
		return false
	}
	if !s.Page.Equals(&other.Page) {
		return false
	}
	return true
}

// Clone returns a clone of the Page.
func (s *Page) Clone() *Page {
	return &Page{
		Id:   s.Id,
		Page: *s.Page.Clone(),
	}
}

// IsEmpty returns true if the Page is empty.
func (s *Page) IsEmpty() bool {
	if s.Id.String() != "" && s.Id != (uuid.UUID{}) {
		return false
	}
	if !s.Page.IsEmpty() {
		return false
	}
	return true
}

// IsNotEmpty returns true if the Page is not empty.
func (s *Page) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Clear clears the Page.
func (s *Page) Clear() {
	s.Id = uuid.UUID{}
	s.Page.Clear()
}

// Validate validates the Page.
func (s *Page) Validate() error {
	if s.IsEmpty() {
		return mo.ErrorPageIsEmpty
	}
	if err := s.Page.Validate(); err != nil {
		return err
	}
	return nil
}

// Pages contains a slice of *Page and total number of pages.
type Pages struct {
	Pages     []Page `json:"pages"`      // Pages is the slice of *Page.
	TotalRows int64  `json:"total_rows"` // TotalRows is the total number of rows.
}
