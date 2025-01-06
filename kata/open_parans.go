package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "")
		dlen := len(data)

		if dlen%2 != 0 {
			fmt.Println("no")
			continue
		}

		stack := make([]string, dlen/2)
		top := -1
		isValid := true

		for _, c := range data {
			if c == "(" {
				top++
				stack[top] = "("
			} else {
				if top == -1 || stack[top] != "(" {
					isValid = false
					break
				} else {
					top--
				}
			}
		}

		if !isValid || top > -1 {
			fmt.Println("no")
		} else {
			fmt.Println("yes")
		}
	}
}
