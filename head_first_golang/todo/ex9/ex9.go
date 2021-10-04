package ex9

import "fmt"

type Liters float64
type Gallons float64
type Milliliters float64

func BaseType() {
	carFuel := Liters(19).toGallons()
	carFuel = Liters(19).toGallons()
	busFuel := Milliliters(10).toGallons()
	busFuel1 := Gallons(10).toMilliliters()
	fmt.Println("Litters:", carFuel, "Gallons:", busFuel, busFuel1)
}

func (l Liters) toGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) toGallons() Liters {
	return Liters(m * 3.785)
}

func (g Gallons) toLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) toMilliliters() Milliliters {
	return Milliliters(g * 0.264)
}

type Population int

func Pool() {
	population := Population(572)
	fmt.Println("Sleepy Creek Country population: ", population)
	fmt.Println("Congratulation? Kelvin and Anna! It's a girl!")
	population += 1
	fmt.Println("Sleepy Creek Country population: ", population)
}

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hello", m)
}

func SayMain() {
	value := MyType("a MyType value")
	value.sayHi()
	anotherValue := MyType("another value")
	anotherValue.sayHi()

	value = MyType("MyType value ")
	value.MethodWithParameters(4, true)

	fmt.Println(value.WithReturn())
}

func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func (m MyType) WithReturn() int {
	return len(m)
}

type Number int

func (n Number) add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) minus(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func NumberMain() {
	ten := Number(10)
	ten.add(4)
	ten.minus(5)
}
