package usecase

import (
	"errors"
	"fib/repository"
	"math/big"
	"strconv"
)

type InterfaceFibonacciUsecase interface{
  Calculate(fibonncciStringNnmber string)(*big.Int,error)
}

type fibonacciUsecase struct{
  fr repository.InterfaceFibonacciRepository
}

func CreateFibonacciUsecase(fr repository.InterfaceFibonacciRepository) InterfaceFibonacciUsecase{
  return &fibonacciUsecase{fr:fr}
}

func (fc *fibonacciUsecase)Calculate(fibonacciStringNumber string)(*big.Int,error){
  //クエリパラメータは数字を返すので数値に変換
  fibonacciNumber ,err := strconv.Atoi(fibonacciStringNumber)
  if err!=nil{
    return big.NewInt(-1),errors.New("Convert error")
  }

  //適切ではないフィボナッチ数列の番号だったらエラーを返す
  if fibonacciNumber <= 0 {
    return big.NewInt(-1),errors.New("error")
  }

  //

  fibonacciTable := make([]*big.Int,fibonacciNumber+2)
  fibonacciTable[1] = big.NewInt(1)
  fibonacciTable[2] = big.NewInt(1)

  for i:=3;i<=fibonacciNumber;i++{
    fibonacciTable[i] = big.NewInt(0)
    fibonacciTable[i].Add(fibonacciTable[i],fibonacciTable[i-1])
    fibonacciTable[i].Add(fibonacciTable[i],fibonacciTable[i-2])
  }


  return fibonacciTable[fibonacciNumber],nil
}
