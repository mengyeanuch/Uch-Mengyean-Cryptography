package main

import (
    "bufio"
    "encoding/base64"
    "encoding/hex"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter a string: ")
    input, err := reader.ReadString('\n')
    if err != nil && len(input) == 0 {
        fmt.Println("Error reading input:", err)
        return
    }

    input = strings.TrimRight(input, "\r\n")
    data := []byte(input)

    fmt.Println("Original text:", input)

    fmt.Print("Binary: ")
    for _, b := range data {
        fmt.Printf("%08b ", b)
    }
    fmt.Println()

    fmt.Println("Hexadecimal:", hex.EncodeToString(data))
    fmt.Println("Base64:", base64.StdEncoding.EncodeToString(data))
}