package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type User2 struct {
	gorm.Model
	Name string
}

type Profile2 struct {
	gorm.Model
	Name    string
	User2   User2
	User2ID int
}

// `gorm:"association_foreignkey:ID;foreignkey:User2ID"`
func init() {
	db, err = gorm.Open("mysql", "root:123456@/gff?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&User2{}, &Profile2{})
	// profileTest := Profile2{Name: "asdfasdfasdfsfsdfasdfsf", User2ID: 1}
	// db.Create(&profileTest)
	// user := User2{Name: "Jinzhu"}
	// db.Create(&user)
}

func main() {
	var profile Profile2
	var userToFind User2
	userToFind.ID = 1

	if err := db.First(&userToFind).Related(&profile).Error; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", profile)
}
