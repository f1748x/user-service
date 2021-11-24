package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Dresse struct {
	Id     int64
	UserId int64
	Phone  string
	Dress  string
	Status int64
}

type DressRepo interface {
	//地址添加
	DressPush(ctx context.Context, phone string, dress string) error
	DressList(ctx context.Context, page int64, pageSize int64) error
	DressUpdate(ctx context.Context, id int64, phone string, dress string) error
}

type DressUseCase struct {
	repo DressRepo
	log  *log.Helper
}

func NewDressUseCase(repo DressRepo, logger log.Logger) *DressUseCase {
	return &DressUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
