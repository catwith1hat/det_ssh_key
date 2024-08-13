package main

import (
	"crypto/ed25519"
	"crypto/sha512"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mikesmitty/edkey"
	"golang.org/x/crypto/ssh"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: det_key <seed> <private_key_file> <public_key_file>")
		os.Exit(1)
	}

	seedFile := os.Args[1]
	privateKeyFile := os.Args[2]
	publicKeyFile := os.Args[3]

	data, err := ioutil.ReadFile(seedFile)
	if err != nil {
		fmt.Println("Error reading seed key file:", err)
		os.Exit(1)
	}

	// Hash the private key data using SHA-512
	hashedData := sha512.Sum512(data)

	// Clamp output as specified by the Curve25519 paper.
	hashedData[0] &= 0xF8
	hashedData[31] &= 0x3F
	hashedData[31] |= 0x40

	privateKey := ed25519.PrivateKey(hashedData[:])

	publicKey := privateKey.Public().(ed25519.PublicKey)

	pemBlock := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: edkey.MarshalED25519PrivateKey(privateKey),
	}

	pemEncoded := pem.EncodeToMemory(pemBlock)

	// Write the PEM private key to a file
	err = ioutil.WriteFile(privateKeyFile, pemEncoded, 0600) // 0600 for read/write permissions only for the owner
	if err != nil {
		fmt.Println("Error writing PEM private key file:", err)
		os.Exit(1)
	}

	fmt.Println("PEM private key written to:", privateKeyFile)

	sshPublicKey, err := ssh.NewPublicKey(publicKey)
	if err != nil {
		fmt.Println("Error creating SSH public key:", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(publicKeyFile, ssh.MarshalAuthorizedKey(sshPublicKey), 0644)
	if err != nil {
		fmt.Println("Error writing public key file:", err)
		os.Exit(1)
	}

	fmt.Println("Public key written to:", publicKeyFile)
}
