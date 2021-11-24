package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id        int64
	NickName  string
	Pwd       string
	AvatarUrl string
	Country   string
	Province  string
	City      string
	Uname     string
	Status    int64
}
type Dress struct {
	Id     int64
	UserId int64
	Phone  string
	Dress  string
	Status int64
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserList(ctx context.Context, page int64, pageSize int64) ([]*User, int64, error)
	GetUser(ctx context.Context, nickname string, pwd string) map[string]interface{}
	UserIsOk(ctx context.Context, nickname string) error
	UserAddDress(ctx context.Context, phone, dress string) error
	GetUserAddDressList(ctx context.Context, page, pageSize int64) ([]*Dress, error)
	DressUpdate(ctx context.Context, id int64, phone, dress string) error
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

//用户列表
func (c *UserUseCase) List(ctx context.Context, page int64, pageSize int64) ([]*User, int64, error) {
	// p := make(map[string]interface{})
	// p["status"] = "ok"
	// count := 2
	return c.repo.GetUserList(ctx, page, pageSize)
}

//创建用户
func (c *UserUseCase) Create(ctx context.Context, user *User) error {
	return c.repo.CreateUser(ctx, user)
}

func (c *UserUseCase) GetUser(ctx context.Context, nickname string, pwd string) map[string]interface{} {
	return c.repo.GetUser(ctx, nickname, pwd)
}
func (c *UserUseCase) UserIsOk(ctx context.Context, nickname string) error {
	return c.repo.UserIsOk(ctx, nickname)
}

//添加地址
func (c *UserUseCase) UserAddDress(ctx context.Context, phone, dress string) error {
	return c.repo.UserAddDress(ctx, phone, dress)
}

//DressUpdate
func (c *UserUseCase) GetUserAddDressList(ctx context.Context, page, pageSize int64) ([]*Dress, error) {
	return c.repo.GetUserAddDressList(ctx, page, pageSize)
}

func (c *UserUseCase) DressUpdate(ctx context.Context, id int64, phone string, dress string) error {
	return c.repo.DressUpdate(ctx, id, phone, dress)
}
