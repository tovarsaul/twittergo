package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct {
	Mail string             `json:"mail"`
	ID   primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}
