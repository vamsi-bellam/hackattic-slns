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
		nums := strings.Split(scanner.Text(), " ")
		from, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("Error in parsing number")
		}
		to, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println("Error in parsing number")
		}

		for from <= to {
			if from%15 == 0 {
				fmt.Println("FizzBuzz")
			} else if from%3 == 0 {
				fmt.Println("Fizz")
			} else if from%5 == 0 {
				fmt.Println("Buzz")
			} else {
				fmt.Printf("%d\n", from)
			}
			from++
		}
	}
}
