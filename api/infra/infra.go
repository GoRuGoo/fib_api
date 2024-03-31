package infra

import (
	"fib/interface/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func NewRouter(fc controller.InterfaceFibonacciController) *gin.Engine{
  r := gin.Default()
  config := cors.DefaultConfig()
  config.AllowOrigins = []string{"*"}
  r.Use(cors.New(config))


  r.GET("/fib",fc.Calculate)
  return r
}
