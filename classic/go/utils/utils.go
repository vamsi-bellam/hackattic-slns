package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file!")
		log.Fatal(err)
	}
}

func FetchData(problem string) []byte {
	fmt.Println("Fetching the required data...")

	url := fmt.Sprintf(
		"https://hackattic.com/challenges/%s/problem?access_token=%s",
		problem, os.Getenv("ACCESS_TOKEN"))
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error in getting problem data!")
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error while reading response!")
		log.Fatal(err)
	}

	fmt.Println("Fetched the required data!")

	return data
}

func SubmitSolution(problem string, solution []byte) {

	fmt.Println("Sending solution...!")

	url := fmt.Sprintf(
		"https://hackattic.com/challenges/%s/solve?access_token=%s",
		problem, os.Getenv("ACCESS_TOKEN"))
	resp, err := http.Post(url, "json", bytes.NewBuffer(solution))

	if err != nil {
		fmt.Println("Error sending solution data!")
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatal("Error")
	}

	fmt.Println("Solution submitted succesfully!")

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solution Response - %s", string(result))
}
