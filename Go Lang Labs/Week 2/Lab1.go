package main
import "fmt"

func main() {
	var a, b int

    fmt.Print("Enter the first integer: ")
    fmt.Scan(&a)

    fmt.Print("Enter the second integer: ")
    fmt.Scan(&b)

    // Assignment operator =
    x := a
    x = b
    fmt.Println("After = :", x)

    // Addition assignment operator +=
    x = a
    x += b
   fmt.Printf("%d + %d = %d\n",a,b,x)

    // Subtraction assignment operator -=
    x = a
    x -= b
    fmt.Printf("%d - %d = %d\n",a,b,x)

    // Multiplication assignment operator *=
    x = a
    x *= b
    fmt.Printf("%d * %d = %d\n",a,b,x)

    // Division assignment operator /=
    if b != 0 {
        x = a
        x /= b
        fmt.Printf("%d / %d = %d\n",a,b,x)

        // Modulus assignment operator %=
        x = a
        x %= b
        fmt.Printf("%d % %d = %d\n",a,b,x)
    } else {
        fmt.Println("Division and modulus by zero are not allowed.")
    }
	
}