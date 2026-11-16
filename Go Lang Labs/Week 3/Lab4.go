package main

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"regexp"
)

func main() {
	// 1. Leaked Google search URL from the slide
	leakedURL := "https://www.google.com/search?q=%62%36%63%63%62%34%65%63%65%35%34%35%34%64%63%61%65%35%31%37%37%38%62%33%65%32%33%39%65%62%63%32&rlz=1C5CHF"

	// 2. Parse and URL-decode the 'q' parameter to reveal the MD5 hash
	parsedURL, _ := url.Parse(leakedURL)
	targetHash := parsedURL.Query().Get("q")
	fmt.Println("Decoded MD5 Hash:", targetHash)

	// 3. Verify the candidate password ("meowmeow") against the MD5 hash and regex
	candidate := "meowmeow"
	computedHash := fmt.Sprintf("%x", md5.Sum([]byte(candidate)))

	if computedHash == targetHash {
		flag := fmt.Sprintf("cryptoCTF{%s}", candidate)
		regex := regexp.MustCompile(`^cryptoCTF\{(?:\x6d\x65\x6f\x77){2}\}$`)

		if regex.MatchString(flag) {
			fmt.Println("[+] Valid Flag:", flag)
		}
	}
}