package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type AccountData map[string]map[string]int
type Person struct {
	name    string
	balance int
}

func formattedBalance(balance int) string {
	isNeg := false
	if balance < 0 {
		isNeg = true
		balance = -1 * balance
	}

	s := strconv.Itoa(balance)
	l := len(s)
	if l < 3 {
		return s
	}

	nums := strings.Split(s, "")
	finalString := ""
	count := 0
	for i := len(nums) - 1; i >= 0; i-- {
		finalString = nums[i] + finalString
		count++
		if count%3 == 0 && i != 0 {
			finalString = "," + finalString
		}
	}
	if isNeg {
		finalString = "-" + finalString
	}
	return finalString
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	data := make([]Person, 0)

	for scanner.Scan() {
		line := scanner.Text()
		var record AccountData
		err := json.Unmarshal([]byte(line), &record)
		if err != nil {
			fmt.Println("Error in parsing json!")
		}
		d := Person{"", 0}

		for k, v := range record {
			if k != "extra" {
				bal, ok := v["balance"]
				if !ok {
					fmt.Println("No balance field")
				}
				(&d).name = k
				(&d).balance = bal
			}
		}

		extra, ok := record["extra"]
		if ok {
			bal, ok := extra["balance"]
			if ok {
				(&d).balance = bal
			}
		}
		data = append(data, d)
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].balance < data[j].balance
	})

	for _, person := range data {
		fmt.Printf("%s: %s\n", person.name, formattedBalance(person.balance))
	}
}
