package repositories

import (
	"context"
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
	FindUser(ctx context.Context, req models.GetUserFilterRequest) ([]models.UserDB, int, error)
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
		respError := errorutil.GetInternalServerError(err, &constants.MongoDbError)
		return &respError
	}

	return nil
}

func (repo *userRepository) FindUser(ctx context.Context, req models.GetUserFilterRequest) ([]models.UserDB, int, error) {

	filter, err := buildFindUserFilter(req)
	if err != nil {
		return nil, 0, err
	}

	findOpts := &options.FindOptions{}
	if req.Pagination != nil {
		skip := (req.Pagination.CurrentPage - 1) * req.Pagination.ItemPerPage

		findOpts = options.Find().
			SetSkip(skip).
			SetLimit(req.Pagination.ItemPerPage).
			SetSort(bson.D{
				{
					Key:   "createdAt",
					Value: -1,
				},
			})
	}

	cur, err := repo.mongoCollection.Find(ctx, filter, findOpts)
	if err != nil {
		respError := errorutil.GetInternalServerError(err, &constants.MongoDbError)
		return nil, 0, &respError
	}

	result := []models.UserDB{}
	if err := cur.All(ctx, &result); err != nil {
		return nil, 0, err
	}

	totalItems := len(result)
	if req.Pagination != nil && len(result) > 0 {
		total, err := repo.mongoCollection.CountDocuments(ctx, filter)
		if err != nil {
			respError := errorutil.GetInternalServerError(err, &constants.MongoDbError)
			return nil, 0, &respError
		}
		totalItems = int(total)
	}

	return result, totalItems, nil
}

func buildFindUserFilter(req models.GetUserFilterRequest) (bson.M, error) {
	filters := []bson.M{}

	if len(req.UserId) > 0 {
		id, err := primitive.ObjectIDFromHex(req.UserId)
		if err != nil {
			respError := errorutil.GetInternalServerError(err, &constants.UserIdFormatIncorrectError)
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

	return bson.M{"$and": filters}, nil
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
		respError := errorutil.GetInternalServerError(err, &constants.MongoDbError)
		return &respError
	}

	return nil
}

func (repo *userRepository) DeleteUser(ctx context.Context, userId primitive.ObjectID) error {

	filter := bson.M{"_id": userId}

	_, err := repo.mongoCollection.DeleteOne(ctx, filter)
	if err != nil {
		respError := errorutil.GetInternalServerError(err, &constants.MongoDbError)
		return &respError
	}

	return nil
}
