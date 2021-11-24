package data

import (
	"fmt"
	"user-service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserDataRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	// conf.Data.Database.Source
	// db, err := gorm.Open(mysql.Open(conf.Data_Database.Source, &gorm.Config{}))
	// db, err := gorm.Open(mysql.Open(conf.Data_Database), &gorm.Config{})
	fmt.Println("准备链接数据库----------")

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	d := &Data{
		db: db,
	}

	return d, func() {
		log.Info("messge", "closing the data resources")
	}, nil
}
