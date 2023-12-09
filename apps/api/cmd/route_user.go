package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"go.uber.org/zap"

	"github.com/arganaphang/helpdesk/apps/api/domain"
	"github.com/arganaphang/helpdesk/apps/api/dto"
)

type route struct {
	service *domain.Services
}

func NewUser(app *gin.Engine, service *domain.Services) {
	route := route{service: service}
	app.POST("/login", route.login)

	r := app.Group("/user")
	r.GET("/:id", route.getByID)
	r.POST("/", route.create)
}

func (r route) login(ctx *gin.Context) {
	body := &dto.Login{}
	if err := ctx.ShouldBind(body); err != nil {
		zap.L().Warn("request body not valid", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "request body not valid",
		})
		return
	}
	user, err := r.service.UserService.GetByEmail(ctx, body.Email)
	if err != nil {
		zap.L().Warn("user not found", zap.String("email", body.Email), zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}
	// TODO: Compare Hash
	// TODO: Create JWT
	// TODO: Create Cookies
	// TODO: Return JWT
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (r route) getByID(ctx *gin.Context) {
	id, err := ulid.Parse(ctx.Param("id"))
	if err != nil {
		zap.L().Warn("id not valid", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id not valid",
		})
		return
	}

	user, err := r.service.UserService.GetByID(ctx, id)
	if err != nil {
		zap.L().Warn("failed to get user", zap.Any("id", id), zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to get user",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (r route) create(ctx *gin.Context) {
	body := &dto.CreateUser{}
	if err := ctx.ShouldBind(body); err != nil {
		zap.L().Warn("request body not valid", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "request body not valid",
		})
		return
	}

	// TODO: Hash password
	user, err := r.service.UserService.Create(ctx, domain.User{Name: body.Name, Email: body.Email, Password: body.Password})
	if err != nil {
		zap.L().Warn("failed to create user", zap.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to create user",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}
