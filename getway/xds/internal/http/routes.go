package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/authservice/getway/xds/internal/app/handler"
)

func routes(svr *gin.Engine) {
	svr.POST("/v3/discovery:endpoints", handler.Endpoints)
	svr.POST("/v3/discovery:clusters", handler.Clusters)
	svr.POST("/v3/discovery:listeners", handler.Listeners)
	svr.POST("/v3/discovery:endpoints", handler.Routes)
}
