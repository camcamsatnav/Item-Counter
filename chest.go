package main

import (
	"github.com/buger/jsonparser"
	"strconv"
)

func chest(chests [][]byte, search string) (int, [][]byte) {
	var count = 0
	var boxes [][]byte
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
	return count, boxes
}
