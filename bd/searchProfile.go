package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/KevinDanae/twitterClone/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SearchProfile searches for a user in the database
func SearchProfile(ID string) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitterClone")
	col := db.Collection("users")

	var profile models.User
	objId, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objId,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())
		return profile, err
	}

	return profile, nil

}
