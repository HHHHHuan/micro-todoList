package model

import "gorm.io/gorm"

/*
* @Author: hh
* @Date:   2022/4/15 16:07
 */

type Task struct {
	gorm.Model
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index; not null"`
	Status    int    `gorm:"default:0"`
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64 `gorm:"default:0"`
}
