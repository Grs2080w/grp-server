package jwt_encode_refresh

import (
	"log"
	"time"

	"github.com/Grs2080w/grp_server/core/crypto/sshutils"
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	Username string `json:"username"`
}

// refresh token exp 720 hours or 30 days

func (t Token) CreateRefreshToken() string {
	signer, err := sshutils.LoadKeyPrivate()
	if err != nil {
		log.Fatalf("Error loading private key: %v", err)
	}

	// create the token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"username": t.Username,
		"exp":      time.Now().Add(time.Hour * 720).Unix(),
	})

	// Assign the token
	tokenString, err := token.SignedString(signer)
	if err != nil {
		log.Fatalf("Error signing token: %v", err)
	}

	return tokenString

}