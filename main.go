package main

import "fmt"

import "strings"

func main() {
	fmt.Println(123)
	s := "123,456"
	st := strings.Split(s, ",")
	fmt.Printf("%+v", st)
}
