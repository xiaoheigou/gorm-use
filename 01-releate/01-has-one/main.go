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
	// 1.first drop examle tables form last operation
	// db.DropTableIfExists("users", "profiles")
	// 2. create Profile1 ,User1
	// profileTest := Profile{Name: "heheh"}
	// db.Create(&profileTest)
	// user := User{Name: "Jinzhu", Age: 18, ProfileID: 1}
	// db.Create(&user)
	var profile Profile
	var userToFind User
	userToFind.ID = 1

	if err := db.First(&userToFind).Related(&profile, "Profile").Error; err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(userToFind)

	fmt.Printf("%+v\n", profile)

}
