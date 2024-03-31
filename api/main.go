package main

import (
	"fib/infra"
	"fib/interface/controller"
	"fib/repository"
	"fib/usecase"
)

func main() {
  fibonacciRepository := repository.NewFibonacciRepository()
  fibonacciUsecase := usecase.CreateFibonacciUsecase(fibonacciRepository)
  fibonacciController := controller.CreateFibonacciController(fibonacciUsecase)

  var controllerList infra.ControllerList
  infra.SetFibonacciController(&controllerList,fibonacciController)

  r := controllerList.NewRouter()
  r.Run(":80")
}
