package example

import (
	"fmt"
	"regexp"
	"strings"
)

const refString = "Mary had a little lamb"
const refString1 = "Mary_had a little_lamb"
const refString2 = "Mary*had,a%little_lamb"

/* Разделить предложение по словам. */

func Fields() {
	words := strings.Fields(refString)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}

func Split() {
	words := strings.Split(refString1, "_")
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}

func SplitFunc() {
	splitFun := func(r rune) bool {
		return strings.ContainsRune("*%,_", r)
	}

	words := strings.FieldsFunc(refString2, splitFun)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}

func Regex() {
	words := regexp.MustCompile("[*,%_]").Split(refString2, -1)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}

func Example() {
	lookFor := "lamb"
	contain := strings.Contains(refString, lookFor)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)

	lookFor = "wolf"
	contain = strings.Contains(refString, lookFor)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)

	startsWith := "Mary"
	starts := strings.HasPrefix(refString, startsWith)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, startsWith, starts)

	endWith := "lamb"
	ends := strings.HasSuffix(refString, endWith)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, endWith, ends)
}
