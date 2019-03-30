package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

// User 包含多个 emails, UserID 为外键
type User3 struct {
	gorm.Model
	Emails []Email
}

type Email struct {
	gorm.Model
	Email   string
	User3ID uint
}

func init() {
	db, err = gorm.Open("mysql", "root:123456@/gff?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&User3{}, &Email{})
	// user := User3{Model: gorm.Model{ID: 1}}
	// email := Email{User3ID: 1, Email: "asdfsdf"}
	// db.Create(&user)
	// db.Create(&email)

}

func main() {
	defer db.Close()
	var user3 User3
	user3.ID = 1
	emails := make([]Email, 0)
	db.First(&user3).Related(&emails)
	fmt.Printf("%v", emails)
}
