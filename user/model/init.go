package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func DataBase(dsn string)  {
	//dsn:="root:root@tcp(127.0.0.1:3306)/go_bubble_demo?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy:schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	DB=db
	migration()
}
