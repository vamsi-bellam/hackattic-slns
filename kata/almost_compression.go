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
		final := ""

		count := 1

		for i := 1; i < dlen; i++ {
			if data[i] != data[i-1] {
				if count <= 2 {
					for count > 0 {
						final += data[i-1]
						count--
					}
				} else {
					final += fmt.Sprintf("%d%s", count, data[i-1])
				}
				count = 1
			} else {
				count++
			}
		}

		if count <= 2 {
			for count > 0 {
				final += data[dlen-1]
				count--
			}
		} else {
			final += fmt.Sprintf("%d%s", count, data[dlen-1])
		}

		fmt.Println(final)
	}
}
