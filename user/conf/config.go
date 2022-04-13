package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"user/model"
)

var (
	Db         			string
	DbHost     			string
	DbPort     			string
	DbUserName     		string
	DbPassWord 			string
	DbName     			string
)

func Init()  {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("load ini file failed..")
		panic(err)
	}
	LoadMysqlData(file)
	dsn:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUserName, DbPassWord, DbHost, DbPort, DbName)
	model.DataBase(dsn)
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUserName = file.Section("mysql").Key("DbUserName").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
