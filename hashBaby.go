package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

//start main OMIT

func GetKeccak256Hash(str string) {
	h := crypto.Keccak256Hash([]byte(str))
	fmt.Println(h.Hex())
}

func main() {
	GetKeccak256Hash("Don’t wait. Life goes faster than you think!")
	GetKeccak256Hash("Don’t wait. Life goes faster than you think.")
}

//end main OMIT
