package database

import (
	"fmt"
	"waysgallery/models"
	"waysgallery/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Transaction{})

	if err != nil {
		fmt.Println(err)
		panic("migration error")
	}
	fmt.Println("migration success")
}
