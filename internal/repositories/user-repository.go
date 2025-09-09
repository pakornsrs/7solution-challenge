package repositories

import (
	"context"
	"errors"
	"math"
	"pakornssn/7solution-challenge/internal/constants"
	"pakornssn/7solution-challenge/internal/models"
	errorutil "pakornssn/7solution-challenge/pkg/error-util"
	"regexp"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user models.UserDB) error
	FindUser(ctx context.Context, req models.GetUserFilterRequest) ([]models.UserDB, models.Pagination, error)
	UpdateUser(ctx context.Context, req models.UpdateUserDb) error
	DeleteUser(ctx context.Context, userId primitive.ObjectID) error
}

type userRepository struct {
	mongoCollection *mongo.Collection
}

func NewUserRepository(dbClient *mongo.Client) IUserRepository {
	return &userRepository{
		mongoCollection: getCollectionUser(dbClient),
	}
}

func getCollectionUser(client *mongo.Client) *mongo.Collection {
	return client.Database(constants.DatabaseName).Collection(constants.Collection_User)
}

func (repo *userRepository) CreateUser(ctx context.Context, user models.UserDB) error {

	_, err := repo.mongoCollection.InsertOne(ctx, user)
	if err != nil {
		respError := errorutil.GetServerErrorResponse(500, err, &constants.MongoDbError)
		return &respError
	}

	return nil
}

func (repo *userRepository) FindUser(ctx context.Context, req models.GetUserFilterRequest) ([]models.UserDB, models.Pagination, error) {

	result := []models.UserDB{}
	pagination := models.Pagination{}

	filter, err := repo.buildFindUserFilter(req)
	if err != nil {
		return nil, pagination, err
	}

	findOpts := &options.FindOptions{}
	if req.Pagination != nil {
		findOpts, pagination, err = repo.buildPaginationOption(ctx, filter, *req.Pagination)
		if err != nil {
			return nil, pagination, err
		}
	}

	cur, err := repo.mongoCollection.Find(ctx, filter, findOpts)
	if err != nil {
		respError := errorutil.GetServerErrorResponse(500, err, &constants.MongoDbError)
		return nil, pagination, &respError
	}

	if err := cur.All(ctx, &result); err != nil {
		return nil, pagination, err
	}

	// default pagination
	if req.Pagination == nil {
		itemsCount := len(result)
		pagination = models.Pagination{
			ItemPerPage: int64(itemsCount),
			CurrentPage: 1,
			TotalPage:   1,
			TotalItem:   int64(itemsCount),
		}
	}

	return result, pagination, nil
}

func (repo *userRepository) buildFindUserFilter(req models.GetUserFilterRequest) (bson.M, error) {
	filters := []bson.M{}

	if len(req.UserId) > 0 {
		id, err := primitive.ObjectIDFromHex(req.UserId)
		if err != nil {
			respError := errorutil.GetServerErrorResponse(500, err, &constants.UserIdFormatIncorrectError)
			return nil, &respError
		}
		filter := bson.M{
			"_id": id,
		}

		filters = append(filters, filter)
	}

	if len(req.Email) > 0 {
		pat := regexp.QuoteMeta(strings.TrimSpace(req.Email))

		filter := bson.M{
			"email": bson.M{
				"$regex":   pat,
				"$options": "i",
			},
		}
		filters = append(filters, filter)
	}

	if len(filters) == 0 {
		return bson.M{}, nil
	}

	return bson.M{"$and": filters}, nil
}

func (repo *userRepository) buildPaginationOption(ctx context.Context, filter bson.M, pagination models.PaginationRequest) (*options.FindOptions, models.Pagination, error) {

	total, err := repo.mongoCollection.CountDocuments(ctx, filter)
	if err != nil {
		respError := errorutil.GetServerErrorResponse(500, err, &constants.MongoDbError)
		return nil, models.Pagination{}, &respError
	}

	totalItems := int(total)
	totalPage := math.Ceil(float64(totalItems) / float64(pagination.ItemPerPage))

	if totalPage < float64(pagination.CurrentPage) {
		respError := errorutil.GetServerErrorResponse(400, errors.New("current page not over total page"), &constants.BadRequestError)
		return nil, models.Pagination{}, &respError
	}

	skip := (pagination.CurrentPage - 1) * pagination.ItemPerPage

	findOpts := options.Find().
		SetSkip(skip).
		SetLimit(pagination.ItemPerPage).
		SetSort(bson.D{
			{
				Key:   "createdAt",
				Value: -1,
			},
		})

	paginationResp := models.Pagination{
		ItemPerPage: pagination.ItemPerPage,
		CurrentPage: pagination.CurrentPage,
		TotalPage:   int64(totalPage),
		TotalItem:   int64(totalItems),
	}

	return findOpts, paginationResp, nil
}

func (repo *userRepository) UpdateUser(ctx context.Context, req models.UpdateUserDb) error {

	filter := bson.M{"_id": req.Id}

	updated := bson.M{}

	if len(req.Email) > 0 {
		updated["email"] = req.Email
	}

	if len(req.Name) > 0 {
		updated["name"] = req.Name
	}

	updated["updatedAt"] = time.Now()

	_, err := repo.mongoCollection.UpdateOne(ctx, filter, updated)
	if err != nil {
		respError := errorutil.GetServerErrorResponse(500, err, &constants.MongoDbError)
		return &respError
	}

	return nil
}

func (repo *userRepository) DeleteUser(ctx context.Context, userId primitive.ObjectID) error {

	filter := bson.M{"_id": userId}

	_, err := repo.mongoCollection.DeleteOne(ctx, filter)
	if err != nil {
		respError := errorutil.GetServerErrorResponse(500, err, &constants.MongoDbError)
		return &respError
	}

	return nil
}
