package main

import "fmt"

func howManyElements(data []any) int {
	return len(data)
}

func main() {
	data := []interface{}{1, 2, 3, 4, 5}
	count := howManyElements(data)
	fmt.Println(count)

	data = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	count = howManyElements(data)
	fmt.Println(count)

	data = []interface{}{1, 1, 1, 5, 5, 5}
	count = howManyElements(data)
	fmt.Println(count)

	data = []interface{}{"Edo", "Budi", "Joko", "Tono"}
	count = howManyElements(data)
	fmt.Println(count)

	data = []interface{}{"Edo", "Budi", "Joko", "Tono", "Edo", "Budi", "Joko", "Tono"}
	count = howManyElements(data)
	fmt.Println(count)

	data = []interface{}{true, false, true, false, true, false}
	count = howManyElements(data)
	fmt.Println(count)
}
