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
	Name      string
	Age       int     `gorm:"default:18"`
	Profile   Profile `gorm:"ForeignKey:ProfileID"` // 使用ProfileID作为外键
	ProfileID int
}
type Profile struct {
	gorm.Model
	Name string
}

func init() {
	db, err = gorm.Open("mysql", "root:123456@/gff?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&User{}, &Profile{})
}

func main() {
	defer db.Close()
	var profile Profile
	var userToFind User

	if err := db.Where(User{Model: gorm.Model{ID: 1}}).First(&userToFind).Error; err != nil {
		fmt.Println(err)
	}
	if err := db.Model(&userToFind).Related(&profile).Error; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", profile)

	fmt.Printf("%+v", userToFind)

}
