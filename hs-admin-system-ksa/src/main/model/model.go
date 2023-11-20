package model

type Page[T any] struct {
	Offset int    `form:"offset"`
	Limit  int    `form:"limit"`
	Sort   string `form:"sort"`
	Order  string `form:"order"`
}

type DataPage[T any] struct {
	PageNo     int    `json:"pageNo"`
	PageSize   int    `json:"pageSize"`
	OrderBy    string `json:"orderBy"`
	Order      string `json:"order"`
	DataList   []T    `json:"dataList"`
	TotalCount int64  `json:"totalCount"`
}

type PageVo[T any] struct {
	Rows  []T   `json:"rows"`
	Total int64 `json:"total"`
}

func (page Page[T]) ToDataPage() DataPage[T] {
	dataPage := new(DataPage[T])
	var limit = 20
	if page.Limit != 0 {
		limit = page.Limit
	}
	var orderBy = "id"
	if page.Sort != "" {
		orderBy = page.Sort
	}

	var order = "DESC"
	if page.Sort != "" {
		order = page.Order
	}

	dataPage.PageNo = page.Offset/limit + 1
	dataPage.PageSize = limit
	dataPage.OrderBy = orderBy
	dataPage.Order = order
	return *dataPage
}

func (dataPage DataPage[T]) ToPageVo() PageVo[T] {
	pageVo := new(PageVo[T])
	pageVo.Rows = dataPage.DataList
	pageVo.Total = dataPage.TotalCount
	return *pageVo
}
