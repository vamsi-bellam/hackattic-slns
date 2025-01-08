package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/vamsi-bellam/hackattic-slns/classic/go/utils"
)

type Response struct {
	Include string `json:"include"`
}

type Solution struct {
	Files []string `json:"files"`
}

func toMD5Hash(tohash string) string {
	hashed := md5.Sum([]byte(tohash))
	return hex.EncodeToString(hashed[:])
}

func main() {
	utils.LoadEnv()

	problem := "collision_course"

	var response Response

	if err := json.Unmarshal(utils.FetchData(problem), &response); err != nil {
		fmt.Println("Error in parsing response to given format!")
		log.Fatal(err)
	}
	fmt.Println(response.Include)

	// Two different blocks(in hex form) with same MD5 hash value - Taken from https://www.mscs.dal.ca/~selinger/md5collision/
	block1 := "d131dd02c5e6eec4693d9a0698aff95c2fcab58712467eab4004583eb8fb7f8955ad340609f4b30283e488832571415a085125e8f7cdc99fd91dbdf280373c5bd8823e3156348f5bae6dacd436c919c6dd53e2b487da03fd02396306d248cda0e99f33420f577ee8ce54b67080a80d1ec69821bcb6a8839396f9652b6ff72a70"

	block2 := "d131dd02c5e6eec4693d9a0698aff95c2fcab50712467eab4004583eb8fb7f8955ad340609f4b30283e4888325f1415a085125e8f7cdc99fd91dbd7280373c5bd8823e3156348f5bae6dacd436c919c6dd53e23487da03fd02396306d248cda0e99f33420f577ee8ce54b67080280d1ec69821bcb6a8839396f965ab6ff72a70"

	decodedStringOfBlock1, err := hex.DecodeString(block1)
	if err != nil {
		fmt.Println("Error decoding block1 to string")
		log.Fatal(err)
	}

	decodedStringOfBlock2, err := hex.DecodeString(block2)
	if err != nil {
		fmt.Println("Error decoding block2 to string")
		log.Fatal(err)
	}

	file1 := string(decodedStringOfBlock1) + response.Include
	file2 := string(decodedStringOfBlock2) + response.Include

	fmt.Printf("The MD5 hashed values for file1 and file2 are - %s and %s\n",
		toMD5Hash(file1), toMD5Hash(file2))

	solution := Solution{
		[]string{
			base64.StdEncoding.EncodeToString([]byte(file1)),
			base64.StdEncoding.EncodeToString([]byte(file2))}}

	sol, err := json.Marshal(solution)

	if err != nil {
		fmt.Println("Error in encoding to json!")
		log.Fatal(err)
	}

	utils.SubmitSolution(problem, sol)
}
