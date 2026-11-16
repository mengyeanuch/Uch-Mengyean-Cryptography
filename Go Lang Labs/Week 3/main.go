package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"golang-labs/week3/utils/crack"
)

func main() {
	_, sourceFile, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Error locating the Week 3 directory")
		return
	}

	week3Dir := filepath.Dir(sourceFile)
	wordlist := filepath.Join(week3Dir, "..", "nord_vpn.txt")
	logPath := filepath.Join(week3Dir, "verbose.txt")

	logFile, err := os.Create(logPath)
	if err != nil {
		fmt.Println("Error creating verbose.txt:", err)
		return
	}
	defer logFile.Close()

	// Lab #1: MD5
	md5Target := "6a85dfd77d9cb35770c9dc6728d73d3f"
	if pass := crack.CrackMD5(wordlist, md5Target, logFile); pass != "" {
		fmt.Printf("[+] Lab #1 Cracked (MD5): %s\n", pass)
	}

	// Lab #2: SHA1
	sha1Target := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
	if pass := crack.CrackSHA1(wordlist, sha1Target, logFile); pass != "" {
		fmt.Printf("[+] Lab #2 Cracked (SHA1): %s\n", pass)
	}

	// Lab #3: SHA512
	sha512Target := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
	if pass := crack.CrackSHA512(wordlist, sha512Target, logFile); pass != "" {
		fmt.Printf("[+] Lab #3 Cracked (SHA512): %s\n", pass)
	}
}
