package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username, Password string
}

func main() {
	dsn := "host=127.0.0.1 user=busrai password=123 dbname=go_blog port=5432"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{})

	db.Create(&User{Username: "Büşra", Password: "123"})
	var user User
	db.First(&user, "1")
	db.First(&user, "password = ?", "123")
	fmt.Println(user.Username)

	var users []User
	db.Find(&users, "password = ?", "123")
	fmt.Println(users)

	db.First(&user, 1)
	db.Model(&user).Update("username", "gorm")

	db.Model(&user).Updates(User{Username: "busra", Password: "1234567"})

	db.Delete(&user, 1)
}
