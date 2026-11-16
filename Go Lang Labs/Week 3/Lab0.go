package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/sha3"
)

func compareHashes(name, hashA, hashB string) {
	status := "No Match!"
	if hashA == hashB {
		status = "Match!"
	}
	fmt.Printf("Hash (%s):\n", name)
	fmt.Printf("  Output A = %s\n", hashA)
	fmt.Printf("  Output B = %s\n", hashB)
	fmt.Printf("  => %s\n", status)
}

func proofMe(txt1, txt2 string) {
	compareHashes("MD5", fmt.Sprintf("%x", md5.Sum([]byte(txt1))), fmt.Sprintf("%x", md5.Sum([]byte(txt2))))
	compareHashes("SHA1", fmt.Sprintf("%x", sha1.Sum([]byte(txt1))), fmt.Sprintf("%x", sha1.Sum([]byte(txt2))))
	compareHashes("SHA256", fmt.Sprintf("%x", sha256.Sum256([]byte(txt1))), fmt.Sprintf("%x", sha256.Sum256([]byte(txt2))))
	compareHashes("SHA512", fmt.Sprintf("%x", sha512.Sum512([]byte(txt1))), fmt.Sprintf("%x", sha512.Sum512([]byte(txt2))))
	compareHashes("SHA3-256", fmt.Sprintf("%x", sha3.Sum256([]byte(txt1))), fmt.Sprintf("%x", sha3.Sum256([]byte(txt2))))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("======== Hashing Program ========")
	fmt.Print("Please input value 1: ")
	txt1, _ := reader.ReadString('\n')
	txt1 = strings.TrimSpace(txt1)

	fmt.Print("Please input value 2: ")
	txt2, _ := reader.ReadString('\n')
	txt2 = strings.TrimSpace(txt2)

	proofMe(txt1, txt2)
}