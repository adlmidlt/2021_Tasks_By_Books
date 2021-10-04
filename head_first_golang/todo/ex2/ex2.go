package ex2

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func CallMethod() {
	var now = time.Now()
	var year = now.Year()
	fmt.Println(year)
}

func CallMethodContinue() {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}

func PassFile() {
	fmt.Println("Введи код: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}

func FunWithMagneto() {
	fileInfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}

func SubstitutionName() {
	// Так нельзя делать, замещение имен
	/*var int int = 12
	var append string = "minutes of bonus footage"
	var fmt string = "DVD"*/
}

func TransformStringToInt() {
	fmt.Print("Enter a grade: ")
	readers := bufio.NewReader(os.Stdin)
	input, err := readers.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	great, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if great >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println(status)
}

func Example() {
	var a = "a"
	a = "a"
	b := "b"
	if true {
		c := "c"
		if true {
			d := "d"
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
		}
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		//fmt.Println(d)
	}
	fmt.Println(a)
	fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(d)
}

func GameFrom1To100() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)

	var printLostAttempts = 7 // Показать оставшиеся попытки.
	const countAttempts = 7   // Количество попыток.

	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("You have 7 attempts")
	fmt.Println("Can you guess it?")

	for x := 7; x >= countAttempts; x++ {
		readers := bufio.NewReader(os.Stdin)
		input, err := readers.ReadString('\n')
		if err != nil {
			log.Fatal(err)

		}

		guess, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("LOW")
		} else if guess > target {
			fmt.Println("HIGH")
		} else {
			fmt.Println("Success", target)
			break
		}
		printLostAttempts--
		fmt.Println("You have", printLostAttempts, "attempts")

		if printLostAttempts == 0 {
			fmt.Println("Learn algorithm binary tree")
			fmt.Println("Guess number was:", target)
		}
	}
}
