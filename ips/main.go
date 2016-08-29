package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// http://stackoverflow.com/questions/10006459/regular-expression-for-ip-address-validation
var IP_REGEX = regexp.MustCompile(`(([1-9]?\d|1\d\d|2[0-5][0-5]|2[0-4]\d)\.){3}(1\d\d|2[0-5][0-5]|2[0-4]\d|[1-9]?\d)`)

func Ips(input string) []string {
	return IP_REGEX.FindAllString(input, -1)
}

func main() {
	var results = []string{}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		input := s.Text()
		matches := Ips(input)
		results = append(results, matches...)
	}

	for _, ip := range results {
		fmt.Println(ip)
	}
}
