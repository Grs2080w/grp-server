package jwt_decode

import (
	"errors"
	"fmt"
	"log"

	"github.com/Grs2080w/grp_server/core/crypto/sshutils"
	"github.com/golang-jwt/jwt/v5"
)

func DecodeToken(tokenString string) (string, error) {

	publicKey, err := sshutils.LoadKeyPublic()
	if err != nil {
		log.Fatalf("Error loading public key: %v", err)
	}
    
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return publicKey, nil
    })

    if err != nil {
		return "", errors.New("invalid token")
    }

    if !token.Valid {
        return "", errors.New("invalid token")
    }

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error parsing claims")
		return "", errors.New("invalid token")
	}

	username, ok := claims["username"].(string)
	if !ok {
		fmt.Println("Error parsing username")
		return "", errors.New("invalid token")
	}

    return username, nil
}



