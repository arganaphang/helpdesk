package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	logger, err := zap.NewProduction()
	if os.Getenv("MODE") == "development" {
		logger, err = zap.NewDevelopment()
	}
	zap.ReplaceGlobals(zap.Must(logger, err))
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()

	db := NewDatabase()
	service := Initialize(db)

	app.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	NewIssue(app, service)
	NewUser(app, service)

	app.Run("0.0.0.0:8000")
}
