package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	fields := strings.Split(data, ",")
	age, _ := strconv.Atoi(fields[1])
	return User{
		Name:    fields[0],
		Age:     age,
		Address: fields[2],
	}
}

func main() {
	data := "Novan Arri Setiadi,23,Banjarmasin"
	user := ConvertData(data)
	fmt.Println(user)
}
