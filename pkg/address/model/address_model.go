package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Country string             `bson:"country,omitempty" json:"country,omitempty"`
	State   string             `bson:"state,omitempty" json:"state,omitempty"`
	City    string             `bson:"city,omitempty" json:"city,omitempty"`
}

type AddressUpdated struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        Address
}

type AddressDeleted struct {
	DeletedCount int64 `json:"deletedCount"`
}
