package model

/*
* @Author: hh
* @Date:   2022/4/15 16:06
 */

//执行数据迁移
func migration() {
	//自动迁移模式
	DB.Debug().AutoMigrate(&Task{})
}
