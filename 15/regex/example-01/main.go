package main

import (
	"fmt"
	"regexp"
)

func main() {
	testString1 := "john.doe@somehost.com"
	testString2 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa!"
	regex := regexp.MustCompile("^([a-zA-Z0-9])(([\\-.]|[_]+)?([a-zA-Z0-9]+))*(@){1}[a-z0-9]+[.]{1}(([a-z]{2,3})|([a-z]{2,3}[.]{1}[a-z]{2,3}))$")

	fmt.Println(regex.MatchString(testString1))
	// expected output: true
	fmt.Println(regex.MatchString(testString2))
	// expected output: false
}
