package domain

import (
	me "github.com/octoposprime/op-be-book/internal/domain/model/entity"
)

// This is the service layer of the domain layer.
type Service struct {
}

// NewService creates a new *Service.
func NewService() *Service {
	return &Service{}
}

// ValidatePage validates the page.
func (s *Service) ValidatePage(page *me.Page) error {
	if err := page.Validate(); err != nil {
		return err
	}
	return nil
}
