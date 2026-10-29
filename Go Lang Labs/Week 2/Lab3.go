
package main

import "fmt"

func runLab3(n1 int, n2 int) {
	fmt.Println("=== Lab 3 ===")

	myAnd(n1, n2)
	myOr(n1, n2)
	myXor(n1, n2)
	myNot(n1, n2)
	myLeftShift(n1, n2)
	myRightShift(n1, n2)

	fmt.Println("\n=== Assignment Operators ===")

	x := n1
	x &= n2
	fmt.Println("AND assignment:", x)

	x = n1
	x |= n2
	fmt.Println("OR assignment:", x)

	x = n1
	x ^= n2
	fmt.Println("XOR assignment:", x)

	x = n1
	x <<= uint(n2)
	fmt.Println("Left shift assignment:", x)

	x = n1
	x >>= uint(n2)
	fmt.Println("Right shift assignment:", x)
}

func myAnd(n1 int, n2 int) {
	fmt.Printf("%d AND %d = %b\n", n1, n2, n1&n2)
}

func myOr(n1 int, n2 int) {
	fmt.Printf("%d OR %d = %b\n", n1, n2, n1|n2)
}

func myXor(n1 int, n2 int) {
	fmt.Printf("%d XOR %d = %b\n", n1, n2, n1^n2)
}

func myNot(n1 int, n2 int) {
	fmt.Printf("NOT %d = %b\n", n1, ^n1)
	fmt.Printf("NOT %d = %b\n", n2, ^n2)
}

func myLeftShift(n1 int, n2 int) {
	fmt.Printf("%d << %d = %b\n", n1, n2, n1<<uint(n2))
}

func myRightShift(n1 int, n2 int) {
	fmt.Printf("%d >> %d = %b\n", n1, n2, n1>>uint(n2))
}

func main() {
	var n1, n2 int

	fmt.Print("Enter two non-negative integers: ")
	fmt.Scan(&n1, &n2)

	runLab3(n1, n2)
}
