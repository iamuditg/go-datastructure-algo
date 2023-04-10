package main

import "fmt"

func main() {
	s := "Mr  john smith"
	fmt.Println(URLify("Mr  john smith", len(s)))
}
