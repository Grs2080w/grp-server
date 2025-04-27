package pwd_manager

// "github.com/Grs2080w/grp_server/core/crypto/pwd_manager"

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

func EncryptAES_GCM(pwd []byte, key []byte) (string, error) {

	nonce := make([]byte, 12)
	_, err := rand.Read(nonce)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, pwd, nil)

	chipherText := base64.StdEncoding.EncodeToString(cipherText)

	return chipherText, nil
}

func DecryptAES_GCM(cipherStr string, key []byte) ([]byte, error) {

	cipherText, err := base64.StdEncoding.DecodeString(cipherStr)

	if err != nil {
		return nil, err
	}
	
	nonce := cipherText[:12]
	cipherText = cipherText[12:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}


/*
func main() {
	key := []byte("gabrieledemisneh")
	plainText := []byte("Texto muito secreto!")

	cipherText, err := EncryptAES_GCM(plainText, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Texto cifrado: %x\n", cipherText)

	cipher := "apwZXUmfVkZrFt4mrVXFLspPaj94ZSKo2HLil/5Qzm/GIrzHhko/LrmBHok="

	key2 := []byte("1234567887654321")

	decryptedText, err := DecryptAES_GCM(cipher, key2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Texto original: %s\n", decryptedText)
}

*/