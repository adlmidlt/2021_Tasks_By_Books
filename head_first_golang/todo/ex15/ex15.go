package ex15

import (
	"fmt"
	"log"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	message := []byte("Hello, web!")
	_, err := w.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func writer(w http.ResponseWriter, message string) {
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(w http.ResponseWriter, r *http.Request) {
	writer(w, "Hello, web!")
}

func frenchHandler(w http.ResponseWriter, r *http.Request) {
	writer(w, "Salut, web!")
}

func hindiHandler(w http.ResponseWriter, r *http.Request) {
	writer(w, "Namaste, web!")
}

func ViewHandlerFunc() {
	http.HandleFunc("/hello", viewHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func Pool() {
	callFunction(functionA)
	callTwice(functionA)
	callWithArguments(functionC)
	printReturnValue(functionB)
}

func callFunction(passedFunction func()) {
	passedFunction()
}

func callTwice(passedFunction func()) {
	passedFunction()
	passedFunction()
}

func callWithArguments(passedFunction func(string, bool)) {
	fmt.Println("This sentence is", false)
}

func printReturnValue(passedFunction func() string) {
	fmt.Println(passedFunction())
}

func functionA() {
	fmt.Println("function called")
}

func functionB() string {
	fmt.Println("function called")
	return "Returning from function"
}

func functionC(a string, b bool) {
	fmt.Println("function called")
	fmt.Println(a, b)
}
