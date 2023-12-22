package dto

type SortQuery struct {
	SortBy   *string `form:"sort_by"`
	SortType *string `form:"sort_type"`
}
