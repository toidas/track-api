package approuter

import (
	"../../controllers/appcontroller"

	"github.com/gin-gonic/gin"
)

func AppRouter(route *gin.Engine) {
	app := route.Group("")
	app.GET("/:token", appcontroller.GetAppWithToken)
}
