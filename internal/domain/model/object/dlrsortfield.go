package domain

// PageSortField is a type that represents the sort fields of a page.
type PageSortField int8

const (
	PageSortFieldNONE PageSortField = iota
	PageSortFieldId
	//PageSortFieldName
	PageSortFieldCreatedAt
	PageSortFieldUpdatedAt
)
