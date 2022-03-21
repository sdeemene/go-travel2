package place

import (
	"context"
	"github.com/stdeemene/go-travel2/pkg/place/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlaceRepositoryImp interface {
	CreatePlace(place *model.Place, ctx context.Context) (interface{}, error)
	GetPlace(filter primitive.M, ctx context.Context) (model.Place, error)
	DeletePlace(filter primitive.M, ctx context.Context) (int64, error)
	UpdatePlace(filter primitive.M, place model.Place, ctx context.Context) (model.Place, error)
	GetPlaces(ctx context.Context) (*mongo.Cursor, error)
	SearchPlace(filter interface{}, ctx context.Context) (*mongo.Cursor, error)
}

type PlaceServiceImp interface {
	CreatePlace(payload *model.PlaceReq, ctx context.Context) (*model.Place, error)
	GetPlace(id string, ctx context.Context) (*model.Place, error)
	DeletePlace(id string, ctx context.Context) (model.PlaceDeleted, error)
	UpdatePlace(id string, place model.Place, ctx context.Context) (model.PlaceUpdated, error)
	GetPlaces(ctx context.Context) ([]model.Place, error)
	SearchPlace(filter interface{}, ctx context.Context) ([]model.Place, error)
}
