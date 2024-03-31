package infra

import "fib/interface/controller"

type ControllerList struct{
  fc controller.InterfaceFibonacciController
}

func SetFibonacciController(cl *ControllerList, fc controller.InterfaceFibonacciController){
  cl.fc = fc
}
