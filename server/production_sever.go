package production

import "github.com/gin-gonic/gin"

func ProductionRouter(route *gin.Engine) {
	route.LoadHTMLGlob("templates/production/*.html")
}
