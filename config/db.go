package config

import (
	"FP-Sanbercode-Go-48-Tubagus_Saifulloh/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func ConnectDatabase(username, password, host, port, database string) *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	log.Println("Connecting", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&models.User{}, models.Post{}, models.Category{}, models.Comment{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
