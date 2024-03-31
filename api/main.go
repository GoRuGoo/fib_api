package main

import (
	"fib/infra"
	"fib/interface/controller"
	"fib/repository"
	"fib/usecase"
  "fmt"
  "os"
)

func main() {
  fibonacciRepository := repository.NewFibonacciRepository()
  fibonacciUsecase := usecase.CreateFibonacciUsecase(fibonacciRepository)
  fibonacciController := controller.CreateFibonacciController(fibonacciUsecase)

  var controllerList infra.ControllerList
  infra.SetFibonacciController(&controllerList,fibonacciController)

  r := controllerList.NewRouter()
  port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.Run(port)
}
