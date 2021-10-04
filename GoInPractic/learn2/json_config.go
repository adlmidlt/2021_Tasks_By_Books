package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	if err := decoder.Decode(&conf); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)
}
