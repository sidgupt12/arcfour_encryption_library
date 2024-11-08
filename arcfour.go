/* arcfour.go */
package main

import (
	"fmt"
	"os"
)

// used assert for only testing purposes, should be removed in production
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

func rc4byte(p *Arcfour) int {

	p.i = (p.i + 1) % 256
	p.j = (p.j + p.s[p.i]) % 256

	p.s[p.i], p.s[p.j] = p.s[p.j], p.s[p.i]

	temp := ((p.s[p.i]) + (p.s[p.j])) % 256
	p.k = int(p.s[temp])
	return p.k

}

func rc4encrypt(p *Arcfour, ct string, size int) string {
	encrypted := make([]byte, size)

	for i := 0; i < size; i++ {
		encrypted[i] = ct[i] ^ byte(rc4byte(p))
	}
	return string(encrypted)
}

func rc4decrypt(p *Arcfour, key string, keylen int) string {
	return rc4encrypt(p, key, keylen)
}

//this function was limited to only []byte input
// func printbin(input []byte, size int) {

// 	assert(size > 0, "size <= 0")

//		for i := 0; i < size; i++ {
//			if i%2 == 0 {
//				fmt.Printf(" ")
//			}
//			fmt.Printf("%02x", input[i])
//		}
//		fmt.Println()
//	}

// this function solved the problem of only []byte input
func printbin(input interface{}, size int) {
	assert(size > 0, "size <= 0")

	switch v := input.(type) {
	case string:
		// If the input is a string, treat it as []byte
		for i := 0; i < size; i++ {
			if i%2 == 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%02x", v[i]) // String characters are just bytes
		}
	case []int:
		// If the input is a []int8 (slice), treat it as bytes
		for i := 0; i < size; i++ {
			if i%2 == 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%02x", v[i]) // Print as hex values
		}
	default:
		panic("Unsupported input type")
	}
	fmt.Println()
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Text to be encrypted")
		os.Exit(1)
	}
	key := "tomatoes" // can be 8 bits to 2048 bits
	//from := "Shall I compare thee to a summer's day?"
	from := os.Args[1]

	skey := len(key)
	stext := len(from)

	fmt.Println("Initializing encryption...")
	os.Stdout.Sync()
	rc4 := rc4init(key, skey)
	fmt.Println("Done")

	fmt.Printf("Original text: '%s'\n ->", from)

	//encrypted = rc4encrypt(from, stext)

	// byteSlice := make([]byte, 256)
	// for i := 0; i < 256; i++ {
	// 	byteSlice[i] = byte(rc4.s[i]) // Convert int8 to byte
	// }

	fmt.Printf("key: %s\n", key)
	fmt.Print("key in binary: ")
	printbin(key, skey)

	// printbin(rc4.s[:], skey)

	encrypted := rc4encrypt(&rc4, from, stext)
	fmt.Printf("encrypted form : '%s'\n", encrypted)

	//we reinitialize the rc4 struct to decrypt because the state of the struct has changed
	rc4 = rc4init(key, skey)
	decrypted := rc4decrypt(&rc4, encrypted, stext)
	fmt.Printf("decrypted form : '%s'\n", decrypted)

}
