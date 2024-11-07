# RC4 Encryption in Go

This project is a Go implementation of the RC4 encryption algorithm, which is a stream cipher that encrypts and decrypts data byte-by-byte. This implementation demonstrates key generation, encryption/decryption processes, and the ability to visualize encrypted data in a hexadecimal format.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [File Structure](#file-structure)
- [Contributing](#contributing)

## Overview

RC4 (Rivest Cipher 4) is a widely-used stream cipher designed by Ronald Rivest in 1987. The cipher is known for its simplicity and speed, making it suitable for encrypting data streams. However, it has known weaknesses and is no longer recommended for use in modern cryptographic systems.

This Go implementation includes:

- **RC4 Key Scheduling Algorithm (KSA)**: Initializes the state vector based on the provided key.
- **RC4 Pseudo-Random Generation Algorithm (PRGA)**: Generates the keystream used for encryption/decryption.
- **Encryption/Decryption**: Since RC4 is symmetric, the same function is used for both encryption and decryption.
- **Hexadecimal Visualization**: The encrypted data is printed in hexadecimal format for easier analysis and debugging.

## Features

- **Key Scheduling Algorithm (KSA)**: Initialize RC4 state with a user-defined key.
- **Encryption/Decryption**: Encrypt and decrypt data using the RC4 cipher.
- **Hexadecimal Output**: Visualize both original and encrypted data in a hexadecimal format.
- **Simple Interface**: Easy to use for encrypting or decrypting strings.

## Installation

### Prerequisites

- Go 1.16 or later installed on your machine. You can download and install Go from the official site: https://golang.org/dl/

### Steps to Install

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/rc4-encryption-go.git
    cd rc4-encryption-go
    ```

2. Build and run the program:

    ```bash
    go run main.go
    ```

This will run the example RC4 encryption and print the original, encrypted, and decrypted messages along with their hexadecimal representations.

## Usage

Once you've cloned and set up the project, you can modify the input data (`key` and `from`) in the `main.go` file to test encryption and decryption with your own values.

### Example

```go
key := "tomatoes" // The encryption key (8 bits to 2048 bits)
from := "Shall I compare thee to a summer's day?" // The plaintext message

// Initialize RC4 encryption
arcfour := rc4init(key, len(key))

// Encrypt the message
encrypted := rc4encrypt(arcfour, from)

// Output the encrypted result
fmt.Printf("Encrypted: '%s'\n", encrypted)
