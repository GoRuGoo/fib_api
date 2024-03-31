package repository

type InterfaceFibonacciRepository interface{}

type fibonacciRepository struct{}


func NewFibonacciRepository()InterfaceFibonacciRepository{
  return & fibonacciRepository{}
}
