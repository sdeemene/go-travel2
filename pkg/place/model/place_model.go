package model

import (
	"github.com/stdeemene/go-travel2/pkg/address/model"
	"github.com/stdeemene/go-travel2/utils/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Place struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Price       float64            `bson:"price,omitempty" json:"price,omitempty"`
	IsAvailable bool               `bson:"isAvailable,omitempty" json:"isAvailable,omitempty"`
	Phone       string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Email       string             `bson:"email,omitempty" json:"email,omitempty"`
	Address     *model.Address     `bson:"address,omitempty" json:"address,omitempty"`
	CreatedAT   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAT   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type PlaceUpdated struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        Place
}

type PlaceDeleted struct {
	DeletedCount int64 `json:"deletedCount"`
}

func (p *Place) Initialize() error {
	p.CreatedAT = helper.CurrentDate()
	p.UpdatedAT = helper.CurrentDate()
	p.IsAvailable = true
	return nil
}

func (p *Place) UpdateTimeStamp() error {
	p.UpdatedAT = helper.CurrentDate()
	return nil
}

type PlaceReq struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	AddressID   string  `json:"addressId"`
}
