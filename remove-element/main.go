package main

import "fmt"

func removeLeftRight(arr []interface{}, position string) []interface{} {
	if position == "left" {
		return arr[1:]
	} else {
		return arr[:len(arr)-1]
	}
}

func main() {
	arr1 := []interface{}{1, 2, 3, 4, 5}
	fmt.Println(removeLeftRight(arr1, "left"))
	fmt.Println(removeLeftRight(arr1, "right"))

	arr2 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(removeLeftRight(arr2, "left"))
	fmt.Println(removeLeftRight(arr2, "right"))
}
