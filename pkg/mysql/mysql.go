package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DataBaseInit() {
	var err error
	dsn := "root:@tcp(localhost:3306)/waysgallery?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("connect to database")
}
