package jwt_encode_verify

import (
	"log"
	"time"

	"github.com/Grs2080w/grp_server/core/crypto/sshutils"
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	Username string `json:"username"`
	Type_verification     string `json:"type_verification"`
}

// acess token exp 1 hour

func (t Token) CreateVerifyToken() string {
	signer, err := sshutils.LoadKeyPrivate()
	if err != nil {
		log.Fatalf("Error loading private key: %v", err)
	}

	// create the token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"username": t.Username,
		"type_verification": t.Type_verification,
		"exp":      time.Now().Add(time.Minute * 1).Unix(),
	})

	// Assign the token
	tokenString, err := token.SignedString(signer)
	if err != nil {
		log.Fatalf("Error signing token: %v", err)
	}

	return tokenString

}

