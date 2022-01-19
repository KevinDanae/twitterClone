package routes

import (
	"errors"
	"strings"

	"github.com/KevinDanae/twitterClone/bd"
	"github.com/KevinDanae/twitterClone/models"
	"github.com/dgrijalva/jwt-go"
)

// Email is a var that stores the email of the user
var Email string

// IdUser is a var that stores the id of the user
var IdUser string

// ProcessToken is a function that processes the token
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	key := []byte("my_secret_key")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("invalid token")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err == nil {
		_, exist, _ := bd.CheckUserExist(claims.Email)
		if exist {
			Email = claims.Email
			IdUser = claims.ID.Hex()
		}

		return claims, exist, IdUser, nil
	}

	if !tkn.Valid {
		return claims, false, "", errors.New("invalid token")
	}

	return claims, false, "", err
}
