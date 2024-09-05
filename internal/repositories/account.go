package repositories

import (
	"context"
	"github/michaellimmm/gooddata-demo/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	accountCollection = "account"
)

type AccountRepositories interface {
	FindAccountByEmail(email string) (models.Account, error)
	SaveAccount(ctx context.Context, account models.Account) error
}

func (r *Repositories) FindAccountByEmail(ctx context.Context, email string) (models.Account, error) {
	result := models.Account{}

	filter := filterNotDeleted()
	filter["email"] = email

	col := r.getCollection(accountCollection)
	err := col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *Repositories) SaveAccount(ctx context.Context, account models.Account) error {
	if account.ID.IsZero() {
		account.ID = primitive.NewObjectID()
	}

	now := time.Now()
	if account.CreatedAt.IsZero() {
		account.CreatedAt = now
	}

	if account.UpdatedAt.IsZero() {
		account.UpdatedAt = now
	}

	col := r.getCollection(accountCollection)
	_, err := col.InsertOne(ctx, account)
	if err != nil {
		return err
	}

	return nil
}
