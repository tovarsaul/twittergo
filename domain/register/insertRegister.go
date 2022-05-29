package register

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittergo/domain/config"
	"twittergo/domain/utils"
	"twittergo/models"
)

func InsertRegister(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := config.MongoDBConnection.Database("twitter")
	col := db.Collection("user")
	user.Password, _ = utils.Encrypt(user.Password)
	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, err
}
