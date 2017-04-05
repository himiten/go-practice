package gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/gommon/log"
	"time"
	//"fmt"
	//"path/filepath"
)
type User struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name string
	Balance float64
}
func (User) TableName() string{
	return "user"
}
var (
	db *gorm.DB
	err error
)
func init(){
	db, err = gorm.Open("sqlite3", "../test.db")
	if err != nil {
		log.Fatalf("failed to connect database,error:%#v",err)
	}
	//defer db.Close()
}

func GetUserById(id uint)  *User{
	var user User
	db.Debug().First(&user,id)
	return &user
}