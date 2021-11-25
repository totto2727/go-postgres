package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type User struct {
	Name string `query:"name"`
	Age  int    `query:"age"`
}

type user struct {
	Name string `query:"name"`
	Age  int    `query:"age"`
}

func GetUsers(query *User) []user {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=mydb password=password sslmode=disable")

	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	defer db.Close()

	var users []user
	db.Debug().Where(query).Find(&users)

	return users
}
