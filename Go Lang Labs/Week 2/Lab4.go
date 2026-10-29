package main

import "fmt"

func calculate(a, b int, choice int) {
    switch choice {
    case 1:
        fmt.Println("Result:", a+b)

    case 2:
        fmt.Println("Result:", a-b)

    case 3:
        fmt.Println("Result:", a*b)

    case 4:
        if b == 0 {
            fmt.Println("Error: division by zero")
        } else {
            fmt.Println("Result:", float64(a)/float64(b))
        }

    case 5:
        if b == 0 {
            fmt.Println("Error: modulo by zero")
        } else {
            fmt.Println("Result:", a%b)
        }

    default:
        fmt.Println("Invalid operator")
    }
}

func main() {
    for {
        fmt.Println("\n===== Mini Calculator =====")
        fmt.Println("1) Add  2) Sub  3) Mul")
        fmt.Println("4) Div  5) Mod  6) Exit")
        fmt.Print("Choose: ")

        var choice int
        if _, err := fmt.Scan(&choice); err != nil {
            fmt.Println("Invalid input")
            return
        }

        if choice == 6 {
            fmt.Println("Goodbye!")
            return
        }

        if choice < 1 || choice > 5 {
            fmt.Println("Invalid operator")
            continue
        }

        var a, b int

        fmt.Print("Enter a: ")
        if _, err := fmt.Scan(&a); err != nil {
            fmt.Println("Invalid input")
            return
        }

        fmt.Print("Enter b: ")
        if _, err := fmt.Scan(&b); err != nil {
            fmt.Println("Invalid input")
            return
        }

        calculate(a, b, choice)
    }
}