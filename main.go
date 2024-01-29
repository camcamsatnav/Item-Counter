package main

import (
	"fmt"
	"github.com/Tnze/go-mc/level"
	"github.com/Tnze/go-mc/save"
	"github.com/Tnze/go-mc/save/region"
	"github.com/buger/jsonparser"
	"github.com/midnightfreddie/nbt2json"
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
	var totalCount = 0
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

	var search = "minecraft:spruce_log"

	count, box := Chest(chests, search)
	totalCount += count
	boxes = append(boxes, box...)

	count = Shulkerbox(boxes, search)
	totalCount += count

	fmt.Println(totalCount)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
