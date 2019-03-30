package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Friends []*User `gorm:"many2many:friendships;association_jointable_foreignkey:friend_id"`
}

func init() {
	db, err = gorm.Open("mysql", "root:123456@/gff?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&User{})
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
	var user User
	db.Preload("Friends").First(&user, "id = ?", 1)
	fmt.Println(user)
	db.Model(&user).Association("Friends").Append(&User{}, &User{})
}
