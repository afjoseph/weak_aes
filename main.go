package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"time"

	"github.com/afjoseph/commongo/print"
	"github.com/afjoseph/weakaes/bruteforcer"
)

func main() {
	print.SetLevel(print.LOG_DEBUG)

	// - Encrypt
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	plaintext := []byte("bunnyfoofoo")
	// -- Make the cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	// -- Make a nonce
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	// -- Make a new GCM
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	// -- Seal
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	print.Infof("plaintext: %s\n", string(plaintext))
	print.Infof("ciphertext: %x\n", ciphertext)
	print.Infof("nonce: %x\n", nonce)
	print.Infof("real key: %x\n", key)

	// - Bruteforce decryption
	// XXX We're the attackers here and we're assuming the following:
	// * We got the nonce
	// * We got 29 bytes out of the 32 bytes of the key
	// * We got the ciphertext
	ciphertext, _ = hex.DecodeString("8103e755f7e7bf50807df25b485b7bd8c3627e6745feb38f2cd075")
	nonce, _ = hex.DecodeString("c44ac1deb893353852c570ab")
	start := time.Now()
	if !bruteforcer.Run(key, uint32(len(key))-3, func(potentialKey []uint8) bool {
		// -- Make a new cipher
		block, err = aes.NewCipher(key)
		if err != nil {
			return true
		}
		// -- Make a GCM
		aesgcm, err = cipher.NewGCM(block)
		if err != nil {
			return true
		}
		// -- Open the seal
		decryptedText, err := aesgcm.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			return true
		}
		print.Infof("FOUND KEY!\n")
		print.Infof("potentialKey: %x\n", potentialKey)
		print.Infof("actualKey: \t%x\n", key)
		print.Infof("decryptedText: %s\n", string(decryptedText))
		return false
	}) {
		elapsed := time.Since(start)
		print.Infof("This took: %s\n", elapsed)
		return
	}
	print.Warnln("Failed to find a suitable key")
}
