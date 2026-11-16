package crack

import (
	"bufio"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
)

func CrackSHA512(wordlistPath, targetHash string, logFile *os.File) string {
	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println("Error opening wordlist:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		hash := fmt.Sprintf("%x", sha512.Sum512([]byte(word)))
		fmt.Fprintf(logFile, "[SHA512] Trying: %s -> %s\n", word, hash)

		if hash == strings.ToLower(targetHash) {
			return word
		}
	}
	return ""
}