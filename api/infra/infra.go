package infra

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func (cl ControllerList)NewRouter() *gin.Engine{
  r := gin.Default()
  config := cors.DefaultConfig()
  config.AllowOrigins = []string{"*"}
  r.Use(cors.New(config))


  r.GET("/fib",cl.fc.Calculate)
  return r
}
