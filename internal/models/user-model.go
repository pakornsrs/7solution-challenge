package models

import (
	"pakornssn/7solution-challenge/internal/constants"
	errorutil "pakornssn/7solution-challenge/pkg/error-util"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserDB struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func (source *UserDB) ToUserResponse() UserResponse {
	return UserResponse{
		UserId:    source.Id.Hex(),
		Name:      source.Name,
		Email:     source.Email,
		CreatedAt: source.CreatedAt,
		UpdatedAt: source.UpdatedAt,
	}
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
		errorResp := errorutil.GetServerErrorResponse(500, err, &constants.UserIdFormatIncorrectError)
		return UpdateUserDb{}, &errorResp
	}

	return UpdateUserDb{
		Id:        id,
		Name:      source.UpdatedName,
		Email:     source.UpdatedEmail,
		UpdatedAt: time.Time{},
	}, nil
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (source *CreateUserRequest) ToUserDb() (UserDB, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(source.Password), bcrypt.DefaultCost)
	if err != nil {
		respError := errorutil.GetServerErrorResponse(500, err, &constants.UnknownError)
		return UserDB{}, &respError
	}

	currentTime := time.Now()

	return UserDB{
		Name:      source.Name,
		Email:     source.Email,
		Password:  string(password),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}, nil
}

type UserResponse struct {
	UserId    string    `json:"userId"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type GetAllUserResponse struct {
	List       []UserResponse `json:"list"`
	Pagination Pagination     `json:"pagination"`
}

type UpdateUserResponse struct {
	UpdatedUser UserResponse `json:"updatedUser"`
}
