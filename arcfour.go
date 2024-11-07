/* arcfour.go */
package main

import (
	"fmt"
	"os"
)

func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

type Arcfour struct {
	// ...
}

func rc4init(key string, keylen int) Arcfour {
	// ...
}

func rc4byte() int {

}

func rc4encrypt(key string, keylen int) string {

}

func printbin(input string, size int) {

	assert(size > 0, "size <= 0")

	for i := 0; i < size; i++ {
		if (i+1)%2 == 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%02x", input[i])
	}
	fmt.Println()
}

func main() {

	//Arcfour * rc4

	var encrypted, decrypted string

	key := "tomatoes" // can be 8 bits to 2048 bits
	from := "Shall I compare thee to a summer's day?"

	skey := len(key)
	stext := len(from)

	fmt.Println("Initializing encryption...")
	os.Stdout.Sync()
	//rc4 = rc4init(key, skey)
	fmt.Println("Done")

	fmt.Printf("'%s'\n ->", from)
	//encrypted = rc4encrypt(from, stext)

	printbin(key, skey)

}
