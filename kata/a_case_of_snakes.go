package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isStartsWithType(s string) bool {
	types := []string{"p", "d", "f", "b", "w", "ch", "fn", "u16", "u32", "u64", "i64", "i16", "i32", "sz"}
	for _, typ := range types {
		if strings.HasPrefix(s, typ) {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile("[A-Z][a-z0-9]*")
		strs := re.FindAllString(line, -1)

		if !isStartsWithType(line) {
			firstWord := ""
			for i := 0; i < len(line); i++ {
				if line[i] >= 65 && line[i] <= 90 {
					break
				}
				firstWord += string(line[i])
			}
			if len(firstWord) > 0 {
				strs = append([]string{firstWord}, strs...)
			}
		}

		for i, v := range strs {
			strs[i] = strings.ToLower(v)
		}
		fmt.Println(strings.Join(strs, "_"))
	}
}
