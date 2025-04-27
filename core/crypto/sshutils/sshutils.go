package sshutils

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// "C:/Users/Gabriel Santos/workspace/grp@server/ssh/"

func LoadKeyPrivate() (*rsa.PrivateKey, error) {
	privateBytes, err := os.ReadFile("C:/Users/Gabriel Santos/workspace/grp@server/ssh/.ssh")
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func LoadKeyPublic() (*rsa.PublicKey, error) {
	publicBytes, err := os.ReadFile("C:/Users/Gabriel Santos/workspace/grp@server/ssh/.ssh-pub.pem")
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}


// ssh-keygen -p -m PEM -f ~/.ssh/id_rsa
// C:\Users\Gabriel Santos\workspace\grp@server