package main

import (
	"crypto/rand"
	"fmt"
)

// Returns a random 15 character string from the collection
func GeneratePassword() string {

	var collection = []byte("0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM~!@#$%^&*(()-?/")

	pwdLength := 15
	pwd := make([]byte, pwdLength)
	rand.Read(pwd)

	for i := 0; i < pwdLength; i++ {
		pwd[i] = collection[int(pwd[i]%byte(len(collection)))]
	}

	return string(pwd)
}

// TODO:: find how to implement
//
// Copies generated string to clipboard
func CopyToClipboard() {
	fmt.Println(`
    *************************
       copied to clipboard
    `)
}
