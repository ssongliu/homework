package main

import (
	"fmt"
	"homework/db"
	"homework/model"
	"log"

	"github.com/pkg/errors"
)

var ErrNoRows error

func main() {
	initDB()
	searchName := "zhangsan"
	go myService(searchName)
	select {}
}

func initDB() {
	dbi := db.DbConfig{
		Host:     "localhost",
		Port:     3306,
		Name:     "mytest",
		User:     "root",
		Password: "123456789",
	}
	err := dbi.Init()
	if err != nil {
		log.Fatal(err)
	}
}

func myService(username string) error {
	users, err := selectUsers(username)
	if err != nil {
		if errors.Cause(err) == ErrNoRows {
			fmt.Printf("select users by name err %s: no such user \n", username)
			// return
		} else {
			fmt.Printf("select users err: %v \n", err)
			// return
		}
	}
	// handle users
	fmt.Println(users)
	return nil
}

func selectUsers(username string) ([]model.User, error) {
	var users []model.User
	if err := db.DB.Where("user_name == ?", username).Find(&users).Error; err != nil {
		return users, errors.Wrap(err, "select user list by name err")
	}
	return users, nil
}
