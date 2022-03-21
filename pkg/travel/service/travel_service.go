package service

import (
	"context"
	"github.com/hako/durafmt"
	"github.com/stdeemene/go-travel2/pkg/place"
	"github.com/stdeemene/go-travel2/pkg/travel"
	"github.com/stdeemene/go-travel2/pkg/travel/model"
	"github.com/stdeemene/go-travel2/pkg/user"
	"github.com/stdeemene/go-travel2/utils/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type TravelService struct {
	travelRepo   travel.TravelRepositoryImp
	userService  user.UserServiceImp
	placeService place.PlaceServiceImp
}

func NewTravelService(travelRepo travel.TravelRepositoryImp, userService user.UserServiceImp, placeService place.PlaceServiceImp) travel.TravelServiceImp {
	return TravelService{travelRepo: travelRepo, userService: userService, placeService: placeService}
}

func (service TravelService) CreateTravel(payload *model.TravelReq, ctx context.Context) (*model.Travel, error) {
	travelDate, err := helper.ConvertToDate(payload.TravelAT)
	if err != nil {
		return nil, err
	}

	returnDate, err := helper.ConvertToDate(payload.ReturnAT)
	if err != nil {
		return nil, err
	}

	newTravel := &model.Travel{
		TravelAT: travelDate,
		ReturnAT: returnDate,
	}
	err = newTravel.Initialize()
	if err != nil {
		return nil, err
	}

	getUser, err := service.userService.GetUser(payload.UserID, ctx)
	if err != nil {
		return nil, err
	}
	newTravel.User = getUser

	interval := returnDate.Sub(travelDate)
	intervalInDays := interval.Hours() / 24
	duration, err := durafmt.ParseString(interval.String())
	if err != nil {
		return nil, err
	}
	newTravel.Duration = duration.String()

	getPlace, err := service.placeService.GetPlace(payload.PlaceID, ctx)
	if err != nil {
		return nil, err
	}
	newTravel.Place = getPlace

	calculatedAmount := intervalInDays * getPlace.Price
	amount := strconv.FormatFloat(calculatedAmount, 'f', 2, 64)
	totalAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return nil, err
	}
	newTravel.TotalAmount = totalAmount

	result, err := service.travelRepo.CreateTravel(newTravel, ctx)
	if err != nil {
		return nil, err
	}
	id := result.(primitive.ObjectID).Hex()
	return service.GetTravel(id, ctx)
}
func (service TravelService) GetTravel(id string, ctx context.Context) (*model.Travel, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	getPlace, err := service.travelRepo.GetTravel(filter, ctx)
	if err != nil {
		return nil, err
	}
	return &getPlace, nil
}
func (service TravelService) DeleteTravel(id string, ctx context.Context) (model.TravelDeleted, error) {
	result := model.TravelDeleted{
		DeletedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.M{"_id": _id}
	_, err = service.GetTravel(id, ctx)
	if err != nil {
		return result, err
	}
	res, err := service.travelRepo.DeleteTravel(filter, ctx)
	if err != nil {
		return result, err
	}
	result.DeletedCount = res
	return result, nil
}
func (service TravelService) UpdateTravel(id string, place model.Travel, ctx context.Context) (model.TravelUpdated, error) {
	result := model.TravelUpdated{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.M{"_id": _id}

	_, err = service.GetTravel(id, ctx)
	if err != nil {
		return result, err
	}

	res, err := service.travelRepo.UpdateTravel(filter, place, ctx)
	if err != nil {
		return result, nil
	}
	result.ModifiedCount = 1
	result.Result = &res
	return result, nil
}
func (service TravelService) GetTravels(ctx context.Context) ([]model.Travel, error) {
	var travels []model.Travel
	cursor, err := service.travelRepo.GetTravels(ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &travels)
	if err != nil {
		return nil, err
	}
	if len(travels) == 0 {
		return nil, helper.NewError("no result found")
	}
	return travels, nil
}
func (service TravelService) SearchTravel(filter interface{}, ctx context.Context) ([]model.Travel, error) {
	var travels []model.Travel
	cursor, err := service.travelRepo.SearchTravel(filter, ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &travels)
	if err != nil {
		return nil, err
	}
	if len(travels) == 0 {
		return nil, helper.NewError("no result found")
	}
	return travels, nil
}
func (service TravelService) GetUserTravels(id string, ctx context.Context) ([]model.Travel, error) {
	getUser, err := service.userService.GetUser(id, ctx)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"user": getUser}
	var travels []model.Travel
	cursor, err := service.travelRepo.GetUserTravels(filter, ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &travels)
	if err != nil {
		return nil, err
	}
	if len(travels) == 0 {
		return nil, helper.NewError("no result found")
	}
	return travels, nil
}
func (service TravelService) RateTravel(id string, rateValue *model.TravelRateReq, ctx context.Context) (model.TravelUpdated, error) {
	rating := &model.TravelRateReq{
		RateValue: rateValue.RateValue,
	}
	result := model.TravelUpdated{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.M{"_id": _id}

	_, err = service.GetTravel(id, ctx)
	if err != nil {
		return result, err
	}
	update := bson.M{"$set": bson.M{"rateValue": rating.RateValue}}

	res, err := service.travelRepo.RateTravel(filter, update, ctx)
	if err != nil {
		return result, err
	}

	getTravel, err := service.GetTravel(id, ctx)
	if err != nil {
		return result, nil
	}
	result.ModifiedCount = res
	result.Result = getTravel
	return result, nil
}
