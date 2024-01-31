package main

import (
	"github.com/buger/jsonparser"
	"strconv"
	"strings"
)

// Chest takes in a 2D slice of chests and a string to search for.
// Finds the count of that item in the chests.
// Returns the count and shulker boxes.
func Chest(chests [][]byte, search string) (int, [][]byte) {
	count := 0
	var boxes [][]byte
	for i := 0; i < len(chests); i++ {
		_, _ = jsonparser.ArrayEach(chests[i], func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			res, _, _, _ := jsonparser.Get(value, "[1]", "value")
			if string(res) == search {
				res2, _, _, _ := jsonparser.Get(value, "[2]", "value")
				temp, _ := strconv.Atoi(string(res2))
				count += temp
			} else if strings.HasSuffix(string(res), "shulker_box") {
				boxes = append(boxes, value)
			}
		}, "nbt", "[4]", "value", "list")

	}
	return count, boxes
}
