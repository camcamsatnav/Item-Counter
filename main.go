package main

import (
	"fmt"
	"github.com/Tnze/go-mc/level"
	"github.com/Tnze/go-mc/save"
	"github.com/Tnze/go-mc/save/region"
	"github.com/buger/jsonparser"
	"github.com/midnightfreddie/nbt2json"
	"strconv"
	"strings"
)

func main() {
	nbt2json.UseJavaEncoding()
	regionExtract(1, 1)
}

func regionExtract(x int, z int) {
	var mcRegion, err = region.Open("r.0.-2.mca")
	var chests [][]byte
	var boxes [][]byte
	if err == nil {
		for i := 0; i < 32; i++ {
			for j := 0; j < 32; j++ {
				//go routine here
				if !mcRegion.ExistSector(i, j) {
					continue
				}
				var chunk, _ = mcRegion.ReadSector(i, j)
				var c save.Chunk
				_ = c.Load(chunk)
				levelChunk, err := level.ChunkFromSave(&c)
				if err != nil {
					continue
				}
				data := levelChunk.BlockEntity
				for i := 0; i < len(data); i++ {
					if strings.Contains(data[i].Data.String(), "Items") {
						jsonOb, err := nbt2json.Nbt2Json(data[i].Data.Data, "")
						check(err)
						res, datatype, _, _ := jsonparser.Get(jsonOb, "nbt", "[2]", "value")
						if datatype == jsonparser.String && string(res) == "minecraft:chest" {
							chests = append(chests, jsonOb)
						}
						//_ = os.WriteFile("final.json", jsonOb, 0666) //replace with new thing -> channel later
					}
				}

			}
		}
		_ = mcRegion.Close()
	}
	var search = "minecraft:redstone"
	var count = 0
	for i := 0; i < len(chests); i++ {
		_, _ = jsonparser.ArrayEach(chests[i], func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			res, _, _, _ := jsonparser.Get(value, "[1]", "value")
			if string(res) == search {
				res2, _, _, _ := jsonparser.Get(value, "[2]", "value")
				//fmt.Println(string(res))
				//fmt.Println(string(res2))
				temp, _ := strconv.Atoi(string(res2))
				count += temp
			} else if string(res) == "minecraft:white_shulker_box" {
				boxes = append(boxes, value)
			}
		}, "nbt", "[4]", "value", "list")
	}
	//fmt.Println(string(boxes[0]))
	for i := 0; i < len(boxes); i++ {
		_, _ = jsonparser.ArrayEach(boxes[i], func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			res, _, _, _ := jsonparser.Get(value, "[2]", "value")
			if string(res) == search {
				res2, _, _, _ := jsonparser.Get(value, "[1]", "value")
				temp, _ := strconv.Atoi(string(res2))
				count += temp
			}
		}, "[2]", "value", "[0]", "value", "[1]", "value", "list")
	}
	fmt.Println(count)
	//fmt.Println(len(chests))
	//fmt.Println(string(boxes[0]))
	//search := "minecraft:melon_seeds"
	//res, _ := http.Get(`http://localhost:3696/chest/` + search)
	//data, _ := ioutil.ReadAll(res.Body)
	//var count Count
	//_ = json.Unmarshal(data, &count)
	//
	//fmt.Println(count.Count)

	//jsonStr, _ := json.Marshal(items)
	//var data []string //going typescript mode LOLSMOOOOOOTH
	//err = json.Unmarshal(jsonStr, &data)
	//check(err)
	//for i := 0; i < len(data); i++ {
	//	var data2 map[string]interface{}
	//	err = json.Unmarshal([]byte(data[i]), &data2)
	//	check(err)
	//	fmt.Println(data2["Items"])
	//}

	// BROOOOOO
	//unpause i dont got this
	//for i := 0; i < len(items); i++ {
	//	if strings.Contains(items[i], "minecraft:chest") {
	//		var chest Chest
	//		//fmt.Println(items[i])
	//		_ = json.Unmarshal([]byte(items[i]), &chest)
	//		//fmt.Println(chest)
	//	}
	//	//println(items[i])
	//	//hmm := items[i].id
	//}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
