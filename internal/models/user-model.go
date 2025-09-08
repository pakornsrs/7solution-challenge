package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDB struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	IsActive  bool               `bson:"isActive"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
