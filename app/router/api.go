package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"github.com/quangdangfit/gocommon/logger"

	"github.com/quangdangfit/getjob/app/api"
	"github.com/quangdangfit/getjob/app/middleware"
	"github.com/quangdangfit/getjob/pkg/http/wrapper"
)

func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		userAPI *api.UserAPI,
	) error {
		//--------------------------------API-----------------------------------
		apiRoute := r.Group("/api/v1")
		{
			apiRoute.POST("/login", wrapper.Wrap(userAPI.Login))
			apiRoute.POST("/register", wrapper.Wrap(userAPI.Register))
		}
		jwtMid := middleware.JWT(container)
		apiRoute.GET("/profile", jwtMid, wrapper.Wrap(userAPI.GetProfile))

		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return err
}
