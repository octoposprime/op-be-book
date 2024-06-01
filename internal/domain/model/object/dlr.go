package domain

import (
	"fmt"
)

// Page is a struct that represents the object of a page basic values.
type Page struct {
	PageData   string     `json:"page_data"`   // PageData is the data of the page.
	PageType   PageType   `json:"page_type"`   // PageType is the type of the page.
	PageStatus PageStatus `json:"page_status"` // PageStatus is the status of the page.
	Tags       []string   `json:"tags"`        // Tags is the tags of the page.
}

// NewPage creates a new *Page.
func NewPage(pageData string,
	pageType PageType,
	pageStatus PageStatus,
	tags []string) *Page {
	return &Page{
		PageData:   pageData,
		PageType:   pageType,
		PageStatus: pageStatus,
		Tags:       tags,
	}
}

// NewEmptyPage creates a new *Page with empty values.
func NewEmptyPage() *Page {
	return &Page{
		PageData:   "",
		PageType:   PageTypeNONE,
		PageStatus: PageStatusNONE,
		Tags:       []string{},
	}
}

// String returns a string representation of the Page.
func (s *Page) String() string {
	return fmt.Sprintf("PageData: %v, "+
		"PageType: %v, "+
		"PageStatus: %v, "+
		"Tags: %v",
		s.PageData,
		s.PageType,
		s.PageStatus,
		s.Tags)
}

// Equals returns true if the Page is equal to the other Page.
func (s *Page) Equals(other *Page) bool {
	if s.PageData != other.PageData {
		return false
	}
	if s.PageType != other.PageType {
		return false
	}
	if s.PageStatus != other.PageStatus {
		return false
	}
	for i := range s.Tags {
		if s.Tags[i] != other.Tags[i] {
			return false
		}
	}
	return true
}

// Clone returns a clone of the Page.
func (s *Page) Clone() *Page {
	return &Page{
		PageData:   s.PageData,
		PageType:   s.PageType,
		PageStatus: s.PageStatus,
		Tags:       s.Tags,
	}
}

// IsEmpty returns true if the Page is empty.
func (s *Page) IsEmpty() bool {
	if s.PageData != "" {
		return false
	}
	if s.PageType != PageTypeNONE {
		return false
	}
	if s.PageStatus != PageStatusNONE {
		return false
	}
	if len(s.Tags) != 0 {
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
	s.PageData = ""
	s.PageType = PageTypeNONE
	s.PageStatus = PageStatusNONE
	s.Tags = []string{}
}

// Validate validates the Page.
func (s *Page) Validate() error {
	if s.IsEmpty() {
		return ErrorPageIsEmpty
	}
	if s.PageData == "" {
		return ErrorPagePageDataIsEmpty
	}
	return nil
}
