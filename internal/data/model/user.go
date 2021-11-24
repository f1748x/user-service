package model

type User struct {
	Id        int64  `gorm:"column:id;" json:"id" form:"id"`
	NickName  string `gorm:"column:nickname;" json:"nickname" form:"nickname" binding:"required"`
	Pwd       string `gorm:"column:pwd;" json:"pwd" form:"pwd" binding:"required"`
	AvatarUrl string `gorm:"column:avatarurl;" json:"avatarurl" form:"avatarurl" binding:"required"`
	Country   string
	Province  string
	City      string
	Uname     string
	Status    int64
}

func (User) TableName() string {
	return "user"
}
