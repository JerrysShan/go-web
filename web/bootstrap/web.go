package bootstrap

import (
	"go-web/web/router"

	"github.com/gin-gonic/gin"
)

func Start() {
	app := gin.Default()
	router.InitRouter(app)
	app.Run(":3000")
}
