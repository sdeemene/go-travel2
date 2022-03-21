package service

import (
	"context"
	"fmt"
	"github.com/stdeemene/go-travel2/utils/helper"

	"github.com/stdeemene/go-travel2/pkg/address"
	"github.com/stdeemene/go-travel2/pkg/address/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddressService struct {
	addressRepo address.AddressRepositoryImp
}

func NewAddressService(addressRepo address.AddressRepositoryImp) address.AddressServiceImp {
	return AddressService{addressRepo: addressRepo}
}

func (service AddressService) CreateAddress(address model.Address, ctx context.Context) (model.Address, error) {
	result, err := service.addressRepo.CreateAddress(address, ctx)
	if err != nil {
		return address, err
	}
	id := result.(primitive.ObjectID).Hex()
	fmt.Println("inserted id ", id)
	return service.GetAddress(id, ctx)
}

func (service AddressService) GetAddress(id string, ctx context.Context) (model.Address, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	getAddress, err := service.addressRepo.GetAddress(filter, ctx)
	if err != nil {
		return getAddress, err
	}
	fmt.Println("Record Found: ", getAddress)
	return getAddress, nil
}
func (service AddressService) DeleteAddress(id string, ctx context.Context) (model.AddressDeleted, error) {
	result := model.AddressDeleted{
		DeletedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	if err != nil {
		return result, err
	}

	_, err = service.GetAddress(id, ctx)
	if err != nil {
		return result, err
	}
	res, err := service.addressRepo.DeleteAddress(filter, ctx)
	if err != nil {
		return result, err
	}
	result.DeletedCount = res
	return result, nil
}

func (service AddressService) UpdateAddress(id string, address model.Address, ctx context.Context) (model.AddressUpdated, error) {
	result := model.AddressUpdated{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	filter := bson.M{"_id": _id}

	_, err = service.GetAddress(id, ctx)
	if err != nil {
		return result, err
	}

	res, err := service.addressRepo.UpdateAddress(filter, address, ctx)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = 1
	result.Result = res
	fmt.Println("Record updated: ", result)
	return result, nil

}

func (service AddressService) GetAddresses(ctx context.Context) ([]model.Address, error) {
	var addresses []model.Address
	cursor, err := service.addressRepo.GetAddresses(ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &addresses)
	if err != nil {
		return nil, err
	}
	if len(addresses) == 0 {
		return nil, helper.NewError("no result found")
	}
	return addresses, nil
}

func (service AddressService) SearchAddress(filter interface{}, ctx context.Context) ([]model.Address, error) {
	if filter == nil {
		filter = bson.M{}
	}
	var addresses []model.Address
	cursor, err := service.addressRepo.SearchAddress(filter, ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &addresses)
	if err != nil {
		return nil, err
	}
	if len(addresses) == 0 {
		return nil, helper.NewError("no result found")
	}
	return addresses, nil
}
