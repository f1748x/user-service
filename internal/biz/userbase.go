package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type Userl struct {
	Id   int64
	Name string
}

type UserRepol interface {
	CreateUser(ctx context.Context, user *User) error
}

type UserUseCases struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCases(repo UserRepo, logger log.Logger) *UserUseCase {
	fmt.Println("----------------biz")
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCases) Create(ctx context.Context, user *User) error {
	fmt.Println("----------------biz")
	return uc.repo.CreateUser(ctx, user)
}
