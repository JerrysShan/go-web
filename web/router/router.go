package router

import (
	"go-web/web/handlers/home"
	"go-web/web/handlers/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {

	app.GET("/", home.Home)

	v1 := app.Group("/v1")
	{
		u := v1.Group("/user")
		{
			u.GET("/index", user.Index)
			u.GET("/info/", user.Show)
			u.POST("/create", user.Create)
			u.Put("/", user.Update)
			u.DELETE("/", user.Del)
		}
	}
}
