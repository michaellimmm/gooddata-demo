package usecases

import (
	"github/michaellimmm/gooddata-demo/internal/repositories"
	"github/michaellimmm/gooddata-demo/pkg/gooddata"
)

type Usecases interface {
	Register
	Login
}

type usecases struct {
	repo        repositories.Repositories
	goodDataApi gooddata.GooddataAPI
}

func NewUsecases(repo repositories.Repositories, goodDataApi gooddata.GooddataAPI) Usecases {
	return &usecases{
		repo:        repo,
		goodDataApi: goodDataApi,
	}
}
