package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

// Use pointer value
type User struct {
	gorm.Model
	Name string
	Age  int `gorm:"default:18"`
}
type Profile struct {
	gorm.Model
	UserID int
	User   User
	Name   string
}

func init() {
	db, err = gorm.Open("mysql", "root:123456@/gff?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&User{})
}
func main() {

	user := User{Name: "Jinzhu", Age: 18}

	db.NewRecord(user) // => returns `true` as primary key is blank

	db.Create(&user)

	db.NewRecord(user) // => return `false` after `user` created
	defer db.Close()
}
