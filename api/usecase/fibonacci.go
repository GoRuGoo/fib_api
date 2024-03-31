package usecase

import "fib/repository"

type InterfaceFibonacciUsecase interface{
  Calculate()(int,error)
}

type fibonacciUsecase struct{
  fr repository.InterfaceFibonacciRepository
}

func CreateFibonacciUsecase(fr repository.InterfaceFibonacciRepository) InterfaceFibonacciUsecase{
  return &fibonacciUsecase{fr:fr}
}

func (fc *fibonacciUsecase)Calculate()(int,error){
  return 3,nil
}
