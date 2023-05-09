package main

import "fmt"

func removeUnrelated(dataMap map[string]interface{}, key string) map[string]interface{} {
	delete(dataMap, key)
	return dataMap
}

func main() {
	dataMap := map[string]interface{}{
		"name":    "Novan",
		"age":     20,
		"address": "Jakarta",
	}
	key := "address"
	result := removeUnrelated(dataMap, key)
	fmt.Println(result)

	dataMap2 := map[string]interface{}{
		"run":  "lari",
		"jump": "loncat",
		"swim": "berenang",
	}
	key2 := "jump"
	result2 := removeUnrelated(dataMap2, key2)
	fmt.Println(result2)

	dataMap3 := map[string]interface{}{
		"atu":   "ichi",
		"dua":   "ni",
		"tiga":  "san",
		"empat": "yon",
		"lima":  "go",
	}
	key3 := "tiga"
	result3 := removeUnrelated(dataMap3, key3)
	fmt.Println(result3)
}
