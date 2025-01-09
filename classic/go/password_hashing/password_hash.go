package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/vamsi-bellam/hackattic-slns/classic/go/utils"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

type Pbkdf2 struct {
	Rounds int    `json:"rounds"`
	Hash   string `json:"hash"`
}

type Scrypt struct {
	N       int    `json:"N"`
	R       int    `json:"r"`
	P       int    `json:"p"`
	Buflen  int    `json:"buflen"`
	Control string `json:"_control"`
}

type Response struct {
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Pbkdf2   Pbkdf2 `json:"pbkdf2"`
	Scrypt   Scrypt `json:"scrypt"`
}

type Solution struct {
	Sha256 string `json:"sha256"`
	Hmac   string `json:"hmac"`
	Pbkdf2 string `json:"pbkdf2"`
	Scrypt string `json:"scrypt"`
}

func main() {

	utils.LoadEnv()
	problem := "password_hashing"

	var response Response

	if err := json.Unmarshal(utils.FetchData(problem), &response); err != nil {
		fmt.Println("Error in parsing response to given format!")
		log.Fatal(err)
	}

	fmt.Println(response)

	// Used to understand details of Sha256 - https://www.simplilearn.com/tutorials/cyber-security-tutorial/sha-256-algorithm
	sha := sha256.Sum256([]byte(response.Password))
	// %x is string of hex value which is in bytes in sha
	fmt.Printf("Sha256 of %s is %x\n", response.Password, sha)

	salt, err := base64.StdEncoding.DecodeString(response.Salt)

	if err != nil {
		fmt.Println("Error in decoding salt!")
		log.Fatal(err)
	}

	fmt.Printf("Decoded Salt : %x\n", salt)

	hm := hmac.New(sha256.New, salt)
	// Hmm, there is difference b/w Write vs Sum - https://groups.google.com/g/golang-nuts/c/S19r1lGNu_A
	// Write - adds passed data to running hash(it includes passed data to running hash and updates hash)
	// Where as Sum appends the current hash to passed data, so it doesn't change the underlying hash state
	// So, using Sum with password won't work here!!
	_, err = hm.Write([]byte(response.Password))
	if err != nil {
		fmt.Println("Error in writing password to hash!")
		log.Fatal(err)
	}

	hmacSha := hm.Sum(nil)
	fmt.Printf("HAMC of %s is %x\n", response.Password, hmacSha)

	pbkdfSha := pbkdf2.Key([]byte(response.Password), salt,
		response.Pbkdf2.Rounds, sha256.Size, sha256.New)
	fmt.Printf("pbkdf of %s is %x\n", response.Password, pbkdfSha)

	scryptSha, err := scrypt.Key([]byte(response.Password), salt, response.Scrypt.N,
		response.Scrypt.R, response.Scrypt.P, response.Scrypt.Buflen)
	if err != nil {
		fmt.Println("Error in encrypting using scrypt!")
		log.Fatal(err)
	}

	fmt.Printf("scrypt of %s is %x\n", response.Password, scryptSha)

	solution := Solution{
		Sha256: hex.EncodeToString(sha[:]),
		Hmac:   hex.EncodeToString(hmacSha),
		Pbkdf2: hex.EncodeToString(pbkdfSha),
		Scrypt: hex.EncodeToString(scryptSha)}

	fmt.Println(solution)

	sol, err := json.Marshal(solution)
	if err != nil {
		fmt.Println("Error in encoding to json!")
		log.Fatal(err)
	}

	utils.SubmitSolution(problem, sol)
}
