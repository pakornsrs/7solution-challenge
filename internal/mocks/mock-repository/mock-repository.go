package mockrepositories

import (
	"context"
	"pakornssn/7solution-challenge/internal/models"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserRepository struct {
	mock.Mock
}

func (mock *MockUserRepository) CreateUser(ctx context.Context, request models.UserDB) error {
	result := mock.Called(ctx, request)

	res0 := result.Error(0)
	return res0
}

func (mock *MockUserRepository) FindUser(ctx context.Context, req models.GetUserFilterRequest) ([]models.UserDB, models.Pagination, error) {
	result := mock.Called(ctx, req)

	var res0 []models.UserDB
	if val0, pass := result.Get(0).([]models.UserDB); pass {
		res0 = val0
	}

	var res1 models.Pagination
	if val1, pass := result.Get(1).(models.Pagination); pass {
		res1 = val1
	}

	res2 := result.Error(2)
	return res0, res1, res2
}

func (mock *MockUserRepository) UpdateUser(ctx context.Context, request models.UpdateUserDb) error {
	result := mock.Called(ctx, request)

	res0 := result.Error(0)
	return res0
}

func (mock *MockUserRepository) DeleteUser(ctx context.Context, userId primitive.ObjectID) error {
	result := mock.Called(ctx, userId)

	res0 := result.Error(0)
	return res0
}
