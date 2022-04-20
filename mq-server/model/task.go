package model

/*
* @Author: hh
* @Date:   2022/4/18 19:32
 */

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index; not null"`
	Status    int    `gorm:"default:0"`
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64 `gorm:"default:0"`
}
