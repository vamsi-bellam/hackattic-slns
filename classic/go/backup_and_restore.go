package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

type Response struct {
	Dump string `json:"dump"`
}

type Solution struct {
	AliveSsns []string `json:"alive_ssns"`
}

func fetchData(problem string) []byte {
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

func sendData(problem string, solution []byte) {

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
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file!")
		log.Fatal(err)
	}

	problem := "backup_restore"

	var response Response

	if err := json.Unmarshal(fetchData(problem), &response); err != nil {
		fmt.Println("Error in parsing response to given format!")
		log.Fatal(err)
	}

	fmt.Println("Decoding the fetched data...")

	decodedDump, err := base64.StdEncoding.DecodeString(response.Dump)
	if err != nil {
		fmt.Println("Error in decoding dump!")
		log.Fatal(err)
	}

	fmt.Println("Decoded the required data!")

	gzipReader, err := gzip.NewReader(bytes.NewReader(decodedDump))
	if err != nil {
		fmt.Println("Error while decompressing data!")
		log.Fatal(err)
	}

	fmt.Println("Decompressing the decoded data...")

	decompressedData, err := io.ReadAll(gzipReader)

	if err != nil {
		fmt.Println("Error while reading decompressed data!")
		log.Fatal(err)
	}

	fmt.Println("Decompressed the decoded data!")

	fmt.Println("Dumping the data to a file db_dump..")

	err = os.WriteFile("db_dump", decompressedData, 0644)

	if err != nil {
		fmt.Println("Error while writing to file!")
		log.Fatal(err)
	}
	fmt.Println("Dumped the data to a file db_dump!")

	fmt.Println("Creating database to restore backup data...")

	_, err = exec.Command("sh", "-c",
		`psql -U postgres -c "create database ssns"`).Output()
	if err != nil {
		fmt.Println("Error while creating to database!")
		log.Fatal(err)
	}
	fmt.Println("Created database to restore backup data!")

	fmt.Println("Restoring data to db...")

	_, err = exec.Command("sh", "-c",
		"psql -U postgres -d ssns < db_dump").Output()
	if err != nil {
		fmt.Println("Error while restore to database!")
		log.Fatal(err)
	}

	fmt.Println("Restored data to db!")

	fmt.Println("Fetching alived ssns...")

	output, err := exec.Command("sh", "-c",
		`psql -U postgres -d ssns -c " select ssn from criminal_records where status = 'alive' "`).Output()

	if err != nil {
		fmt.Println("Error while getting ssn records!")
		log.Fatal(err)
	}

	fmt.Println("Fetced alived ssns!")

	fmt.Println("Dropping the created database...")

	_, err = exec.Command("sh", "-c",
		`psql -U postgres -c "drop database ssns"`).Output()
	if err != nil {
		fmt.Println("Error while dropping database!")
		log.Fatal(err)
	}
	fmt.Println("Dropped the created database!")

	fmt.Println("Deleting the inter file db_dump...")

	err = os.Remove("db_dump")
	if err != nil {
		fmt.Println("Error in deleting file!")
		log.Fatal(err)
	}
	fmt.Println("Deleted the inter file db_dump!")

	records := strings.Split(string(output), "\n")
	records = records[2:(len(records) - 3)]

	var solution Solution

	for _, record := range records {
		solution.AliveSsns = append(solution.AliveSsns, strings.TrimSpace(record))
	}

	sol, err := json.Marshal(solution)
	if err != nil {
		fmt.Println("Error in encoding to json!")
		log.Fatal(err)
	}

	sendData(problem, sol)
}
