package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {

	var text string

	fmt.Print("Enter text for hash: ")
	fmt.Scanf("%s", &text)

	hashResult := md5.Sum([]byte(text))

	hashResultString := hex.EncodeToString(hashResult[:])
	fmt.Println(hashResultString)
}
