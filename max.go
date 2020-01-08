package main

import "fmt"


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    x := 3
    y := 4
    z := 5

    max_xy := max(x, y) 
    max_xz := max(x, z)
    min_xy := min(x, y)
    min_xz := min(x, z)
    

    fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
    fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
    fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z))
    fmt.Println("min(%d, %d) = %d\n", x, y, min_xy)
    fmt.Println("min(%d, %d) = %d\n", x, z, min_xz)
}
