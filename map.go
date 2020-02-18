package main

import "fmt"

func main() {
	maps := map[string]int{
		"age":  22,
		"suji": 22,
	}

	fmt.Println(maps)

	mapsyn := make(map[string]string)
	fmt.Println(mapsyn)

}
