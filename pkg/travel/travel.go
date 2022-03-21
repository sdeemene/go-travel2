package travel

import (
	"context"
	"github.com/stdeemene/go-travel2/pkg/travel/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TravelRepositoryImp interface {
	CreateTravel(travel *model.Travel, ctx context.Context) (interface{}, error)
	GetTravel(filter primitive.M, ctx context.Context) (model.Travel, error)
	DeleteTravel(filter primitive.M, ctx context.Context) (int64, error)
	UpdateTravel(filter primitive.M, travel model.Travel, ctx context.Context) (model.Travel, error)
	GetTravels(ctx context.Context) (*mongo.Cursor, error)
	SearchTravel(filter interface{}, ctx context.Context) (*mongo.Cursor, error)
	GetUserTravels(filter interface{}, ctx context.Context) (*mongo.Cursor, error)
	RateTravel(filter primitive.M, update primitive.M, ctx context.Context) (int64, error)
}

type TravelServiceImp interface {
	CreateTravel(payload *model.TravelReq, ctx context.Context) (*model.Travel, error)
	GetTravel(id string, ctx context.Context) (*model.Travel, error)
	DeleteTravel(id string, ctx context.Context) (model.TravelDeleted, error)
	UpdateTravel(id string, place model.Travel, ctx context.Context) (model.TravelUpdated, error)
	GetTravels(ctx context.Context) ([]model.Travel, error)
	SearchTravel(filter interface{}, ctx context.Context) ([]model.Travel, error)
	GetUserTravels(id string, ctx context.Context) ([]model.Travel, error)
	RateTravel(id string, rateValue *model.TravelRateReq, ctx context.Context) (model.TravelUpdated, error)
}
