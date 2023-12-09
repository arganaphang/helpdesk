package dto

import "github.com/oklog/ulid/v2"

type CreateIssue struct {
	Title         string `form:"title"`
	Detail        string `form:"detail"`
	CustomerName  string `form:"customer_name"`
	CustomerEmail string `form:"customer_email"`
}

type IssueQueryParams struct {
	Q             *string    `form:"q"`
	CustomerName  *string    `form:"customer_name"`
	CustomerEmail *string    `form:"customer_email"`
	TakenBy       *ulid.ULID `form:"taken_by"`
	IsSolved      *bool      `form:"is_solved"`
	SortQuery
	PaginationQuery
}
