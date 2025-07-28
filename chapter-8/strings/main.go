package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello, World!", "o"))
	fmt.Println(strings.Count("Hello, World!", "o"))
	fmt.Println(strings.Count("Hello, World!", "O")) // case-sensitive
	fmt.Println(strings.Index("Hello, World!", "o"))
	fmt.Println(strings.HasPrefix("Hello, World!", "He"))
	fmt.Println(strings.HasSuffix("Hello, World!", "!"))
	fmt.Println(strings.ToLower("Hello, World!"))
	fmt.Println(strings.CutPrefix("Hello, World!", "Hello"))
	fmt.Println(strings.Join([]string{"Hello", "World"}, ", "))
	fmt.Println(strings.Split("Hello, World!", ", "))
	fmt.Println(strings.TrimSpace("   Hello, World!   "))
	fmt.Println(strings.ReplaceAll("Hello, World!", "o", "0"))
	fmt.Println(strings.Repeat("Hello ", 3))
	fmt.Println(strings.ToUpper("Hello, World!"))
	fmt.Println(strings.Fields("Hello, World! This is a test."))
	fmt.Println(strings.Trim("%%%Hello, World!%%%", "%"))
	fmt.Println(strings.Trim("   Hello, World!   ", " "))
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	fmt.Println(strings.Map(func(r rune) rune {
		if r == 'a' {
			return 'b'
		} else if r == 'b' {
			return 'c'
		}
		return r
	}, "aaba"))
}
