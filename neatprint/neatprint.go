package main

import (
	"fmt"
	"strings"
)

const (
	OFFSET = 5
)

func neatPrint(strs [][]string) {
	sliceOfStrs := strs[:]
	maxL, maxR := maxLeftAndRight(sliceOfStrs)
	fmt.Println(strings.Repeat("#", maxL+maxR+OFFSET))
	for _, v := range sliceOfStrs {
		fmt.Print("#" + padString(v[0], maxL, true))
		fmt.Print(" - ")
		fmt.Print(padString(v[1], maxR, false) + "#\n")
	}
	fmt.Println(strings.Repeat("#", maxL+maxR+OFFSET))
}
func maxLeftAndRight(strs [][]string) (maxLeft int, maxRight int) {
	var maxL, maxR int
	for _, v := range strs {
		if maxL < len(v[0]) {
			maxL = len(v[0])
		}
		if maxR < len(v[1]) {
			maxR = len(v[1])
		}
	}
	return maxL + 1, maxR + 1
}

func padString(input string, padLength int, leftPadding bool) string {
	if padLength < len(input) {
		return input
	}
	repeat := padLength - len(input)
	if leftPadding {
		return strings.Repeat(" ", repeat) + input
	} else {
		return input + strings.Repeat(" ", repeat)
	}
}
func main() {
	stringPairs := [][]string{
		{"price", "display price of item"},
		{"owner", "show owner of item"},
		{"update", "update database and then sync cache with database"},
		{"delete", "delete item"},
		{"disassociate", "remove item from owner"},
		{"billing inquiries", "history of billing"}}
	neatPrint(stringPairs)
}
