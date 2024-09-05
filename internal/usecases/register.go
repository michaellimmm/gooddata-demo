package usecases

import (
	"context"
	"errors"
	"fmt"
	"github/michaellimmm/gooddata-demo/generated/analytics/v1"
	"github/michaellimmm/gooddata-demo/internal/models"
	"github/michaellimmm/gooddata-demo/internal/utils"
	"github/michaellimmm/gooddata-demo/pkg/gooddata"
	"log/slog"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type Register interface {
	Register(ctx context.Context, req *analytics.RegisterAccountRequest) (*analytics.RequestAccountResponse, error)
}

func (u *usecases) Register(
	ctx context.Context,
	req *analytics.RegisterAccountRequest,
) (*analytics.RequestAccountResponse, error) {
	err := u.validateRegisterAccountRequest(ctx, req)
	if err != nil {
		slog.Error("RegisterAccountRequest is not valid", err)
		return nil, err
	}

	newPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		slog.Error("failed to hash password", err)
		return nil, err
	}

	account := models.Account{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: newPassword,
		TenantID: req.GetTenantId(),
	}
	if err := u.repo.SaveAccount(ctx, account); err != nil {
		slog.Error("failed to save account", err)
		return nil, err
	}

	// create user
	// todo add convention for user ID
	user := gooddata.User{
		ID:        fmt.Sprintf("u_%s", req.GetTenantId()),
		Firstname: req.GetTenantId(),
		UserGroups: []*gooddata.UserGroup{
			{ID: "ug_tenant", Name: "tenant_group"},
		},
	}
	_, err = u.goodDataApi.CreateUser(user)
	if err != nil {
		slog.Error("failed to create user", err)
		return nil, err
	}

	// assign user data filter
	udf := gooddata.UserDataFilter{
		ID:          fmt.Sprintf("udf_%s", req.GetTenantId()),
		Description: fmt.Sprintf("user data filter for %s", req.GetTenantId()),
		Maql:        fmt.Sprintf(`{label/TENANT_ID} = "%s"`, req.GetTenantId()),
		User: &gooddata.User{
			ID: user.ID,
		},
	}
	workspaceID := os.Getenv("WORKSPACE_ID")
	_, err = u.goodDataApi.CreateUserDataFilter(workspaceID, udf)
	if err != nil {
		slog.Error("failed to create user data filter", err)
		return nil, err
	}

	// generate token
	kid := os.Getenv("KID")
	privateKey := os.Getenv("PRIVATE_KEY")
	token, err := GenerateToken(privateKey, TokenKey{
		Kid: kid,
		Sub: fmt.Sprintf("u_%s", req.GetTenantId()),
	})
	if err != nil {
		slog.Error("failed to generate token", err)
		return nil, err
	}

	result := analytics.RequestAccountResponse{
		Name:        account.Name,
		Email:       account.Email,
		TenantId:    account.TenantID,
		AccessToken: token,
	}

	return &result, nil
}

func (u *usecases) validateRegisterAccountRequest(ctx context.Context, req *analytics.RegisterAccountRequest) error {
	existingAccount, err := u.repo.FindAccountByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	if !existingAccount.IsEmpty() {
		return errors.New("account is already registered")
	}

	return nil
}
