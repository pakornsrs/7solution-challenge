package services

import (
	"context"
	mockrepositories "pakornssn/7solution-challenge/internal/mocks/mock-repository"
	"pakornssn/7solution-challenge/internal/models"
	authutil "pakornssn/7solution-challenge/pkg/auth-util"
	timeutil "pakornssn/7solution-challenge/pkg/time-util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRegister(t *testing.T) {
	mockPrimitiveId, _ := primitive.ObjectIDFromHex("68c0398a940bf6a1577e0716")
	newPrimitiveId = func() primitive.ObjectID {
		return mockPrimitiveId
	}

	mockTimeUtc := time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC)
	timeutil.TimeNowUte = func() time.Time {
		return mockTimeUtc
	}

	authutil.EncryptPassword = func(pwd string) (string, error) {
		return "mock-encrypt-password", nil
	}

	t.Run("success_register", func(t *testing.T) {
		// arrange
		mockUserRepository := &mockrepositories.MockUserRepository{}

		mockListUser := []models.UserDB{}
		mockUserRepository.On(
			"FindUser",
			mock.Anything,
			mock.Anything,
		).Return(mockListUser, 0, nil)

		mockUserRepository.On(
			"CreateUser",
			mock.Anything,
			mock.Anything,
		).Return(nil)

		ctx := context.Background()

		mockReq := models.CreateUserRequest{
			Name:     "user name",
			Email:    "username@gmail.com",
			Password: "123456",
		}

		service := userService{
			userRepository: mockUserRepository,
		}

		// act
		resp, err := service.Register(ctx, mockReq)

		// assert
		assert.NoError(t, err)

		expectResp := models.UserResponse{
			UserId:    "68c0398a940bf6a1577e0716",
			Name:      "user name",
			Email:     "username@gmail.com",
			CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
			UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
		}

		assert.Equal(t, expectResp, resp)

		mockUserRepository.AssertNumberOfCalls(t, "FindUser", 1)
		expectedFilter := models.GetUserFilterRequest{
			Email: "username@gmail.com",
		}
		mockUserRepository.AssertCalled(t, "FindUser", mock.Anything, expectedFilter)

		mockUserRepository.AssertNumberOfCalls(t, "CreateUser", 1)
		expectSaveUser := models.UserDB{
			Id:        mockPrimitiveId,
			Name:      "user name",
			Email:     "username@gmail.com",
			Password:  "mock-encrypt-password",
			CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
			UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
		}
		mockUserRepository.AssertCalled(t, "CreateUser", mock.Anything, expectSaveUser)
	})

	t.Run("email_existed_register_fail", func(t *testing.T) {
		// arrange
		mockUserRepository := &mockrepositories.MockUserRepository{}

		mockListUser := []models.UserDB{
			{
				Id:        mockPrimitiveId,
				Name:      "user name",
				Email:     "username@gmail.com",
				CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
				UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
			},
		}
		mockUserRepository.On(
			"FindUser",
			mock.Anything,
			mock.Anything,
		).Return(mockListUser, 0, nil)

		mockUserRepository.On(
			"CreateUser",
			mock.Anything,
			mock.Anything,
		).Return(nil)

		ctx := context.Background()

		mockReq := models.CreateUserRequest{
			Name:     "user name 2",
			Email:    "username@gmail.com",
			Password: "123456",
		}

		service := userService{
			userRepository: mockUserRepository,
		}

		// act
		resp, err := service.Register(ctx, mockReq)

		// assert
		assert.ErrorContains(t, err, "This e-main has been registed")

		expectResp := models.UserResponse{}
		assert.Equal(t, expectResp, resp)

		mockUserRepository.AssertNumberOfCalls(t, "FindUser", 1)
		expectedFilter := models.GetUserFilterRequest{
			Email: "username@gmail.com",
		}
		mockUserRepository.AssertCalled(t, "FindUser", mock.Anything, expectedFilter)

		mockUserRepository.AssertNumberOfCalls(t, "CreateUser", 0)
	})
}

func TestGetUserById(t *testing.T) {

	mockPrimitiveId, _ := primitive.ObjectIDFromHex("68c0398a940bf6a1577e0716")

	mockTimeUtc := time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC)
	timeutil.TimeNowUte = func() time.Time {
		return mockTimeUtc
	}

	t.Run("succes_get_user", func(t *testing.T) {
		// arrange
		mockUserRepository := &mockrepositories.MockUserRepository{}

		mockListUser := []models.UserDB{
			{
				Id:        mockPrimitiveId,
				Name:      "user name",
				Email:     "username@gmail.com",
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

		mockReq := "68c0398a940bf6a1577e0716"

		service := userService{
			userRepository: mockUserRepository,
		}

		// act
		resp, err := service.GetUserById(ctx, mockReq)

		// assert
		assert.NoError(t, err)

		expectResp := models.UserResponse{
			UserId:    "68c0398a940bf6a1577e0716",
			Name:      "user name",
			Email:     "username@gmail.com",
			CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
			UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
		}

		assert.Equal(t, expectResp, resp)

		mockUserRepository.AssertNumberOfCalls(t, "FindUser", 1)
		expectedFilter := models.GetUserFilterRequest{
			UserId: "68c0398a940bf6a1577e0716",
		}
		mockUserRepository.AssertCalled(t, "FindUser", mock.Anything, expectedFilter)
	})

	t.Run("not_found_user", func(t *testing.T) {
		// arrange
		mockUserRepository := &mockrepositories.MockUserRepository{}

		mockListUser := []models.UserDB{}
		mockUserRepository.On(
			"FindUser",
			mock.Anything,
			mock.Anything,
		).Return(mockListUser, 0, nil)

		ctx := context.Background()

		mockReq := "68c0398a940bf6a1577e0716"

		service := userService{
			userRepository: mockUserRepository,
		}

		// act
		resp, err := service.GetUserById(ctx, mockReq)

		// assert
		assert.ErrorContains(t, err, "Not found user")

		expectResp := models.UserResponse{}
		assert.Equal(t, expectResp, resp)

		mockUserRepository.AssertNumberOfCalls(t, "FindUser", 1)
		expectedFilter := models.GetUserFilterRequest{
			UserId: "68c0398a940bf6a1577e0716",
		}
		mockUserRepository.AssertCalled(t, "FindUser", mock.Anything, expectedFilter)
	})
}

func TestGetAllUser(t *testing.T) {
	mockPrimitiveId1, _ := primitive.ObjectIDFromHex("68c0398a940bf6a1577e0716")
	mockPrimitiveId2, _ := primitive.ObjectIDFromHex("68c03a6685b8677afa7bc52d")
	t.Run("succes_get_all_user_with_pagination", func(t *testing.T) {
		// arrange
		mockUserRepository := &mockrepositories.MockUserRepository{}

		mockListUser := []models.UserDB{
			{
				Id:        mockPrimitiveId1,
				Name:      "user name",
				Email:     "username@gmail.com",
				CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
				UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
			},
			{
				Id:        mockPrimitiveId2,
				Name:      "user name2",
				Email:     "username2@gmail.com",
				CreatedAt: time.Date(2030, 8, 22, 22, 58, 20, 0, time.UTC),
				UpdatedAt: time.Date(2030, 8, 31, 22, 58, 20, 0, time.UTC),
			},
		}
		mockPagination := models.Pagination{
			ItemPerPage: 10,
			CurrentPage: 1,
			TotalPage:   1,
			TotalItem:   2,
		}
		mockUserRepository.On(
			"FindUser",
			mock.Anything,
			mock.Anything,
		).Return(mockListUser, mockPagination, nil)

		ctx := context.Background()

		mockReq := &models.PaginationRequest{
			ItemPerPage: 10,
			CurrentPage: 1,
		}

		service := userService{
			userRepository: mockUserRepository,
		}

		// act
		resp, err := service.GetAllUser(ctx, mockReq)

		// assert
		assert.NoError(t, err)

		expectResp := models.GetAllUserResponse{
			List: []models.UserResponse{
				{
					UserId:    "68c0398a940bf6a1577e0716",
					Name:      "user name",
					Email:     "username@gmail.com",
					CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
					UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
				},
				{
					UserId:    "68c03a6685b8677afa7bc52d",
					Name:      "user name2",
					Email:     "username2@gmail.com",
					CreatedAt: time.Date(2030, 8, 22, 22, 58, 20, 0, time.UTC),
					UpdatedAt: time.Date(2030, 8, 31, 22, 58, 20, 0, time.UTC),
				},
			},
			Pagination: models.Pagination{
				ItemPerPage: 10,
				CurrentPage: 1,
				TotalPage:   1,
				TotalItem:   2,
			},
		}

		assert.Equal(t, expectResp, resp)

		mockUserRepository.AssertNumberOfCalls(t, "FindUser", 1)
		expectedFilter := models.GetUserFilterRequest{
			Pagination: &models.PaginationRequest{
				ItemPerPage: 10,
				CurrentPage: 1,
			},
		}
		mockUserRepository.AssertCalled(t, "FindUser", mock.Anything, expectedFilter)
	})
}

func TestUpdateUser(t *testing.T) {
	mockPrimitiveId, _ := primitive.ObjectIDFromHex("68c0398a940bf6a1577e0716")

	mockTimeUtc := time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC)
	timeutil.TimeNowUte = func() time.Time {
		return mockTimeUtc
	}

	t.Run("success_update_email_and_name", func(t *testing.T) {
		// arrange
		mockUserRepository := &mockrepositories.MockUserRepository{}
		mockListUser := []models.UserDB{}
		mockUserRepository.On(
			"FindUser",
			mock.Anything,
			mock.Anything,
		).Return(mockListUser, 0, nil).Once()

		mockListUser2 := []models.UserDB{
			{
				Id:        mockPrimitiveId,
				Name:      "user name",
				Email:     "username@gmail.com",
				CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
				UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
			},
		}
		mockUserRepository.On(
			"FindUser",
			mock.Anything,
			mock.Anything,
		).Return(mockListUser2, 0, nil).Once()

		mockUserRepository.On(
			"UpdateUser",
			mock.Anything,
			mock.Anything,
		).Return(nil)

		ctx := context.Background()

		mockReq := models.UpdateUserRequest{
			UserId:       "68c0398a940bf6a1577e0716",
			UpdatedName:  "user name",
			UpdatedEmail: "username@gmail.com",
		}

		service := userService{
			userRepository: mockUserRepository,
		}

		// act
		resp, err := service.UpdateUser(ctx, mockReq)

		// assert
		assert.NoError(t, err)

		expectResp := models.UpdateUserResponse{
			UpdatedUser: models.UserResponse{
				UserId:    "68c0398a940bf6a1577e0716",
				Name:      "user name",
				Email:     "username@gmail.com",
				CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
				UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
			},
		}

		assert.Equal(t, expectResp, resp)

		mockUserRepository.AssertNumberOfCalls(t, "FindUser", 2)
		mockUserRepository.AssertNumberOfCalls(t, "UpdateUser", 1)
	})

	t.Run("success_update_email_and_name_but_email_existed", func(t *testing.T) {
		// arrange
		mockUserRepository := &mockrepositories.MockUserRepository{}
		mockListUser := []models.UserDB{
			{
				Id:        mockPrimitiveId,
				Name:      "user name",
				Email:     "username@gmail.com",
				CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
				UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
			},
		}
		mockUserRepository.On(
			"FindUser",
			mock.Anything,
			mock.Anything,
		).Return(mockListUser, 0, nil).Once()

		mockUserRepository.On(
			"UpdateUser",
			mock.Anything,
			mock.Anything,
		).Return(nil)

		ctx := context.Background()

		mockReq := models.UpdateUserRequest{
			UserId:       "68c0398a940bf6a1577e0716",
			UpdatedName:  "user name",
			UpdatedEmail: "username@gmail.com",
		}

		service := userService{
			userRepository: mockUserRepository,
		}

		// act
		resp, err := service.UpdateUser(ctx, mockReq)

		// assert
		assert.ErrorContains(t, err, "This e-main has been registed")

		expectResp := models.UpdateUserResponse{}
		assert.Equal(t, expectResp, resp)

		mockUserRepository.AssertNumberOfCalls(t, "FindUser", 1)
		mockUserRepository.AssertNumberOfCalls(t, "UpdateUser", 0)
	})

	t.Run("success_update_name_only", func(t *testing.T) {
		// arrange
		mockUserRepository := &mockrepositories.MockUserRepository{}
		mockListUser := []models.UserDB{
			{
				Id:        mockPrimitiveId,
				Name:      "user name",
				Email:     "username@gmail.com",
				CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
				UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
			},
		}
		mockUserRepository.On(
			"FindUser",
			mock.Anything,
			mock.Anything,
		).Return(mockListUser, 0, nil).Once()

		mockUserRepository.On(
			"UpdateUser",
			mock.Anything,
			mock.Anything,
		).Return(nil)

		ctx := context.Background()

		mockReq := models.UpdateUserRequest{
			UserId:      "68c0398a940bf6a1577e0716",
			UpdatedName: "user name",
		}

		service := userService{
			userRepository: mockUserRepository,
		}

		// act
		resp, err := service.UpdateUser(ctx, mockReq)

		// assert
		assert.NoError(t, err)

		expectResp := models.UpdateUserResponse{
			UpdatedUser: models.UserResponse{
				UserId:    "68c0398a940bf6a1577e0716",
				Name:      "user name",
				Email:     "username@gmail.com",
				CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
				UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
			},
		}

		assert.Equal(t, expectResp, resp)

		mockUserRepository.AssertNumberOfCalls(t, "FindUser", 1)
		mockUserRepository.AssertNumberOfCalls(t, "UpdateUser", 1)
	})
}

func TestDeleteUser(t *testing.T) {
	mockPrimitiveId, _ := primitive.ObjectIDFromHex("68c0398a940bf6a1577e0716")
	t.Run("delete_success", func(t *testing.T) {
		// arrange
		mockUserRepository := &mockrepositories.MockUserRepository{}

		mockListUser := []models.UserDB{
			{
				Id:        mockPrimitiveId,
				Name:      "user name",
				Email:     "username@gmail.com",
				CreatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
				UpdatedAt: time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC),
			},
		}
		mockUserRepository.On(
			"FindUser",
			mock.Anything,
			mock.Anything,
		).Return(mockListUser, 0, nil)

		mockUserRepository.On(
			"DeleteUser",
			mock.Anything,
			mock.Anything,
		).Return(nil)

		ctx := context.Background()

		mockReq := "68c0398a940bf6a1577e0716"

		service := userService{
			userRepository: mockUserRepository,
		}

		// act
		err := service.DeleteUser(ctx, mockReq)

		// assert
		assert.NoError(t, err)

		mockUserRepository.AssertNumberOfCalls(t, "FindUser", 1)
		expectedFilter := models.GetUserFilterRequest{
			UserId: "68c0398a940bf6a1577e0716",
		}
		mockUserRepository.AssertCalled(t, "FindUser", mock.Anything, expectedFilter)

		mockUserRepository.AssertNumberOfCalls(t, "DeleteUser", 1)
		expectPrimitiveId, _ := primitive.ObjectIDFromHex("68c0398a940bf6a1577e0716")
		mockUserRepository.AssertCalled(t, "DeleteUser", mock.Anything, expectPrimitiveId)
	})

	t.Run("no_user_to_delete", func(t *testing.T) {
		// arrange
		mockUserRepository := &mockrepositories.MockUserRepository{}

		mockListUser := []models.UserDB{}
		mockUserRepository.On(
			"FindUser",
			mock.Anything,
			mock.Anything,
		).Return(mockListUser, 0, nil)

		mockUserRepository.On(
			"DeleteUser",
			mock.Anything,
			mock.Anything,
		).Return(nil)

		ctx := context.Background()

		mockReq := "68c0398a940bf6a1577e0716"

		service := userService{
			userRepository: mockUserRepository,
		}

		// act
		err := service.DeleteUser(ctx, mockReq)

		// assert
		assert.ErrorContains(t, err, "Not found user")

		mockUserRepository.AssertNumberOfCalls(t, "FindUser", 1)
		expectedFilter := models.GetUserFilterRequest{
			UserId: "68c0398a940bf6a1577e0716",
		}
		mockUserRepository.AssertCalled(t, "FindUser", mock.Anything, expectedFilter)

		mockUserRepository.AssertNumberOfCalls(t, "DeleteUser", 0)
	})
}
