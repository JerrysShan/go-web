package router

import (
	"go-web/web/handlers/resource"
	"go-web/web/handlers/role"
	"go-web/web/handlers/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {

	app.GET("/")

	v1 := app.Group("/v1")
	{
		u := v1.Group("/user")
		{
			u.GET("/", user.Index)
			u.POST("/", user.Create)
			u.PUT("/", user.Update)
			u.DELETE("/:id", user.Delete)
		}

		re := v1.Group("/resource")
		{
			re.GET("/", resource.Index)
			re.POST("/", resource.Create)
			re.PUT("/", resource.Update)
			re.DELETE("/:id", resource.Delete)
		}
		ro := v1.Group("/role")
		{
			ro.GET("/", role.Index)
			ro.POST("/", role.Create)
			ro.PUT("/", role.Update)
			ro.DELETE("/:id", role.Delete)
		}
	}
}
