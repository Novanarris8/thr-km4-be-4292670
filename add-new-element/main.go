package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		return append([]int{newData}, data...)
	} else if position == "down" {
		return append(data, newData)
	} else {
		panic("Invalid position")
	}
}

func main() {
	data := []int{1, 2, 3}
	element := 4
	position := "down"

	newData := AddElement(data, element, position)
	fmt.Println(newData)
}
