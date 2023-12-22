package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/arganaphang/helpdesk/apps/api/domain"
	"github.com/arganaphang/helpdesk/apps/api/dto"
)

type issueRoute struct {
	service *domain.Services
}

func NewIssue(app *gin.Engine, service *domain.Services) {
	route := issueRoute{service: service}
	r := app.Group("/issue")
	r.POST("", route.create)
	r.GET("", route.get)
	r.GET(":id", route.getByID)
	r.PUT(":id/take", route.take)
	r.PUT(":id/solved", route.solve)
}

func (r issueRoute) create(ctx *gin.Context) {
	body := dto.CreateIssue{}
	if err := ctx.ShouldBind(&body); err != nil {
		zap.L().Warn("request body not valid", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "request body not valid",
		})
		return
	}
	if _, err := r.service.IssueService.Create(ctx, domain.Issue{
		Title:         body.Title,
		Detail:        body.Detail,
		CustomerName:  body.CustomerName,
		CustomerEmail: body.CustomerEmail,
	}); err != nil {
		zap.L().Warn("failed to create issue", zap.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create issue",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "issue created",
	})
}

func (r issueRoute) get(ctx *gin.Context) {
	queries := dto.IssueQueryParams{}
	if err := ctx.BindQuery(&queries); err != nil {
		zap.L().Warn("failed to bind query params", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse query params",
		})
		return
	}

	err := queries.Pagination.Transform(ctx.Request.URL.Query())
	if err != nil {
		zap.L().Warn("failed to bind query params", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse pagination",
		})
		return
	}

	result, count, err := r.service.IssueService.Get(ctx, queries)
	if err != nil {
		zap.L().Warn("failed to get issue", zap.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get issue",
		})
		return
	}

	queries.Pagination.Finish(*count)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get issues",
		"data":    result,
		"meta":    queries.Pagination,
	})
}

func (r issueRoute) getByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		zap.L().Warn("failed to parse id", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse id",
		})
		return
	}

	result, err := r.service.IssueService.GetByID(ctx, id)
	if err != nil {
		zap.L().Warn("failed to get issue by id", zap.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get issue by id",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get issue by id",
		"data":    result,
	})
}

func (r issueRoute) take(ctx *gin.Context) {
	userID, _ := uuid.Parse("018c92e0-0252-389c-9eff-20d87846e018") // TODO: Update this
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		zap.L().Warn("failed to parse id", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse id",
		})
		return
	}

	result, err := r.service.IssueService.Take(ctx, id, userID)
	if err != nil {
		zap.L().Warn("failed to take issue", zap.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to take issue",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "issue was taken",
		"data":    result,
	})
}

func (r issueRoute) solve(ctx *gin.Context) {
	// Right : 018c92e0-0252-389c-9eff-20d87846e018
	// Wrong: 018c92e0-3e63-2aac-59cb-48dfefeb7242
	userID, _ := uuid.Parse("018c92e0-0252-389c-9eff-20d87846e018") // TODO: Update this
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		zap.L().Warn("failed to parse id", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse id",
		})
		return
	}

	result, err := r.service.IssueService.GetByID(ctx, id)
	if err != nil {
		zap.L().Warn("failed to get issue by id", zap.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get issue by id",
		})
		return
	}

	if result.TakenBy != nil && *result.TakenBy != userID {
		zap.L().Warn("authorization failed", zap.Any("taken_by", *result.TakenBy), zap.Any("taken_by", userID))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "you are not allowed to solved this issue",
		})
		return
	}

	result, err = r.service.IssueService.Solve(ctx, id)
	if err != nil {
		zap.L().Warn("failed to solve issue", zap.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to solve issue",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "issue was solved",
		"data":    result,
	})
}
