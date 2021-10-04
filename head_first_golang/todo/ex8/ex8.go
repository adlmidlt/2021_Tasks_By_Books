package ex8

import (
	"fmt"
	"todo/todo/ex8/magazine"
)

var myStruct struct {
	number float64
	word   string
	toggle bool
}

func Example() {
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)
}

var pet struct {
	name string
	age  int
}

func Example1() {
	pet.name = "Max"
	pet.age = 5
	fmt.Println("Name:", pet.name)
	fmt.Println("Name:", pet.age)
}

type part struct {
	description string
	count       int
}

func showInfo(p part) {
	fmt.Println("Description: ", p.description)
	fmt.Println("Count: ", p.count)
}

func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func PartMain() {
	p := minimumOrder("Hex Bolts")
	fmt.Println(p.description, p.count)

	var bolts part
	bolts.description = "Hex Box"
	bolts.count = 24
	showInfo(bolts)
}

type car struct {
	name     string
	topSpeed float64
}

func nitroBoost(c *car) {
	c.topSpeed += 50
}

func Auto() {
	var porsche car
	porsche.name = "Porshe 911 R"
	porsche.topSpeed = 323
	fmt.Printf("Name:  %s: ", porsche.name)
	fmt.Printf("Speed: %v:", porsche.topSpeed)
}

//type subscriber struct {
//	name   string
//	rate   float64
//	active bool
//}

func applyDiscount(s *magazine.Subscriber) {
	s.Rate = 4.99
}

func ApplyDiscountMain() {
	var s magazine.Subscriber
	applyDiscount(&s)
	fmt.Println(s.Rate)
}

func printInfo(s *magazine.Subscriber) {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Rate: ", s.Rate)
	fmt.Println("Active: ", s.Active)
}

func defaultSubscriber(name string) *magazine.Subscriber {
	s := magazine.Subscriber{
		Name:   name,
		Rate:   5.99,
		Active: true,
	}
	return &s
}

func Subs() {
	subscriber1 := defaultSubscriber("Aman Sight")
	subscriber1.Rate = 4.99
	fmt.Println(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	fmt.Println(subscriber2)
}

type student struct {
	name  string
	grade float64
}

func printInf(s student) {
	fmt.Println("Name: ", s.name)
	fmt.Printf("Grade: %0.1f\n", s.grade)
}

func StudentMain() {
	s := student{
		name:  "Alonzo Cole",
		grade: 92.3,
	}
	printInf(s)

}

func FieldMain() {
	var value = 2
	var pointer = &value
	fmt.Println(*pointer)
}

func Example2() {
	mustang := car{
		name:     "Mustang Cobra",
		topSpeed: 255,
	}
	nitroBoost(&mustang)
	fmt.Println(mustang.name)
	fmt.Println(mustang.topSpeed)
}

func SubsMain() {
	employee := magazine.Employee{
		Name:   "Joy Carr",
		Salary: 6000,
	}
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)

	address := magazine.Address{
		Street:     "123 Oak St",
		City:       "Omaha",
		State:      "NE",
		PostalCode: "213123",
	}
	fmt.Println(address)
}

func AddressMain() {
	address := magazine.Address{
		Street:     "123 Oak St",
		City:       "Omaha",
		State:      "NE",
		PostalCode: "213123",
	}
	subscriber := magazine.Subscriber{
		Name: "Aman Singh",
	}
	subscriber.Address = address
	fmt.Printf("%#v\n", subscriber.Address)
}
