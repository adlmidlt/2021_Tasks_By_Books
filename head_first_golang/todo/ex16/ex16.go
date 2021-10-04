package ex16

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	signatures := getString("signature.txt")
	html, err := template.ParseFiles("ex16/view.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(w, guestbook)
	check(err)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("ex16/new.html")
	check(err)
	err = html.Execute(w, nil)
	check(err)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	signature := r.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("ex16/signature.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintf(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(w, r, "/guestbook", http.StatusFound)
}

func getString(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func MainHandler() {
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/create", createHandler)

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
