package controller

import (
	"fib/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InterfaceFibonacciController interface {
  Calculate(c *gin.Context)
}

type fibonacciController struct{
  uf usecase.InterfaceFibonacciUsecase
}

func CreateFibonacciController(uf usecase.InterfaceFibonacciUsecase) InterfaceFibonacciController{
  return &fibonacciController{uf: uf}
}

func (fc *fibonacciController) Calculate(c *gin.Context){
  fibonacciNumber := c.Query("n")
  if fibonacciNumber == ""{
    c.JSON(http.StatusBadRequest,gin.H{"error":"Query parameter does not exist"})
    return
  }


  c.JSON(http.StatusOK,gin.H{"hello":"Hello"})
  return
}
