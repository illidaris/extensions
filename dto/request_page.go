package dto

import (
	"strings"
)

// ISortField sort field
type ISortField interface {
	GetField() string
	GetIsDesc() bool
}

var _ = ISortField(&SortField{}) // check impl

// SortField default sort field
type SortField struct {
	Field  string `json:"field" form:"field"`   // order by col
	IsDesc bool   `json:"isDesc" form:"isDesc"` // asc or desc, default is asc
}

func (r *SortField) GetField() string {
	return r.Field
}

func (r *SortField) GetIsDesc() bool {
	return r.IsDesc
}

// IPage page request
type IPage interface {
	GetPageIndex() int32
	GetPageSize() int32
	GetBegin() int
	GetSize() int
	GetAfterID() any
	GetSorts() []ISortField
}

var _ = IPage(&Page{}) // check impl

// Page default page request
type Page struct {
	PageIndex int         `json:"page" form:"page" uri:"page" binding:"required,gte=1"`             // currect page no
	AfterId   interface{} `json:"afterId"`                                                          // previous page last id, when sort by pk
	PageSize  int         `json:"pageSize" form:"pageSize" uri:"pageSize" binding:"required,gte=1"` // page size
	Sorts     []string    `json:"sorts" form:"sorts" uri:"sorts"`                                   // eg; field|desc
}

func (dto *Page) GetPageIndex() int32 {
	return int32(dto.PageIndex)
}

func (dto *Page) GetPageSize() int32 {
	return int32(dto.PageSize)
}

func (dto *Page) GetBegin() int {
	return (dto.PageIndex - 1) * dto.PageSize
}

func (dto *Page) GetSize() int {
	return dto.PageSize
}

func (dto *Page) GetAfterID() any {
	return dto.AfterId
}

func (dto *Page) GetSorts() []ISortField {
	s := []ISortField{}
	for _, v := range dto.Sorts {
		words := strings.Split(v, "|")
		if len(words) == 1 {
			s = append(s, &SortField{Field: words[0]})
		} else if len(words) > 1 {
			s = append(s, &SortField{Field: words[0], IsDesc: words[1] == "desc"})
		}
	}
	return s
}
