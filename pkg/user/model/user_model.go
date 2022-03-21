package model

import (
	"time"

	"github.com/google/uuid"
	constant "github.com/stdeemene/go-travel2/config/constant"
	helper "github.com/stdeemene/go-travel2/utils/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Firstname string             `bson:"firstname,omitempty" json:"firstname,omitempty"`
	Lastname  string             `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Gender    string             `bson:"gender,omitempty" json:"gender,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	Phone     string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Password  string             `bson:"password,omitempty" json:"-"`
	//Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	Role      string    `bson:"role,omitempty" json:"role,omitempty"`
	Salt      string    `bson:"salt,omitempty" json:"salt,omitempty"`
	IsActive  bool      `bson:"isActive,omitempty" json:"isActive,omitempty"`
	CreatedAT time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAT time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type UserUpdated struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        User
}

type UserDeleted struct {
	DeletedCount int64 `json:"deletedCount"`
}

func (u *User) ComparePassword(password string) error {
	incoming := []byte(password + u.Salt)
	existing := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err
}

func (u *User) Initialize() error {
	salt := uuid.New().String()
	passwordBytes := []byte(u.Password + salt)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash[:])
	u.Salt = salt
	u.CreatedAT = helper.CurrentDate()
	u.UpdatedAT = helper.CurrentDate()
	u.IsActive = true
	u.Role = constant.UserRole
	return nil
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshReq struct {
	Token string `json:"token"`
}

type SignupReq struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}
