package models

import (
	"pakornssn/7solution-challenge/internal/constants"
	errorutil "pakornssn/7solution-challenge/pkg/error-util"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDB struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type GetUserFilterRequest struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`

	Pagination *PaginationRequest `json:"pagination"`
}

type UpdateUserDb struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type UpdateUserRequest struct {
	UserId       string `json:"userId"`
	UpdatedName  string `json:"updatedName"`
	UpdatedEmail string `json:"updatedEmail"`
}

func (source *UpdateUserRequest) ToUpdateUserDb() (UpdateUserDb, error) {
	id, err := primitive.ObjectIDFromHex(source.UserId)
	if err != nil {
		errorResp := errorutil.GetInternalServerError(err, &constants.UserIdFormatIncorrectError)
		return UpdateUserDb{}, &errorResp
	}

	return UpdateUserDb{
		Id:        id,
		Name:      source.UpdatedName,
		Email:     source.UpdatedEmail,
		UpdatedAt: time.Time{},
	}, nil
}

type AuthUserRequest struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
