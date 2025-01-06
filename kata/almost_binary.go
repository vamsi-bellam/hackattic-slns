package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "")
		dlen := len(data)
		sum := 0

		for i, v := range data {
			if v == "#" {
				sum += int(math.Pow(float64(2), float64(dlen-i-1)))
			}
		}
		fmt.Println(sum)

	}
}
