package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		sum := int64(0)

		for _, num := range nums {
			val, err := strconv.ParseInt(num, 0, 64)
			if err != nil {
				// if not any number system, then it is single ASCII char
				val = int64(num[0])
			}
			sum += val
		}
		fmt.Println(sum)
	}
}
