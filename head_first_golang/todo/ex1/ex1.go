package ex1

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func FmtPrintLn() {
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first go"))
}

func Pool() {
	fmt.Println("Cannonball!")
}

func TemplateGo() {
	fmt.Println("Hello, Go!")
	fmt.Println("Hello, \nGo!")   // \n
	fmt.Println("Hello, \tGo!")   // \t
	fmt.Println("Hello, \"Go!\"") // " "
	fmt.Println("Hello, \\Go!")   // \
}

func RuneLiteral() {
	fmt.Println('\n')
}

func TypesGo() {
	//fmt.Println(math.Floor("head first go"))
	//fmt.Println(strings.Title(2.75))
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf("Hello, Go"))
	fmt.Println(reflect.TypeOf(42.3232))
	fmt.Println(reflect.TypeOf(true))
}

func VariableDeclaration() {
	var quantity = 2
	var length, width = 2.34, 3.434
	var customersName = "Damon Cole"

	fmt.Println(quantity)
	fmt.Println(length * width)
	fmt.Println(customersName)
}

func FunWithMagneto() {
	var originalCount = 10
	var eatenCount = 4
	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")
}

func ShortVariableDeclaration() {
	quantity := 4
	length, width := 1.2, 2.4
	customersName := "Daemon Cole"

	fmt.Println(quantity)
	fmt.Println(length * width)
	fmt.Println(customersName)
}

func Transformation() {
	var length float64 = 1.2
	var width int = 2
	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width?", length > float64(width))
}

func Example() {
	var price int = 100
	fmt.Println("Price is", price, "dollars.")
	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")
	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")
	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", total <= float64(availableFunds))
}
