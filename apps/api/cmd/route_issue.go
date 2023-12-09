package main

import (
	"github.com/gin-gonic/gin"

	"github.com/arganaphang/helpdesk/apps/api/domain"
)

func NewIssue(app *gin.Engine, service *domain.Services) {
	r := app.Group("/issue")
	r.POST("/", func(ctx *gin.Context) {})
	r.GET("/", func(ctx *gin.Context) {})
	r.GET("/:id", func(ctx *gin.Context) {})
	r.PUT("/:id", func(ctx *gin.Context) {})
}
