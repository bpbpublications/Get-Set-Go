package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"
)

func encrypt(data []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    ciphertext := make([]byte, aes.BlockSize+len(data))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

    return ciphertext, nil
}

func main() {
    key := []byte("a very very very very secret key") // Must be 32 bytes
    data := []byte("Hello, Go!")

    encryptedData, err := encrypt(data, key)
    if err != nil {
        fmt.Println("Error encrypting data:", err)
        return
    }

    fmt.Printf("Encrypted Data: %x\n", encryptedData)
}
