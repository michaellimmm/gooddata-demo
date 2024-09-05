package usecases

import (
	"context"
	"errors"
	"fmt"
	"github/michaellimmm/gooddata-demo/generated/analytics/v1"
	"github/michaellimmm/gooddata-demo/internal/utils"
	"log/slog"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type Login interface {
	Login(ctx context.Context, req *analytics.LoginRequest) (*analytics.LoginResponse, error)
}

func (u *usecases) Login(
	ctx context.Context,
	req *analytics.LoginRequest,
) (*analytics.LoginResponse, error) {
	account, err := u.repo.FindAccountByEmail(ctx, req.GetEmail())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("account is not found")
		}

		slog.Error("failed to find account by email", err)
		return nil, err
	}

	isValid := utils.CheckPasswordHash(req.Password, account.Password)
	if !isValid {
		return nil, errors.New("password is invalid")
	}

	kid := os.Getenv("KID")
	privateKey := os.Getenv("PRIVATE_KEY")
	token, err := GenerateToken(privateKey, TokenKey{
		Kid: kid,
		Sub: fmt.Sprintf("u_%s", account.TenantID),
	})
	if err != nil {
		slog.Error("failed to generate token", err)
		return nil, err
	}

	result := analytics.LoginResponse{
		Email:       account.Email,
		Name:        account.Name,
		TenantId:    account.TenantID,
		AccessToken: token,
	}

	return &result, nil
}
