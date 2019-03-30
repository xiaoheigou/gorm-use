package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
	gorm.Model
	Languages []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}

func init() {
	db, err = gorm.Open("mysql", "root:123456@/gff?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&User{}, &Language{})
}

func main() {
	defer db.Close()
	// 1.first drop examle tables form last operation
	// db.DropTableIfExists("users", "profiles")
	// 2. create Profile1 ,User1
	// profileTest := Profile{Name: "heheh"}
	// db.Create(&profileTest)
	// user := User{Name: "Jinzhu", Age: 18, ProfileID: 1}
	// db.Create(&user)

	// var profile Profile
	// var userToFind User
	// userToFind.ID = 1

	// if err := db.First(&userToFind).Related(&profile, "Profile").Error; err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// // fmt.Println(userToFind)

	// fmt.Printf("%+v\n", profile)

	var users []User
	language := Language{}

	db.First(&language, "id = ?", 1)
	if err := db.Model(&language).Related(&users, "Users").Error; err != nil {
		fmt.Println(err)
	}
	//// SELECT * FROM "users" INNER JOIN "user_languages" ON "user_languages"."user_id" = "users"."id" WHERE  ("user_languages"."language_id" IN ('111'))

	fmt.Printf("%+v", users)
}
