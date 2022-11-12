package main


import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"
)

	

func main() {
	plaintext, err := ioutil.ReadFile("/volumes/v1/output.txt")
	//fmt.Println(string(plaintext))
	keyR, err := ioutil.ReadFile("/work/secrets/keyR.bin")
	keyG, err := ioutil.ReadFile("/work/secrets/keyG.bin")
	keyB, err := ioutil.ReadFile("/work/secrets/keyB.bin")

	var key []byte
	for i:=0; i<32; i++ {
		key = append(key, keyR[i] ^ keyG[i] ^ keyB[i])
	}

	fmt.Println("KeyR: ", keyR)
	fmt.Println("KeyG: ", keyG)
	fmt.Println("KeyB: ", keyB)

	

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encrypting with xor of all keys: ", key)

	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Panic(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}


	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	fmt.Println("NONCE: ", nonce)

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	
	// Save back to file
	err = ioutil.WriteFile("/work/out/file1.bin", ciphertext, 0777)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Decrypting . . .")
	ciphertext, err = ioutil.ReadFile("/work/out/file1.bin")
	if err != nil {
		log.Fatal(err)
	}

	nonce = ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]

	
	text, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Decrypted content: \n", string(text))
}
