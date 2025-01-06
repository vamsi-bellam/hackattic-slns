package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		dec, err := base64.StdEncoding.DecodeString(line)
		if err != nil {
			fmt.Println("Error in decoding string!")
		}

		fmt.Println(string(dec))
	}
}
