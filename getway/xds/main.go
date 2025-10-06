package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/authservice/getway/xds/handler/clusters"
	"github.com/sajad-dev/authservice/getway/xds/handler/listeners"
)

func serve() {
	ginserve := gin.Default()

	routes(ginserve)

	log.Println("Run server in 8082")
	ginserve.Run(":8082")
}
func routes(r *gin.Engine) {
	log.Println("Register Route")
	r.Any("/v3/*path", func(c *gin.Context) {
		switch c.Request.URL.Path {
		case "/v3/discovery:clusters":
			clusters.Clusters(c)
		case "/v3/discovery:listeners":
			listeners.Listeners(c)
		default:
			c.AbortWithStatus(404)
		}
	})
}
func main() {
	serve()
}
