package my_directory

import (
	"fmt"
	"io/ioutil"
	"log"
)

func DirMain() {
	files, err := ioutil.ReadDir("ex12/my_directory")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
		} else {
			fmt.Println("File:", file.Name())
		}
	}
}
