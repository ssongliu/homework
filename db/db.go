package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DbConfig struct {
	Host         string
	Port         int
	Name         string
	User         string
	Password     string
	MaxOpenConns int
	MaxIdleConns int
}

func (i *DbConfig) Init() error {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Asia%%2FShanghai",
		i.User,
		i.Password,
		i.Host,
		i.Port,
		i.Name)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return err
	}

	db.SingularTable(true)
	db.DB().SetMaxOpenConns(i.MaxOpenConns)
	db.DB().SetMaxIdleConns(i.MaxIdleConns)
	DB = db
	DB.LogMode(false)
	return nil
}
