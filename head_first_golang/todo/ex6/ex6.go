package ex6

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func Segments() {
	notes := make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "me"
	fmt.Println(notes)
}

func Pool() {
	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 25.2

	for i, number := range numbers {
		fmt.Println(i, number)
	}
	var letters = []string{"a", "b", "c"}
	for i, letters := range letters {
		fmt.Println(i, letters)
	}
}

func Slices() {
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))
	slice = append(slice, "d", "e")
	fmt.Println(slice, len(slice))
}

func SegmentsAndNil() {
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[9], boolSlice[5])
}

func Example() {
	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "first item")
	}
	fmt.Printf("%#v\n", slice)
}

func Arguments() {
	arguments := os.Args[1:]
	var sum float64 = 0
	for _, arg := range arguments {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	sampleCount := float64(len(arguments))
	fmt.Printf("Arguments: %0.2f\n", sum/sampleCount)
}

func Maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func InRage(min, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func MaximumMain() {
	//fmt.Println(Maximum(12.5, 15.2, 92.3))
	fmt.Println(InRage(1, 100, -21.4, -23.1, 54, 11))
	fmt.Println(InRage(-10, 10, -1, -21.4, -2.1, 54, 11))
}
func sum(numbers ...int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func FunMagneto() {
	fmt.Println(sum(4, 7, 1))
	fmt.Println(sum(9, 2))
}

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func AverageMain() {
	fmt.Println(average(100, 50))
	fmt.Println(average(90.7, 89.7, 98.5, 92.3))
}
