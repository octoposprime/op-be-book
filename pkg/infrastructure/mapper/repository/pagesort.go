package infrastructure

import (
	mo "github.com/octoposprime/op-be-book/internal/domain/model/object"
)

var PageSortMap map[mo.PageSortField]string = map[mo.PageSortField]string{
	mo.PageSortFieldId:        "id",
	mo.PageSortFieldCreatedAt: "created_at",
	mo.PageSortFieldUpdatedAt: "updated_at",
}
