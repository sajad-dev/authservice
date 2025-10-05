package http

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Serve() {
	ginserve := gin.Default()

	routes(ginserve)

	log.Println("Run server in 8082")
	ginserve.Run(":8082")
}
