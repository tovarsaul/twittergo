package register

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"twittergo/domain/config"
	"twittergo/models"
)

func CheckIfUserExists(mail string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := config.MongoDBConnection.Database("twitter")
	col := db.Collection("user")
	condition := bson.M{"mail": mail}
	var result models.User
	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
