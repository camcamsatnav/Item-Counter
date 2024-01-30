package main

import (
	"fmt"
	"github.com/Tnze/go-mc/level"
	"github.com/Tnze/go-mc/save"
	"github.com/Tnze/go-mc/save/region"
	"github.com/buger/jsonparser"
	"github.com/midnightfreddie/nbt2json"
	"strings"
	"sync"
)

// ItemExtractor takes in a region file and a string to search for.
// Finds the count of that item in the chests, hoppers, and shulker boxes.
// Returns the count.
func ItemExtractor(file string, search string) int {
	var chests [][]byte
	var hoppers [][]byte
	var boxes [][]byte

	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	ch := make(chan []byte)
	ch2 := make(chan []byte)
	done := make(chan bool)

	mcRegion, err := region.Open(file)
	if err != nil {
		fmt.Println(file)
		panic(err)
	}
	wg2.Add(2)
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			if !mcRegion.ExistSector(i, j) {
				continue
			}
			chunk, err := mcRegion.ReadSector(i, j)
			if err != nil {
				continue
			}
			wg.Add(1)
			go processChunk(chunk, ch, ch2, &wg)
		}
	}
	go func() {
		wg.Wait()
		close(ch)
		close(ch2)
		done <- true
	}()
	go func() {
		defer wg2.Done()
		for value := range ch {
			chests = append(chests, value)
		}
	}()
	go func() {
		defer wg2.Done()
		for value2 := range ch2 {
			hoppers = append(hoppers, value2)
		}
	}()

	wg2.Wait()
	_ = mcRegion.Close()
	<-done
	var totalCount = 0

	count, box := Chest(chests, search)
	totalCount += count
	boxes = append(boxes, box...)

	count, box = Hopper(hoppers, search)
	totalCount += count
	boxes = append(boxes, box...)

	count = ShulkerBox(boxes, search)
	totalCount += count

	return totalCount
}

// processChunk takes in a chunk, a channel for each block entity type, and a waitgroup.
// Finds the block entities and sends them to the appropriate channel.
func processChunk(chunk []byte, ch chan []byte, ch2 chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	var c save.Chunk
	err := c.Load(chunk)
	if err != nil {
		return
	}
	levelChunk, err := level.ChunkFromSave(&c)
	if err != nil {
		return
	}
	data := levelChunk.BlockEntity
	for i := 0; i < len(data); i++ {
		if strings.Contains(data[i].Data.String(), "Items") {
			jsonOb, err := nbt2json.Nbt2Json(data[i].Data.Data, "")
			fatal(err)
			res, datatype, _, _ := jsonparser.Get(jsonOb, "nbt", "[2]", "value")
			res2, datatype2, _, _ := jsonparser.Get(jsonOb, "nbt", "[3]", "value")
			if datatype == jsonparser.String && string(res) == "minecraft:chest" {
				ch <- jsonOb
			}
			if datatype2 == jsonparser.String && string(res2) == "minecraft:hopper" {
				ch2 <- jsonOb
			} //add droppers and shulkers maybe
		}
	}
}
