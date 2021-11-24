package data

import (
	"context"
	"errors"
	"fmt"
	"user-service/internal/biz"
	"user-service/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type UserDataRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserDataRepo(data *Data, logger log.Logger) biz.UserRepo {
	return UserDataRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (c UserDataRepo) CreateUser(ctx context.Context, user *biz.User) error {
	//	c.data.db.Raw()
	// c.data.db.Raw
	//1.查看用户是否存在数据库

	//c.data.db.Raw("insert into user(nickname,pwd,avatarurl,country,province,city,uname) value()")
	//创建用户
	tx := c.data.db.Model(&model.User{}).Create(&model.User{
		NickName:  user.NickName,
		AvatarUrl: user.AvatarUrl,
		Pwd:       user.Pwd,
		Country:   user.Country,
		Province:  user.Province,
		City:      user.City,
		Uname:     user.Uname,
		Status:    user.Status,
		Id:        user.Id,
	})
	if tx.Error != nil {
		fmt.Println("err-------------")
		fmt.Println(tx.Error)
		return tx.Error
	}
	return nil
}
func (c UserDataRepo) GetUser(ctx context.Context, nickname string, pwd string) map[string]interface{} {
	// c.data.db
	var tmap []map[string]interface{}
	tx := c.data.db.Raw("select * from user where nickname=? and pwd=?", nickname, pwd).Scan(&tmap)
	if tx.Error != nil {
		fmt.Println("查询用户失败!-------")
	}
	return tmap[0]

}
func (c UserDataRepo) UserIsOk(ctx context.Context, nickname string) error {
	var tmap []map[string]interface{}
	tx := c.data.db.Raw("select * from user where nickname=?", nickname).Scan(&tmap)
	if tx.Error != nil {
		fmt.Println("查询用户失败!-------")
		return errors.New("当前查询用户出错！")
	}
	v := tmap[0]
	v["isTrue"] = true
	if len(tmap) != 0 {
		return nil
	}
	return errors.New("当前用户不存在,可以注册该用户名!")
}

//获取用户列表
func (c UserDataRepo) GetUserList(ctx context.Context, page int64, pageSize int64) ([]*biz.User, int64, error) {
	var list []model.User
	spage := (page - 1) * pageSize
	// c.data.db
	tx := c.data.db.Raw("select * from user a where limit ?,?", spage, pageSize).Scan(&list)
	if tx.Error != nil {
		//查询失败
	}
	var count int64
	c.data.db.Count(&count)
	var res []*biz.User
	for _, v := range list {
		res = append(res, &biz.User{
			Id:        v.Id,
			NickName:  v.NickName,
			Pwd:       v.Pwd,
			AvatarUrl: v.AvatarUrl,
			Country:   v.Country,
			Province:  v.Province,
			City:      v.City,
			Uname:     v.Uname,
			Status:    v.Status,
		})
	}
	return res, count, nil

}

func (c UserDataRepo) UserAddDress(ctx context.Context, phone, dress string) error {
	return c.data.db.Model(&model.Dress{}).Create(&model.Dress{
		Id:     int64(123),
		UserId: int64(3434),
		Phone:  phone,
		Dress:  dress,
		Status: int64(1),
	}).Error
}
func (c UserDataRepo) GetUserAddDressList(ctx context.Context, page, pageSize int64) ([]*biz.Dress, error) {
	var m []*biz.Dress
	c.data.db.Raw("select * from dress where limit ?,?", page, pageSize).Scan(m)
	return m, nil
}
func (c UserDataRepo) DressUpdate(ctx context.Context, id int64, phone string, dress string) error {
	var d []*biz.Dress
	tx := c.data.db.Raw("update user set phone=?,dress=? where id=?", phone, dress, id).Scan(d)
	if tx.Error != nil {
		return errors.New("修改失败!")
	}
	return nil
}
