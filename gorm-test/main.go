package main
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
	"fmt"
	"time"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func (Product) TableName() string {
	return "product"
}
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

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=test sslmode=disable password=tiger")
	if err != nil {
		log.Fatalf("failed to connect database,error:%#v",err)
	}
	if err=db.DB().Ping();err!=nil{
		log.Fatalf("database is unusable,err:%v\n",err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	defer db.Close()
	//db.LogMode(true)
	// Migrate the schema
	db=db.Table("test.user")
	db.AutoMigrate(&User{})

	// insert
	db.Create(&User{Name: "kity", Balance: 1000})
	db.Create(&User{Name: "lili", Balance: 100})
	db.Create(&User{Name: "jete", Balance: 500})
	db.Create(&User{Name: "france", Balance: 2000})

	// Read
	var user User
	/*db.First(&user, 1) // find product with id 1
	db.Model(&user).Update("balance", 2000)
	fmt.Printf("id为1的user：%v\n",&user)*/
	db.Debug().Model(&User{}).Select("id,created_at,balance").First(&user, "balance > ?", 100) // find product with code l1212
	//db.Model(&user).Update("balance", 50)
	fmt.Printf("第一个balance>100的user：%#v\n",&user)
	// Update - update product's price to 2000
	db.Where("balance < ?", 100).Delete(User{})
	var names []string
	db.Model(&User{}).Pluck("name", &names)
	fmt.Printf("names：%#v\n",&names)

	// Delete - delete product
	//db.Delete(&product)
}