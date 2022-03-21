package user

import (
	"context"

	"github.com/stdeemene/go-travel2/pkg/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImp interface {
	CreateUser(user *model.User, ctx context.Context) (interface{}, error)
	AuthenticateUser(filter primitive.M, ctx context.Context) (model.User, error)
	GetUser(filter primitive.M, ctx context.Context) (model.User, error)
	DeleteUser(filter primitive.M, ctx context.Context) (int64, error)
	UpdateUser(filter primitive.M, user model.User, ctx context.Context) (model.User, error)
	GetUsers(ctx context.Context) (*mongo.Cursor, error)
	GetUserByEmail(filter primitive.M, ctx context.Context) (model.User, error)
	SearchUser(filter interface{}, ctx context.Context) (*mongo.Cursor, error)
}

type UserServiceImp interface {
	CreateUser(payload *model.SignupReq, ctx context.Context) (*model.User, error)
	AuthenticateUser(credentials *model.LoginReq, ctx context.Context) (*model.User, error)
	GetUser(id string, ctx context.Context) (*model.User, error)
	DeleteUser(id string, ctx context.Context) (model.UserDeleted, error)
	UpdateUser(id string, user model.User, ctx context.Context) (model.UserUpdated, error)
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUserByEmail(email string, ctx context.Context) (*model.User, error)
	SearchUser(filter interface{}, ctx context.Context) ([]model.User, error)
}
