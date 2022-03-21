package repository

import (
	"context"
	dbs "github.com/stdeemene/go-travel2/config/db"
	"github.com/stdeemene/go-travel2/pkg/travel"
	"github.com/stdeemene/go-travel2/pkg/travel/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TravelRepo struct {
	conn *dbs.MongoDB
}

func NewTravelRepo(conn *dbs.MongoDB) travel.TravelRepositoryImp {
	return TravelRepo{conn: conn}
}

func (repo TravelRepo) CreateTravel(travel *model.Travel, ctx context.Context) (interface{}, error) {
	travelColl := repo.conn.TravelCollection()
	result, err := travelColl.InsertOne(ctx, travel)
	return result.InsertedID, err
}
func (repo TravelRepo) GetTravel(filter primitive.M, ctx context.Context) (model.Travel, error) {
	var p model.Travel
	travelColl := repo.conn.TravelCollection()
	err := travelColl.FindOne(ctx, filter).Decode(&p)
	return p, err
}
func (repo TravelRepo) DeleteTravel(filter primitive.M, ctx context.Context) (int64, error) {
	travelColl := repo.conn.TravelCollection()
	result, err := travelColl.DeleteOne(ctx, filter)
	return result.DeletedCount, err
}
func (repo TravelRepo) UpdateTravel(filter primitive.M, travel model.Travel, ctx context.Context) (model.Travel, error) {
	travelColl := repo.conn.TravelCollection()
	err := travelColl.FindOneAndUpdate(ctx, filter, bson.M{"$set": travel}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&travel)
	return travel, err
}
func (repo TravelRepo) GetTravels(ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	filter := bson.D{{}}
	travelColl := repo.conn.TravelCollection()
	cursor, err := travelColl.Find(ctx, filter, findOptions)
	return cursor, err
}
func (repo TravelRepo) SearchTravel(filter interface{}, ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	travelColl := repo.conn.TravelCollection()
	cursor, err := travelColl.Find(ctx, filter, findOptions)
	return cursor, err
}
func (repo TravelRepo) GetUserTravels(filter interface{}, ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	travelColl := repo.conn.TravelCollection()
	cursor, err := travelColl.Find(ctx, filter, findOptions)
	return cursor, err
}
func (repo TravelRepo) RateTravel(filter primitive.M, update primitive.M, ctx context.Context) (int64, error) {
	travelColl := repo.conn.TravelCollection()
	result, err := travelColl.UpdateOne(ctx, filter, update)
	return result.ModifiedCount, err
}
