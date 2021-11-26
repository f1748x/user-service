package service

import (
	"context"
	"fmt"

	log "github.com/go-kratos/kratos/v2/log"

	pb "user-service/api/user"
	"user-service/internal/biz"
)

type UserService struct {
	// pb.UnimplementeduserServer
	pb.UnimplementedUserServer
	u   *biz.UserUseCase
	log *log.Helper
}

func NewuserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		u:   uc,
		log: log.NewHelper(logger),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	//创建用户
	res := &pb.CreateUserRes{}
	create := &biz.User{}
	create.NickName = req.Nickname
	create.AvatarUrl = req.AvatarUrl
	create.City = req.City
	create.Province = req.Province
	create.Pwd = req.Pwd
	create.Uname = req.Uname
	res.Msg = "请求成功!"
	err := s.u.Create(ctx, create)
	if err != nil {
		return &pb.CreateUserRes{Ok: "false", Msg: err.Error()}, err
	}
	return &pb.CreateUserRes{Ok: "true", Msg: "创建成功!"}, nil
}
func (s *UserService) GetUserList(ctx context.Context, req *pb.GetUserListReq) (*pb.GetUserListRes, error) {
	// page := req.Page
	// pageSize := req.PageSize
	var r []*pb.User
	list, count, _ := s.u.List(ctx, req.Page, req.PageSize)
	for _, v := range list {
		r = append(r, &pb.User{
			Id:        v.Id,
			Nickname:  v.NickName,
			Pwd:       v.Pwd,
			AvatarUrl: v.AvatarUrl,
			Country:   v.Country,
			Province:  v.Province,
			City:      v.City,
			Uname:     v.Uname,
			Status:    v.Status,
		})
	}
	return &pb.GetUserListRes{List: r, Count: count, Ok: "ok", Msg: "查询用户列表成功!"}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserRes, error) {
	tmap := s.u.GetUser(ctx, req.Nickname, req.Pwd)
	m := &pb.User{}
	m.AvatarUrl = tmap["avatarurl"].(string)
	m.Nickname = tmap["nickname"].(string)
	m.Id = tmap["id"].(int64)
	m.City = tmap["city"].(string)
	m.Country = tmap["country"].(string)
	m.Province = tmap["province"].(string)
	m.Status = tmap["status"].(int64)
	m.Uname = tmap["uname"].(string)
	fmt.Println("进入了user-service----------")
	return &pb.GetUserRes{UserDetail: m, Msg: "获取成功!", Ok: "ok"}, nil
}
func (s *UserService) UserIsOk(ctx context.Context, req *pb.GetUserIsOkReq) (*pb.GetUserIsOkRes, error) {
	err := s.u.UserIsOk(ctx, req.Nickname)
	if err != nil {
		return &pb.GetUserIsOkRes{IsOk: false, Msg: "当前用户不存在!"}, err
	}
	return &pb.GetUserIsOkRes{
		IsOk: true,
		Msg:  "当前用户存在!",
	}, nil
}

//添加地址
func (s *UserService) DressPush(ctx context.Context, req *pb.DressPushReq) (*pb.DressPushReply, error) {
	err := s.u.UserAddDress(ctx, req.Phone, req.Dress)
	if err != nil {
		return &pb.DressPushReply{Ok: "false", Msg: err.Error()}, nil
	}
	return &pb.DressPushReply{Ok: "true", Msg: "获取地址成功！"}, nil
}

//	获取地址列表
func (s *UserService) DressList(ctx context.Context, req *pb.DressListReq) (*pb.DressListReply, error) {
	m, err := s.u.GetUserAddDressList(ctx, req.Page, req.PageSize)

	if err != nil {
		return &pb.DressListReply{
			//Result: &pb.Dress{},
			Ok:  "false",
			Msg: "获取失败!",
		}, nil
	}
	var a []*pb.Dress
	for _, v := range m {
		a = append(a, &pb.Dress{
			Id:     v.Id,
			Phone:  v.Phone,
			Dress:  v.Dress,
			Status: v.Status,
			Msg:    "卧槽!",
		})
	}
	return &pb.DressListReply{
		Result: a,
		Msg:    "获取列表成功1",
		Ok:     "true",
	}, nil
}
func (s *UserService) DressUpdate(ctx context.Context, req *pb.DressUpdateReq) (*pb.DressUpdateReply, error) {
	return &pb.DressUpdateReply{}, nil
}
