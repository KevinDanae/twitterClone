package bd

import (
	"context"
	"time"

	"github.com/KevinDanae/twitterClone/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitterClone")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var resultado models.User
	err := col.FindOne(ctx, condition).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
