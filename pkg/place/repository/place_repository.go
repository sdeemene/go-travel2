package repository

import (
	"context"
	dbs "github.com/stdeemene/go-travel2/config/db"
	"github.com/stdeemene/go-travel2/pkg/place"
	"github.com/stdeemene/go-travel2/pkg/place/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlaceRepo struct {
	conn *dbs.MongoDB
}

func NewPlaceRepo(conn *dbs.MongoDB) place.PlaceRepositoryImp {
	return PlaceRepo{conn: conn}
}

func (repo PlaceRepo) CreatePlace(place *model.Place, ctx context.Context) (interface{}, error) {
	placeColl := repo.conn.PlaceCollection()
	result, err := placeColl.InsertOne(ctx, place)
	return result.InsertedID, err
}
func (repo PlaceRepo) GetPlace(filter primitive.M, ctx context.Context) (model.Place, error) {
	var p model.Place
	placeColl := repo.conn.PlaceCollection()
	err := placeColl.FindOne(ctx, filter).Decode(&p)
	return p, err
}
func (repo PlaceRepo) DeletePlace(filter primitive.M, ctx context.Context) (int64, error) {
	placeColl := repo.conn.PlaceCollection()
	result, err := placeColl.DeleteOne(ctx, filter)
	return result.DeletedCount, err
}
func (repo PlaceRepo) UpdatePlace(filter primitive.M, place model.Place, ctx context.Context) (model.Place, error) {
	placeColl := repo.conn.PlaceCollection()
	err := placeColl.FindOneAndUpdate(ctx, filter, bson.M{"$set": place}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&place)
	return place, err
}
func (repo PlaceRepo) GetPlaces(ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	filter := bson.D{{}}
	placeColl := repo.conn.PlaceCollection()
	cursor, err := placeColl.Find(ctx, filter, findOptions)
	return cursor, err
}
func (repo PlaceRepo) SearchPlace(filter interface{}, ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	placeColl := repo.conn.PlaceCollection()
	cursor, err := placeColl.Find(ctx, filter, findOptions)
	return cursor, err
}
