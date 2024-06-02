package domain

import (
	"errors"

	smodel "github.com/octoposprime/op-be-shared/pkg/model"
)

var ERRORS []error = []error{
	ErrorNone,
	ErrorPageIsEmpty,
	ErrorPagePageDataIsEmpty,
}

const (
	ErrId       string = "id"
	ErrPage     string = "page"
	ErrPageData string = "pagedata"
)

const (
	ErrEmpty         string = "empty"
	ErrTooShort      string = "tooshort"
	ErrTooLong       string = "toolong"
	ErrNotValid      string = "notvalid"
	ErrInactive      string = "inactive"
	ErrAlreadyExists string = "alreadyexists"
)

var (
	ErrorNone                error = nil
	ErrorPageIsEmpty         error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrPage + smodel.ErrSep + ErrEmpty)
	ErrorPagePageDataIsEmpty error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrPage + smodel.ErrSep + ErrPageData + smodel.ErrSep + ErrEmpty)
)

func GetErrors() []error {
	return ERRORS
}
