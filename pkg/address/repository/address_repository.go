package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"

	dbs "github.com/stdeemene/go-travel2/config/db"
	"github.com/stdeemene/go-travel2/pkg/address"
	"github.com/stdeemene/go-travel2/pkg/address/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AddressRepo struct {
	conn *dbs.MongoDB
}

func NewAddressRepo(conn *dbs.MongoDB) address.AddressRepositoryImp {
	return AddressRepo{conn: conn}
}

func (repo AddressRepo) CreateAddress(address model.Address, ctx context.Context) (interface{}, error) {
	addressColl := repo.conn.AddressCollection()
	result, err := addressColl.InsertOne(ctx, address)
	return result.InsertedID, err
}
func (repo AddressRepo) GetAddress(filter primitive.M, ctx context.Context) (model.Address, error) {
	var addr model.Address
	addressColl := repo.conn.AddressCollection()
	err := addressColl.FindOne(ctx, filter).Decode(&addr)
	return addr, err

}
func (repo AddressRepo) DeleteAddress(filter primitive.M, ctx context.Context) (int64, error) {
	addressColl := repo.conn.AddressCollection()
	result, err := addressColl.DeleteOne(ctx, filter)
	return result.DeletedCount, err
}
func (repo AddressRepo) UpdateAddress(filter primitive.M, address model.Address, ctx context.Context) (model.Address, error) {
	addressColl := repo.conn.AddressCollection()
	err := addressColl.FindOneAndUpdate(ctx, filter, bson.M{"$set": address}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&address)
	return address, err
}
func (repo AddressRepo) GetAddresses(ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	filter := bson.D{{}}
	addressColl := repo.conn.AddressCollection()
	cursor, err := addressColl.Find(ctx, filter, findOptions)
	return cursor, err
}

func (repo AddressRepo) SearchAddress(filter interface{}, ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	addressColl := repo.conn.AddressCollection()
	cursor, err := addressColl.Find(ctx, filter, findOptions)
	return cursor, err
}
