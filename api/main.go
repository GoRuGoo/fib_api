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

  r := infra.NewRouter(fibonacciController)
  r.Run(":80")
}
