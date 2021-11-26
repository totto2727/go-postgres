package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=mydb password=password sslmode=disable")

	if err != nil {
		log.Fatalln("接続失敗", err)
	}

	db.LogMode(true)

	return db
}

type User struct {
	Name string `json:"name" query:"name"`
	Age  int    `json:"age" query:"age"`
}

func GetUsers(query *User) []User {
	db := connectDB()
	defer db.Close()

	var users []User
	db.Where(query).Find(&users)
	// db.Debug().Where(query).Find(&users)

	return users
}

func CreateUser(query *User) {
	db := connectDB()
	defer db.Close()

	db.Create(query)
	// db.Debug().Create(query)
}

func UpdateUsers(query *User) {
	db := connectDB()
	defer db.Close()

	db.Model(&User{}).Where(User{Name: "test"}).Select("age").Updates(User{Age: 1000})
}

func DeleteUser(query *User) {
	db := connectDB()
	defer db.Close()

	db.Where(query).Delete(User{})
	// db.Debug().Where(query).Delete(User{})
}
