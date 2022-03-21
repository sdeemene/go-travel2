package repository

import (
	"context"

	dbs "github.com/stdeemene/go-travel2/config/db"
	"github.com/stdeemene/go-travel2/pkg/user"
	"github.com/stdeemene/go-travel2/pkg/user/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepo struct {
	conn *dbs.MongoDB
}

func NewUserRepo(conn *dbs.MongoDB) user.UserRepositoryImp {
	return UserRepo{conn: conn}
}

func (repo UserRepo) CreateUser(user *model.User, ctx context.Context) (interface{}, error) {
	userColl := repo.conn.UserCollection()
	result, err := userColl.InsertOne(ctx, user)
	return result.InsertedID, err
}
func (repo UserRepo) AuthenticateUser(filter primitive.M, ctx context.Context) (model.User, error) {
	return repo.GetUser(filter, ctx)
}

func (repo UserRepo) GetUser(filter primitive.M, ctx context.Context) (model.User, error) {
	var usr model.User
	userColl := repo.conn.UserCollection()
	err := userColl.FindOne(ctx, filter).Decode(&usr)
	return usr, err
}
func (repo UserRepo) DeleteUser(filter primitive.M, ctx context.Context) (int64, error) {
	userColl := repo.conn.UserCollection()
	result, err := userColl.DeleteOne(ctx, filter)
	return result.DeletedCount, err
}
func (repo UserRepo) UpdateUser(filter primitive.M, user model.User, ctx context.Context) (model.User, error) {
	userColl := repo.conn.UserCollection()
	err := userColl.FindOneAndUpdate(ctx, filter, bson.M{"$set": user}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&user)
	return user, err
}
func (repo UserRepo) GetUsers(ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	filter := bson.D{{}}
	userColl := repo.conn.UserCollection()
	cursor, err := userColl.Find(ctx, filter, findOptions)
	return cursor, err
}
func (repo UserRepo) GetUserByEmail(filter primitive.M, ctx context.Context) (model.User, error) {
	return repo.GetUser(filter, ctx)
}

func (repo UserRepo) SearchUser(filter interface{}, ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	userColl := repo.conn.UserCollection()
	cursor, err := userColl.Find(ctx, filter, findOptions)
	return cursor, err
}
