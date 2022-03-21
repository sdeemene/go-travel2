package model

import (
	placeModel "github.com/stdeemene/go-travel2/pkg/place/model"
	userModel "github.com/stdeemene/go-travel2/pkg/user/model"
	"github.com/stdeemene/go-travel2/utils/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Travel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Duration    string             `bson:"duration,omitempty" json:"duration,omitempty"`
	TotalAmount float64            `bson:"totalAmount,omitempty" json:"totalAmount,omitempty"`
	User        *userModel.User    `bson:"user,omitempty" json:"user,omitempty"`
	Place       *placeModel.Place  `bson:"place,omitempty" json:"place,omitempty"`
	TravelAT    time.Time          `bson:"travelAt,omitempty" json:"travelAt,omitempty"`
	ReturnAT    time.Time          `bson:"returnAt,omitempty" json:"returnAt,omitempty"`
	RateValue   int64              `bson:"rateValue" json:"rateValue,omitempty"`
	CreatedAT   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAT   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type TravelUpdated struct {
	ModifiedCount int64   `json:"modifiedCount"`
	Result        *Travel `json:"travelData"`
}

type TravelDeleted struct {
	DeletedCount int64 `json:"deletedCount"`
}

func (t *Travel) Initialize() error {
	t.CreatedAT = helper.CurrentDate()
	t.UpdatedAT = helper.CurrentDate()
	t.RateValue = 0
	return nil
}

type TravelReq struct {
	UserID   string `json:"userId"`
	PlaceID  string `json:"placeId"`
	TravelAT string `json:"travelDate"`
	ReturnAT string `json:"returnDate"`
}

type TravelRateReq struct {
	RateValue int `json:"rateValue"`
}
