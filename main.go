package main

import (
	"fmt"
	"github.com/midnightfreddie/nbt2json"
	"os"
	"strings"
	"sync"
)

func main() {
	nbt2json.UseJavaEncoding()
	totalCount := 0
	dirName := "C:\\Users\\camde\\AppData\\Roaming\\PrismLauncher\\instances\\1.20.2\\.minecraft\\saves\\CAMCAM WORLD\\region"

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
				//fmt.Println(file.Name() + " is empty")
				return
			}
			if fileInfo.Size() != 0 && strings.Contains(file.Name(), ".mca") {
				ch <- ItemExtractor(dirName+"\\"+file.Name(), "minecraft:redstone")
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
	fmt.Println(totalCount)
}

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}
