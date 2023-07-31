package main

import (
	"fmt"
	"regexp"
	"strings"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile("something")
}

func regexCompiledMatch(text string) bool {
	return re.Match([]byte(text))
}
func regexMatch(pattern, text string) bool {
	match, _ := regexp.MatchString(pattern, text)
	return match
}
func stringContainsMatch(pattern, text string) bool {
	return strings.Contains(text, pattern)
}
func main() {
	pattern := "example"
	text := "This is an example text"
	regexResult := regexMatch(pattern, text)
	containsResult := stringContainsMatch(pattern, text)
	fmt.Printf("Regex Match: %v\n", regexResult)
	fmt.Printf("String Contains Match: %v\n", containsResult)
	fmt.Printf("Re compiled Match: %v\n", regexCompiledMatch((text)))
}
