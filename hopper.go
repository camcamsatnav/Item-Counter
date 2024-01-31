package main

import (
	"github.com/buger/jsonparser"
	"strconv"
	"strings"
)

// Hopper takes in a 2D slice of hoppers and a string to search for.
// Finds the count of that item in the hoppers.
// Returns the count and shulker boxes.
func Hopper(hoppers [][]byte, search string) (int, [][]byte) {
	count := 0
	var boxes [][]byte
	for i := 0; i < len(hoppers); i++ {
		_, _ = jsonparser.ArrayEach(hoppers[i], func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			res, _, _, _ := jsonparser.Get(value, "[1]", "value")
			if string(res) == search {
				res2, _, _, _ := jsonparser.Get(value, "[2]", "value")
				temp, _ := strconv.Atoi(string(res2))
				count += temp
			} else if strings.HasSuffix(string(res), "shulker_box") {
				boxes = append(boxes, value)
			}
		}, "nbt", "[5]", "value", "list")
	}
	return count, boxes
}
