package bd

import (
	"context"
	"time"

	"github.com/KevinDanae/twitterClone/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RegisterUser is a function that registers a new user in to the database and returns the status of the operation
func RegisterUser(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitterClone")
	col := db.Collection("users")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
