package ex5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Notes() {
	notes := [7]string{"do", "re", "mi", "fa"}

	for _, note := range notes {
		fmt.Println(note)
	}

	//var notes = [7]string{"do", "re", "mi", "fa"}
	/*notes[0] = "do"
	notes[1] = "re"
	notes[2] = "rm"
	notes[3] = "fa"*/

	/*	for i := 0; i <=6; i++ {
		fmt.Println(i, notes[i])
	}*/

	// безопаснее
	/*	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}*/

	//fmt.Println(len(notes))

	//fmt.Println(notes[0], notes[1], notes[2], notes[1])
	//fmt.Println(notes)
	//fmt.Printf("%#v\n", notes)

	//fmt.Println()
	//fmt.Println(notes[4])
	//fmt.Println(notes[6])
}

func Numbers() {
	numbers := [4]float64{1, 2, 3, 4}
	var numSum float64
	for _, num := range numbers {
		numSum += num
	}
	fmt.Println(numSum)

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", numSum/sampleCount)

	//var numbers = [4]int{1, 2, 3, 4}
	//numbers[0] = 1
	//numbers[1] = 2
	//fmt.Println(numbers[0], numbers[1], numbers[2], numbers[3])
	//fmt.Println(numbers)
	//fmt.Println(numbers[1])
	//fmt.Println(numbers[2])
	//fmt.Println(numbers[3])
}

func Time() {
	/*	var dates [3]time.Time
		dates[0] = time.Unix(1257894000, 0)
		dates[1] = time.Unix(1447920000, 0)
		dates[2] = time.Unix(1508632200, 0)
		fmt.Println(dates[1])

		fmt.Println(dates)*/
}

func Pool() {
	numbers := [6]int{3, 16, -2, 10, 23, 12}
	for i, number := range numbers {
		if number >= 10 && number <= 20 {
			fmt.Println(i, number)
		}
	}
}

func FileRead() {
	numbers, err := getFloats("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}

func getFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
