package services

import (
	"context"
	"os"
	mockrepositories "pakornssn/7solution-challenge/internal/mocks/mock-repository"
	"pakornssn/7solution-challenge/internal/models"
	timeutil "pakornssn/7solution-challenge/pkg/time-util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAuthenticationUser(t *testing.T) {
	mockPrimitiveId, _ := primitive.ObjectIDFromHex("68c0398a940bf6a1577e0716")
	mockTimeUtc := time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC)
	timeutil.TimeNowUte = func() time.Time {
		return mockTimeUtc
	}

	os.Setenv("JWT_SECRET_KEY", "aJm5+7P3tFbDqZHTq8gW4nqoxzWJ3i28bM6oW8Zz9Xg=")
	os.Setenv("ISSUER", "user-service")
	os.Setenv("AUDIENCE", "user-service-client")

	t.Run("login_success", func(t *testing.T) {
		// arragen
		mockUserRepository := &mockrepositories.MockUserRepository{}
		mockListUser := []models.UserDB{
			{
				Id:        mockPrimitiveId,
				Name:      "user name",
				Email:     "username@gmail.com",
				Password:  "$2a$10$FnuEgIkTYXoYPYKChPp8aOnd4/K625JsM9jsyvexylelZQrbTmGy.",
				CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
				UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
			},
		}
		mockUserRepository.On(
			"FindUser",
			mock.Anything,
			mock.Anything,
		).Return(mockListUser, 0, nil)

		ctx := context.Background()

		mockReq := models.AuthUserRequest{
			Email:    "username@gmail.com",
			Password: "12345678",
		}

		service := authenticationService{
			userRepository: mockUserRepository,
		}

		// act
		resp, err := service.AuthenticationUser(ctx, mockReq)

		// assert
		assert.NoError(t, err)

		expectResp := models.AuthUserResponse{
			User: models.UserResponse{
				UserId:    "68c0398a940bf6a1577e0716",
				Name:      "user name",
				Email:     "username@gmail.com",
				CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
				UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
			},
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyLXNlcnZpY2UiLCJzdWIiOiI2OGMwMzk4YTk0MGJmNmExNTc3ZTA3MTYiLCJhdWQiOlsidXNlci1zZXJ2aWNlLWNsaWVudCJdLCJleHAiOjE5MTEyNTE2MDAsIm5iZiI6MTkxMTI1MDcwMCwiaWF0IjoxOTExMjUwNzAwfQ.UNrvz4nAME_fI0lAPRLch13REdnhZ4EEf0Gu7PsZwRQ",
		}
		assert.Equal(t, expectResp, resp)
	})

	t.Run("login_fail_wrong_password", func(t *testing.T) {
		// arragen
		mockUserRepository := &mockrepositories.MockUserRepository{}
		mockListUser := []models.UserDB{
			{
				Id:        mockPrimitiveId,
				Name:      "user name",
				Email:     "username@gmail.com",
				Password:  "$2a$10$FnuEgIkTYXoYPYKChPp8aOnd4/K625JsM9jsyvexylelZQrbTmGy.",
				CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
				UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
			},
		}
		mockUserRepository.On(
			"FindUser",
			mock.Anything,
			mock.Anything,
		).Return(mockListUser, 0, nil)

		ctx := context.Background()

		mockReq := models.AuthUserRequest{
			Email:    "username@gmail.com",
			Password: "12345678910",
		}

		service := authenticationService{
			userRepository: mockUserRepository,
		}

		// act
		resp, err := service.AuthenticationUser(ctx, mockReq)

		// assert
		assert.ErrorContains(t, err, "Not found user or password is incorrect")

		expectResp := models.AuthUserResponse{}
		assert.Equal(t, expectResp, resp)
	})
}
