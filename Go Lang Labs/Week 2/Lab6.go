package main

import (
    "bufio"
    "encoding/hex"
    "fmt"
    "os"
    "strings"
)

func xorEncrypt(text string, key byte) string {
    data := []byte(text)

    for i := range data {
        data[i] ^= key
    }

    return string(data)
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter plaintext: ")
    plaintext, err := reader.ReadString('\n')
    if err != nil && len(plaintext) == 0 {
        fmt.Println("Error reading plaintext:", err)
        return
    }
    plaintext = strings.TrimRight(plaintext, "\r\n")

    fmt.Print("Enter a single-byte key (0-255): ")
    var key int

    if _, err := fmt.Scan(&key); err != nil || key < 0 || key > 255 {
        fmt.Println("Invalid key: enter an integer from 0 to 255")
        return
    }

    // Encrypt the plaintext.
    ciphertext := xorEncrypt(plaintext, byte(key))

    // Display ciphertext as hexadecimal so all bytes are visible.
    fmt.Println("Ciphertext (hex):", hex.EncodeToString([]byte(ciphertext)))

    // Decrypt using the same XOR function and key.
    decrypted := xorEncrypt(ciphertext, byte(key))

    fmt.Println("Decrypted text:", decrypted)
    fmt.Println("Original recovered:", plaintext == decrypted)
}