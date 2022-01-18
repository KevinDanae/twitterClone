package jwt

import (
	"time"

	"github.com/KevinDanae/twitterClone/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT is a function that generates a JWT token
func GenerateJWT(t models.User) (string, error) {
	key := []byte("my_secret_key")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastName":  t.LastName,
		"birthday":  t.Birthday,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.Website,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
