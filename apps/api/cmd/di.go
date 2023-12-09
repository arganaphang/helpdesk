package main

import (
	"github.com/uptrace/bun"

	"github.com/arganaphang/helpdesk/apps/api/domain"
	"github.com/arganaphang/helpdesk/apps/api/pkg/issue"
	"github.com/arganaphang/helpdesk/apps/api/pkg/user"
)

func Initialize(db *bun.DB) *domain.Services {
	// ? User
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	// ? Issue
	issueRepository := issue.NewRepository(db)
	issueService := issue.NewService(issueRepository)
	services := &domain.Services{
		UserService:  userService,
		IssueService: issueService,
	}
	return services
}
