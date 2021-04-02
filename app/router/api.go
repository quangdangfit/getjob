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
		companyAPI *api.CompanyAPI,
	) error {
		jwtMid := middleware.JWT(container)
		//-------------------------------Public API-----------------------------
		public := r.Group("/api/v1")
		{
			public.POST("/login", wrapper.Wrap(userAPI.Login))
			public.POST("/register", wrapper.Wrap(userAPI.Register))
		}

		//-------------------------------Private API----------------------------
		private := r.Group("/api/v1", jwtMid)
		{
			private.GET("/profile", jwtMid, wrapper.Wrap(userAPI.GetProfile))
			private.POST("/companies", jwtMid, wrapper.Wrap(companyAPI.Create))
		}

		//-------------------------------Internal API----------------------------
		//internal := r.Group("/api/v1/internal")

		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return err
}
