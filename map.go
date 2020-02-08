package main

import "fmt"

func main() {
	maps := map[string]int{
		"age":  22,
		"suji": 22,
	}

	fmt.Println(maps)

	var map2 = make(map[string]string)
	fmt.Println(map2)

}
