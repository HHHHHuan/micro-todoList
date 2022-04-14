package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/*
* @Author: hh
* @Date:   2022/4/14 20:49
 */

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
}
