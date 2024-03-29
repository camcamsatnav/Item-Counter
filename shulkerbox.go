package main

import (
	"github.com/buger/jsonparser"
	"strconv"
)

// ShulkerBox takes in a 2D slice of shulkerboxes and a string to search for.
// Finds the count of that item in the boxes.
// Returns the count.
func ShulkerBox(boxes [][]byte, search string) int {
	count := 0
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
	return count
}
