package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"github.com/quangdangfit/gocommon/logger"

	"github.com/quangdangfit/getjob/app/api"
	"github.com/quangdangfit/getjob/app/dbs"
	"github.com/quangdangfit/getjob/app/repositories"
	"github.com/quangdangfit/getjob/app/router"
	"github.com/quangdangfit/getjob/app/services"
	"github.com/quangdangfit/getjob/pkg"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// Inject packages
	err := pkg.Inject(container)
	if err != nil {
		logger.Error("Failed to inject packages", err)
	}

	// Inject database
	err = dbs.Inject(container)
	if err != nil {
		logger.Error("Failed to inject database", err)
	}

	// Inject repositories
	err = repositories.Inject(container)
	if err != nil {
		logger.Error("Failed to inject repositories", err)
	}

	// Inject services
	err = services.Inject(container)
	if err != nil {
		logger.Error("Failed to inject services", err)
	}

	// Inject APIs
	err = api.Inject(container)
	if err != nil {
		logger.Error("Failed to inject APIs", err)
	}

	return container
}

func InitGinEngine(container *dig.Container) *gin.Engine {
	app := gin.New()
	err := router.RegisterAPI(app, container)
	if err != nil {
		logger.Fatalf("Cannot init Gin Engine: %s", err)
		return nil
	}

	return app
}
