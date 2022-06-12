package profile

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittergo/domain/config"
	"twittergo/models"
)

func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := config.MongoDBConnection.Database("twitter")
	col := db.Collection("user")
	var profile models.User
	objId, _ := primitive.ObjectIDFromHex(ID)

	query := bson.M{
		"_id": objId,
	}
	err := col.FindOne(ctx, query).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Register not find " + err.Error())
		return profile, err
	}
	return profile, err
}
