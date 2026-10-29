package main

import "fmt"

func main() {
    var a, b int

    fmt.Print("Enter the first integer: ")
    fmt.Scan(&a)

    fmt.Print("Enter the second integer: ")
    fmt.Scan(&b)

    bothPositive := a > 0 && b > 0
    oneGreater := a > b || b > a
    notEqual := !(a == b)

    fmt.Println("Both positive:", bothPositive)
    fmt.Println("One greater than the other:", oneGreater)
    fmt.Println("Not equal:", notEqual)
}