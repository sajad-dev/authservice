package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/authservice/getway/xds/internal/app/handler"
)

func routes(r *gin.Engine) {
	log.Println("Register Route")
	r.Any("/v3/*path", func(c *gin.Context) {
		switch c.Request.URL.Path {
		case "/v3/discovery:clusters":
			handler.Clusters(c)
		case "/v3/discovery:listeners":
			handler.Listeners(c)
		default:
			c.AbortWithStatus(404)
		}
	})
}
