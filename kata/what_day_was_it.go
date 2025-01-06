package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		days, err := strconv.Atoi(line)

		if err != nil {
			fmt.Println("Error parsing the number!")
		}

		extraDays := days % 7

		if days < 0 {
			days = -1 * days
			extraDays = 7 - days%7
		}

		day := "Thursday"
		switch extraDays {
		case 1:
			day = "Friday"
		case 2:
			day = "Saturday"
		case 3:
			day = "Sunday"
		case 4:
			day = "Monday"
		case 5:
			day = "Tuesday"
		case 6:
			day = "Wednesday"
		}
		fmt.Println(day)
	}
}
