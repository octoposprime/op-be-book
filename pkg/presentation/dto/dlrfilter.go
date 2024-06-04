package presentation

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	me "github.com/octoposprime/op-be-book/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-book/internal/domain/model/object"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/page"
	tuuid "github.com/octoposprime/op-be-shared/tool/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// PageFilter is a struct that represents the filter dto of a page.
type PageFilter struct {
	proto *pb.PageFilter
}

// NewPageFilter creates a new *PageFilter.
func NewPageFilter(pb *pb.PageFilter) *PageFilter {
	return &PageFilter{
		proto: pb,
	}
}

// String returns a string representation of the PageFilter.
func (s *PageFilter) String() string {
	return fmt.Sprintf("Id: %v, "+
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
		s.proto.Id,
		s.proto.PageType,
		s.proto.PageStatus,
		s.proto.Tags,
		s.proto.CreatedAtFrom,
		s.proto.CreatedAtTo,
		s.proto.UpdatedAtFrom,
		s.proto.UpdatedAtTo,
		s.proto.SearchText,
		s.proto.SortType,
		s.proto.SortField,
		s.proto.Limit,
		s.proto.Offset)
}

// NewPageFilterFromEntity creates a new *PageFilter from entity.
func NewPageFilterFromEntity(entity me.PageFilter) *PageFilter {
	id := entity.Id.String()
	pageType := pb.PageType(entity.PageType)
	pageStatus := pb.PageStatus(entity.PageStatus)
	tags := entity.Tags
	createdAtFrom := timestamppb.New(entity.CreatedAtFrom)
	createdAtTo := timestamppb.New(entity.CreatedAtTo)
	updatedAtFrom := timestamppb.New(entity.UpdatedAtFrom)
	updatedAtTo := timestamppb.New(entity.UpdatedAtTo)
	searchText := entity.SearchText
	sortType := entity.SortType
	sortField := pb.PageSortField(entity.SortField)
	limit := int32(entity.Limit)
	offset := int32(entity.Offset)
	return &PageFilter{
		&pb.PageFilter{
			Id:            &id,
			PageType:       &pageType,
			PageStatus:     &pageStatus,
			Tags:          tags,
			CreatedAtFrom: createdAtFrom,
			CreatedAtTo:   createdAtTo,
			UpdatedAtFrom: updatedAtFrom,
			UpdatedAtTo:   updatedAtTo,
			SearchText:    &searchText,
			SortType:      &sortType,
			SortField:     &sortField,
			Limit:         &limit,
			Offset:        &offset,
		},
	}
}

// ToEntity returns a entity representation of the PageFilter.
func (s *PageFilter) ToEntity() *me.PageFilter {
	id := uuid.UUID{}
	if s.proto.Id != nil {
		id = tuuid.FromString(*s.proto.Id)
	}
	pageType := 0
	if s.proto.PageType != nil {
		pageType = int(*s.proto.PageType)
	}
	pageStatus := 0
	if s.proto.PageStatus != nil {
		pageStatus = int(*s.proto.PageStatus)
	}
	tags := []string{}
	if s.proto.Tags != nil {
		tags = s.proto.Tags
	}
	createdAtFrom := time.Time{}
	if s.proto.CreatedAtFrom != nil {
		createdAtFrom = s.proto.CreatedAtFrom.AsTime()
	}
	createdAtTo := time.Time{}
	if s.proto.CreatedAtTo != nil {
		createdAtTo = s.proto.CreatedAtTo.AsTime()
	}
	updatedAtFrom := time.Time{}
	if s.proto.UpdatedAtFrom != nil {
		updatedAtFrom = s.proto.UpdatedAtFrom.AsTime()
	}
	updatedAtTo := time.Time{}
	if s.proto.UpdatedAtTo != nil {
		updatedAtTo = s.proto.UpdatedAtTo.AsTime()
	}
	searchText := ""
	if s.proto.SearchText != nil {
		searchText = string(*s.proto.SearchText)
	}
	sortType := ""
	if s.proto.SortType != nil {
		sortType = string(*s.proto.SortType)
	}
	sortField := 0
	if s.proto.SortField != nil {
		sortField = int(*s.proto.SortField)
	}
	limit := 0
	if s.proto.Limit != nil {
		limit = int(*s.proto.Limit)
	}
	offset := 0
	if s.proto.Offset != nil {
		offset = int(*s.proto.Offset)
	}
	return &me.PageFilter{
		Id:            id,
		PageType:       mo.PageType(pageType),
		PageStatus:     mo.PageStatus(pageStatus),
		Tags:          tags,
		CreatedAtFrom: createdAtFrom,
		CreatedAtTo:   createdAtTo,
		UpdatedAtFrom: updatedAtFrom,
		UpdatedAtTo:   updatedAtTo,
		SearchText:    searchText,
		SortType:      sortType,
		SortField:     mo.PageSortField(sortField),
		Limit:         limit,
		Offset:        offset,
	}
}
