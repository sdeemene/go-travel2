package service

import (
	"context"
	"github.com/stdeemene/go-travel2/pkg/address"
	"github.com/stdeemene/go-travel2/pkg/place"
	"github.com/stdeemene/go-travel2/pkg/place/model"
	"github.com/stdeemene/go-travel2/utils/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlaceService struct {
	placeRepo      place.PlaceRepositoryImp
	addressService address.AddressServiceImp
}

func NewPlaceService(placeRepo place.PlaceRepositoryImp, service address.AddressServiceImp) place.PlaceServiceImp {
	return PlaceService{placeRepo: placeRepo, addressService: service}
}

func (service PlaceService) CreatePlace(payload *model.PlaceReq, ctx context.Context) (*model.Place, error) {
	newPlace := &model.Place{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Email:       payload.Email,
		Phone:       payload.Phone,
	}
	err := newPlace.Initialize()
	if err != nil {
		return nil, err
	}
	getAddress, err := service.addressService.GetAddress(payload.AddressID, ctx)
	if err != nil {
		return nil, err
	}
	newPlace.Address = &getAddress

	result, err := service.placeRepo.CreatePlace(newPlace, ctx)
	if err != nil {
		return nil, err
	}
	id := result.(primitive.ObjectID).Hex()
	return service.GetPlace(id, ctx)
}
func (service PlaceService) GetPlace(id string, ctx context.Context) (*model.Place, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	getPlace, err := service.placeRepo.GetPlace(filter, ctx)
	if err != nil {
		return nil, err
	}
	return &getPlace, nil
}
func (service PlaceService) DeletePlace(id string, ctx context.Context) (model.PlaceDeleted, error) {
	result := model.PlaceDeleted{
		DeletedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.M{"_id": _id}
	_, err = service.GetPlace(id, ctx)
	if err != nil {
		return result, err
	}
	res, err := service.placeRepo.DeletePlace(filter, ctx)
	if err != nil {
		return result, err
	}
	result.DeletedCount = res
	return result, nil
}
func (service PlaceService) UpdatePlace(id string, place model.Place, ctx context.Context) (model.PlaceUpdated, error) {
	result := model.PlaceUpdated{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.M{"_id": _id}

	_, err = service.GetPlace(id, ctx)
	if err != nil {
		return result, err
	}
	err = place.UpdateTimeStamp()
	if err != nil {
		return result, err
	}
	res, err := service.placeRepo.UpdatePlace(filter, place, ctx)
	if err != nil {
		return result, nil
	}
	result.ModifiedCount = 1
	result.Result = res
	return result, nil
}
func (service PlaceService) GetPlaces(ctx context.Context) ([]model.Place, error) {
	var places []model.Place
	cursor, err := service.placeRepo.GetPlaces(ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &places)
	if err != nil {
		return nil, err
	}
	if len(places) == 0 {
		return nil, helper.NewError("no result found")
	}
	return places, nil
}
func (service PlaceService) SearchPlace(filter interface{}, ctx context.Context) ([]model.Place, error) {
	var places []model.Place
	cursor, err := service.placeRepo.SearchPlace(filter, ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &places)
	if err != nil {
		return nil, err
	}
	if len(places) == 0 {
		return nil, helper.NewError("no result found")
	}
	return places, nil
}
