package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"github.com/quangdangfit/getjob/pkg/logger"
	"github.com/quangdangfit/getjob/src/api"
)

func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		candidate *api.CandidateAPI,
	) error {
		//--------------------------------API-----------------------------------
		//apiRoute := r.Group("/api/v1")
		{
			//apiRoute.GET("/users/:id", user.GetByID)
			//apiRoute.GET("/users", wrapper.Wrap(user.List))
		}
		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return err
}
