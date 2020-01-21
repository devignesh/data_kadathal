package main

import "fmt"

// define a new type
type person struct {
    name string
    age  int
}
func Older(p1, p2 person) (person, int) {
    if p1.age > p2.age {
        return p1, p1.age - p2.age
    }
    return p2, p2.age - p1.age
}
