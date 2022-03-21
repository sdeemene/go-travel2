package address

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/stdeemene/go-travel2/pkg/address/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddressRepositoryImp interface {
	CreateAddress(address model.Address, ctx context.Context) (interface{}, error)
	GetAddress(id primitive.M, ctx context.Context) (model.Address, error)
	DeleteAddress(filter primitive.M, ctx context.Context) (int64, error)
	UpdateAddress(filter primitive.M, address model.Address, ctx context.Context) (model.Address, error)
	GetAddresses(ctx context.Context) (*mongo.Cursor, error)
	SearchAddress(filter interface{}, ctx context.Context) (*mongo.Cursor, error)
}

type AddressServiceImp interface {
	CreateAddress(address model.Address, ctx context.Context) (model.Address, error)
	GetAddress(id string, ctx context.Context) (model.Address, error)
	DeleteAddress(id string, ctx context.Context) (model.AddressDeleted, error)
	UpdateAddress(id string, address model.Address, ctx context.Context) (model.AddressUpdated, error)
	GetAddresses(ctx context.Context) ([]model.Address, error)
	SearchAddress(filter interface{}, ctx context.Context) ([]model.Address, error)
}
