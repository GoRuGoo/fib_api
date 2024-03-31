package usecase

import (
	"errors"
	"fib/repository"
	"strconv"
)

type InterfaceFibonacciUsecase interface{
  Calculate(fibonncciStringNnmber string)(int,error)
}

type fibonacciUsecase struct{
  fr repository.InterfaceFibonacciRepository
}

func CreateFibonacciUsecase(fr repository.InterfaceFibonacciRepository) InterfaceFibonacciUsecase{
  return &fibonacciUsecase{fr:fr}
}

func (fc *fibonacciUsecase)Calculate(fibonacciStringNumber string)(int,error){
  //クエリパラメータは数字を返すので数値に変換
  fibonacciNumber ,err := strconv.Atoi(fibonacciStringNumber)
  if err!=nil{
    return -1,errors.New("Convert error")
  }

  //適切ではないフィボナッチ数列の番号だったらエラーを返す
  if fibonacciNumber <= 0 {
    return -1,errors.New("error")
  }

  fibonacciTable := make([]int,fibonacciNumber+1)
  fibonacciTable[1] = 1
  fibonacciTable[2] = 1

  for i:=3;i<=fibonacciNumber;i++{
    fibonacciTable[i] = fibonacciTable[i-1] + fibonacciTable[i-2]
  }


  return fibonacciTable[fibonacciNumber],nil
}
