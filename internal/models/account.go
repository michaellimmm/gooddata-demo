package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	TenantID  string             `bson:"tenant_id"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}

func (a Account) IsEmpty() bool {
	return a.ID.IsZero() &&
		a.Name == "" &&
		a.Email == "" &&
		a.Password == "" &&
		a.TenantID == "" &&
		a.CreatedAt.IsZero() &&
		a.UpdatedAt.IsZero() &&
		a.DeletedAt == nil
}
