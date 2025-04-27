package randomPassword

import (
	"crypto/rand"
	"math/big"
)

// "github.com/Grs2080w/grp_server/core/utils/ramdomPaswword"


func RandomPassword(length int, upper bool, lower bool, numbers bool, special bool) (string, error) {
    const (
        uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZABCD"
        lowercaseLetters = "abcdefghijklmnopqrstuvwxyzabcd"
        digits           = "012345678901234567890123456789"
        specialChars     = "!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
    )

	var allChars string

	if upper {
		allChars += uppercaseLetters
	}

	if lower {
		allChars += lowercaseLetters
	}

	if numbers {
		allChars += digits
	}

	if special {
		allChars += specialChars
	}

	var r int = 0

    var password string
    for i := 0; i < length; i++ {
        randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
        if err != nil {
            return "", err
        }
        password += string(allChars[randomIndex.Int64()])

		r += 1
    }

    return password, nil
}
