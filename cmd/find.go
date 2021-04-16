package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	flag.Parse()

	args := flag.Args()

	searchPath := args[0]

	err := filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

	if err != nil {
		panic(err)
	}
}
