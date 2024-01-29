package main

import (
	"fmt"
	"github.com/midnightfreddie/nbt2json"
	"os"
	"strings"
)

func main() {
	nbt2json.UseJavaEncoding()
	dir, _ := os.ReadDir(".")
	for _, file := range dir {
		if strings.Contains(file.Name(), ".mca") {
			ItemExtractor(file.Name())
			fmt.Println(file.Name())
		}
	}
}

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}
