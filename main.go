package main

import (
	"flag"
	"fmt"
	"github.com/midnightfreddie/nbt2json"
	"os"
	"strings"
	"sync"
)

func main() {
	// get the flags from command line
	givenSearch := flag.String("s", "", "search item")
	givenDir := flag.String("d", "", "directory")
	flag.Parse()

	if *givenSearch == "" || *givenDir == "" {
		fmt.Println("Usage: main.exe -s <search item> -d <directory>")
		return
	}

	nbt2json.UseJavaEncoding()
	totalCount := 0
	//escape the blackslashes
	dirName := strings.Replace(*givenDir, "\\", "\\\\", -1) + "\\\\region"
	//adding minecraft:
	search := "minecraft:" + *givenSearch

	// channel for the total number of items
	ch := make(chan int)
	var wg sync.WaitGroup

	dir, _ := os.ReadDir(dirName)
	for _, file := range dir {
		wg.Add(1)
		file := file
		go func() {
			defer wg.Done()
			fileInfo, err := os.Stat(dirName + "\\" + file.Name())
			fatal(err)
			if fileInfo.Size() == 0 {
				return
			}
			if fileInfo.Size() != 0 && strings.Contains(file.Name(), ".mca") {
				ch <- ItemExtractor(dirName+"\\"+file.Name(), search)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for i := range ch {
		totalCount += i
	}
	fmt.Println(search + ": " + fmt.Sprint(totalCount))
}

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}
