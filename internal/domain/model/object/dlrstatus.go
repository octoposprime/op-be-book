package domain

// PageStatus is a status that represents the status of a page.
type PageStatus int8

const (
	PageStatusNONE PageStatus = iota
	PageStatusACTIVE
	PageStatusINACTIVE
)
