package intf

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameters(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

func (m MyType) MethodWithParameters(f float64) {
	fmt.Println("MethodWithParameters called with", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "MethodWithReturnValue called"
}

func (m MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}
