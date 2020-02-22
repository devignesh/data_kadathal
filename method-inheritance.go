package main

import "fmt"

type Human struct {
    name  string
    age   int
    phone string
}

type Student struct {
    Human  // anonymous field
    school string
}

type Employee struct {
    Human
    company string
}


func (hu *Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", hu.name, hu.phone)
}

func main() {
    vicky := Employee{Human{"vicky", 22, "111-888-XXXX"}, "Golang Inc"}
    suji := Student{Human{"suji", 25, "222-222-YYYY"}, "pagal_suji"}

    vicky.SayHi()
    suji.SayHi()
}
