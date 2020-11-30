package _map

const (
	DefaultPage     = 1
	DefaultPageSize = 10
)

var DefaultPageRequest = PageList{
	Page:     DefaultPage,
	PageSize: DefaultPageSize,
}

type PageList struct {
	Page     int64 `json:"page" form:"page" validate:"required,number" label:"页码"`
	PageSize int64 `json:"page_size" form:"page_size" validate:"required,number" label:"页码大小"`
}

type IdMap struct {
	Id uint `json:"id" form:"id" validate:"required,number,min=1" label:"id"`
}
