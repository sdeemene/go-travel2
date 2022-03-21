package service

import (
	"context"
	"fmt"

	"github.com/stdeemene/go-travel2/pkg/user"
	"github.com/stdeemene/go-travel2/pkg/user/model"
	"github.com/stdeemene/go-travel2/utils/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRepo user.UserRepositoryImp
}

func NewUserService(userRepo user.UserRepositoryImp) user.UserServiceImp {
	return UserService{userRepo: userRepo}
}

func (service UserService) CreateUser(payload *model.SignupReq, ctx context.Context) (*model.User, error) {
	newUser := &model.User{
		Firstname: payload.Firstname,
		Lastname:  payload.Lastname,
		Gender:    payload.Gender,
		Email:     payload.Email,
		Password:  payload.Password,
		Phone:     payload.Phone,
	}

	err := newUser.Initialize()
	if err != nil {
		return nil, err
	}
	result, err := service.userRepo.CreateUser(newUser, ctx)

	if err != nil {
		return nil, err
	}
	id := result.(primitive.ObjectID).Hex()
	fmt.Println("Record Created: ", result)

	return service.GetUser(id, ctx)
}
func (service UserService) AuthenticateUser(credentials *model.LoginReq, ctx context.Context) (*model.User, error) {
	filter := bson.M{"email": credentials.Email}
	authenticateUser, err := service.userRepo.AuthenticateUser(filter, ctx)

	if err != nil {
		return nil, helper.NewError("invalid email")
	}
	err = authenticateUser.ComparePassword(credentials.Password)
	if err != nil {
		return nil, helper.NewError("invalid password")
	}
	return &authenticateUser, nil
}

func (service UserService) GetUser(id string, ctx context.Context) (*model.User, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	getUser, err := service.userRepo.GetUser(filter, ctx)
	if err != nil {
		return nil, err
	}
	return &getUser, nil
}

func (service UserService) DeleteUser(id string, ctx context.Context) (model.UserDeleted, error) {
	result := model.UserDeleted{
		DeletedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	if err != nil {
		return result, err
	}
	_, err = service.GetUser(id, ctx)
	if err != nil {
		return result, err
	}

	res, err := service.userRepo.DeleteUser(filter, ctx)
	if err != nil {
		return result, err
	}
	result.DeletedCount = res
	return result, nil
}
func (service UserService) UpdateUser(id string, address model.User, ctx context.Context) (model.UserUpdated, error) {
	result := model.UserUpdated{
		ModifiedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	filter := bson.M{"_id": _id}

	_, err = service.GetUser(id, ctx)
	if err != nil {
		return result, err
	}

	res, err := service.userRepo.UpdateUser(filter, address, ctx)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = 1
	result.Result = res
	fmt.Println("Record updated: ", result)
	return result, nil
}
func (service UserService) GetUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	cursor, err := service.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, helper.NewError("no result found")
	}
	return users, nil
}

func (service UserService) GetUserByEmail(email string, ctx context.Context) (*model.User, error) {
	fmt.Println("email", email)
	filter := bson.M{"email": email}
	userByEmail, err := service.userRepo.GetUserByEmail(filter, ctx)
	if err != nil {
		return nil, err
	}
	return &userByEmail, nil
}

func (service UserService) SearchUser(filter interface{}, ctx context.Context) ([]model.User, error) {
	var users []model.User
	cursor, err := service.userRepo.SearchUser(filter, ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, helper.NewError("no result found")
	}
	return users, nil
}
