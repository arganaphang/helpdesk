package dto

type SortQuery struct {
	SortBy   *string `form:"sort_by"`
	SortType *string `form:"sort_type"`
}

type PaginationQuery struct {
	Page    *uint `form:"page"`
	PerPage *uint `form:"per_page"`
}
