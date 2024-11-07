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
	s       [256]int
	i, j, k int
}

func rc4init(key string, size int) Arcfour {
	var p Arcfour
	for x := 0; x < 256; x++ {
		p.s[x] = 0
	}
	p.i, p.j, p.k = 0, 0, 0

	var temp1, temp2 int

	for p.i = 0; p.i < 256; p.i++ {
		p.s[p.i] = p.i
	}

	for p.i = 0; p.i < 256; p.i++ {
		temp1 = p.i % size
		temp2 = p.j + p.s[p.i] + int(key[temp1])
		p.j = temp2 % 256
		p.s[p.i], p.s[p.j] = p.s[p.j], p.s[p.i]
	}
	p.i, p.j = 0, 0

	return p

}

func rc4byte() int {
	return 0
}

func rc4encrypt(key string, keylen int) string {
	return ""
}

func rc4decrypt(key string, keylen int) string {
	return rc4encrypt(key, keylen)
}

func printbin(input string, size int) {

	assert(size > 0, "size <= 0")

	for i := 0; i < size; i++ {
		if i%2 == 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%02x", input[i])
	}
	fmt.Println()
}

func main() {

	//Arcfour * rc4

	var encrypted, decrypted string
	fmt.Println(encrypted, decrypted)

	key := "tomatoes" // can be 8 bits to 2048 bits
	from := "Shall I compare thee to a summer's day?"

	skey := len(key)
	//stext := len(from)

	fmt.Println("Initializing encryption...")
	os.Stdout.Sync()
	//rc4 = rc4init(key, skey)
	fmt.Println("Done")

	fmt.Printf("'%s'\n ->", from)
	//encrypted = rc4encrypt(from, stext)

	printbin(key, skey)

}
