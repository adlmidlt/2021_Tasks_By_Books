package ex3

import (
	"errors"
	"fmt"
)

func Room() {
	truth := true
	negate(&truth)
	fmt.Println(truth)

	lies := false
	negate(&lies)
	fmt.Println(lies)
	/*	var myFloatPointer = createPointer()
		fmt.Println(*myFloatPointer)

		var myBool bool = true
		printPointer(&myBool)*/
	/*var myInt int
	  myInt = 42
	  myIntPointer := &myInt
	  fmt.Println(*myIntPointer)*/
	/*	myInt := 4
		fmt.Println(myInt)

		// & адрес переменной
		myIntPointer := &myInt

		// * указатель на тип...
		*myIntPointer = 8
		fmt.Println(*myIntPointer)
		fmt.Println(myInt)*/
	/*	myFloat := 4.5
		myFloatPointer := &myFloat
		fmt.Println(myFloatPointer)
		fmt.Println(*myFloatPointer)

		myBool := true
		myBoolPointer := &myBool
		fmt.Println(myBoolPointer)
		fmt.Println(*myBoolPointer)*/
	/*	var myInt int
		myIntPointer := &myInt
		fmt.Println(myIntPointer)

		var myFloat float64
		myFloatPointer := &myFloat
		fmt.Println(myFloatPointer)

		var myBool bool
		myBoolPointer := &myBool
		fmt.Println(myBoolPointer)
	*/
	/*	amount := 6
		double(&amount)
		fmt.Println(amount)*/
	/*	quotient, err := divide(5.6, 0.0)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("%0.2f\n", quotient)
		}*/
	/*	var total, amount float64
		amount, err := paintNeeded(4.2, 3.0)
		fmt.Println(err)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%0.2f liters needed\n", amount)
		total += amount

		amount, err = paintNeeded(-5.2, 3.5)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(err)
		fmt.Printf("%0.2f liters needed\n", amount)
		total += amount
		fmt.Printf("Total: %0.2f liters\n", total)*/
	/*	paintNeeded(4.2, 3.0)
		paintNeeded(5.2, 3.5)
		paintNeeded(2.2, 1.5)

		dozen := double(3.0)
		fmt.Println(dozen)
		fmt.Println(double(6.0))

		fmt.Println(status(70))
		fmt.Println(status(10))*/
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func paintNeeded(width, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height %0.2f is invalid", height)
	}
	area := width * height
	return area / 10, nil
}

/*func paintNeeded(width, height float64) {
	area := width * height
	fmt.Printf("%.2f letters needed\n", area  / 10.0)
}*/

func double(number *int) {
	*number *= 2
}

func status(grade float64) string {
	if grade < 60.0 {
		return "failing"
	}
	return "passing"
}

func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return dividend / divisor, nil
}
