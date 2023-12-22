package dto

import (
	"github.com/google/uuid"

	"github.com/arganaphang/helpdesk/apps/api/pkg/common/pagination"
)

type CreateIssue struct {
	Title         string `json:"title"`
	Detail        string `json:"detail"`
	CustomerName  string `json:"customer_name"`
	CustomerEmail string `json:"customer_email"`
}

type IssueQueryParams struct {
	Q             *string    `form:"q"`
	CustomerName  *string    `form:"customer_name"`
	CustomerEmail *string    `form:"customer_email"`
	TakenBy       *uuid.UUID `form:"taken_by"`
	IsSolved      *bool      `form:"is_solved"`
	SortQuery
	pagination.Pagination
}
