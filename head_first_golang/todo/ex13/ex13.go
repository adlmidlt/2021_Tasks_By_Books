package ex13

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Page struct {
	URL  string
	Size int
}

func WebMain() {
	pages := make(chan Page)
	urls := []string{
		"https://example.com/",
		"https://golang.org/",
		"https://golang.org/doc",
	}

	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}

func responseSize(url string, channel chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{
		URL:  url,
		Size: len(body),
	}
}

func ABMain() {
	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println("end main()")
}

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}
