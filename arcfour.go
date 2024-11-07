/* arcfour.go */
package main

import (
	"fmt"
	"os"
)

type s_arcfour struct {
	// ...
}

func rc4init(key string, keylen int) Arcfour {
	// ...
}

func rc4byte() int {

}

func rc4encrypt(key string, keylen int) string {

}

func main() {

	Arcfour * rc4

	var skey, stext int
	var key, from, encrypted, decrypted string

	key = "tomatoes" // can be 8 bits to 2048 bits
	from = "Shall I compare thee to a summer's day?"

	skey = len(key)
	stext = len(from)

	fmt.Println("Initializing encryption...")
	os.Stdout.Sync()
	rc4 = rc4init(key, skey)
	fmt.Println("Done")

	fmt.Printf("'%s'\n ->", from)
	encrypted = rc4encrypt(from, stext)

}
