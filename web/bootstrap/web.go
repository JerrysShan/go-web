package bootstrap

import (
	"go-web/web/router"

	"github.com/gin-gonic/gin"
)

func Start(env string) *gin.Engine {
	app := gin.Default()
	app.Run(":3000")
	router.InitRouter(app)
	return app
}
