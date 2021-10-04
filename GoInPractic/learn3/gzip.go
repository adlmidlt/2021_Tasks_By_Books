package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var i = -1
	var file string
	for _, file = range os.Args[1:] {
		wg.Add(1)

		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
	}
	wg.Wait()
	fmt.Printf("Compreesed %d files\n", i+1)
}

func compress(file string) error {
	in, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer in.Close()

	out, err := os.Create(file + ".gz")
	if err != nil {
		log.Fatalln(err)
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()
	return err
}
