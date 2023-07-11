package main

import (
	"fmt"
	"regexp"
)

func main() {
	testString1 := "<h1>Go Secure Coding Practices Guide</h1>"
	testString2 := "<p>Go Secure Coding Practices Guide</p>"
	testString3 := "<h1>Go Secure Coding Practices Guid</p>"
	// todo: Change to the following to 'fix' unknown escape sequence error:
	// <([a-z][a-z0-9]*)\b[^>]*>.*?<\\/\\1>
	regex := regexp.MustCompile("<([a-z][a-z0-9]*)\b[^>]*>.*?<\/\1>")

	fmt.Println(regex.MatchString(testString1))
	fmt.Println(regex.MatchString(testString2))
	fmt.Println(regex.MatchString(testString3))
}
