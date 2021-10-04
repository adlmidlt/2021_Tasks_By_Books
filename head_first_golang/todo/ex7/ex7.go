package ex7

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetStringsMain() {
	lines, err := GetStrings("string.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts {
		fmt.Printf("%s: %v\n", name, count)
	}

}

func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err = file.Close(); err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func Numbers() {
	numbers := make(map[string]int)
	numbers["I've been assigned"] = 12
	fmt.Printf("%#v\n", numbers["I've been assigned"])
	fmt.Printf("%#v\n", numbers["I haven't been assigned"])
}

func Example() {
	jewelry := make(map[string]float64)
	jewelry["necklace"] = 89.99
	jewelry["earrings"] = 79.99
	clothing := map[string]float64{
		"pants": 59.99,
		"shirt": 39.99,
	}
	fmt.Println("Earrings:", jewelry["earrings"])
	fmt.Println("Necklace:", jewelry["necklace"])
	fmt.Println("Shirt:", clothing["shirt"])
	fmt.Println("Pants:", clothing["pants"])
}

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	}
}

func GradesMain() {
	status("Alma")
	status("Rohit")
}

func FunMagneto() {
	ranks := map[string]int{
		"bronze": 3,
		"silver": 2,
		"gold":   1,
	}

	for medal, rank := range ranks {
		fmt.Printf("The %s medal's rank is %v\n", medal, rank)
	}
}
