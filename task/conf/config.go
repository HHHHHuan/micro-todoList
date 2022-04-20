package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
	"task/model"
)

var (
	Db         			string
	DbHost     			string
	DbPort     			string
	DbUserName     		string
	DbPassWord 			string
	DbName     			string

	RabbitMQ string
	RabbitMQUser string
	RabbitMQPassWord string
	RabbitMQHost string
	RabbitMQPort string
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

	//连接rabbitmq
	LoadRabbitMQ(file)
	pathRabbitMQ:=strings.Join([]string{RabbitMQ,"://",RabbitMQUser,":",RabbitMQPassWord,"@", RabbitMQHost, ":", RabbitMQPort, "/"},"")
	model.RabbitMQ(pathRabbitMQ)
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUserName = file.Section("mysql").Key("DbUserName").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadRabbitMQ(file *ini.File) {
	RabbitMQ = file.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassWord = file.Section("rabbitmq").Key("RabbitMQPassWord").String()
	RabbitMQHost = file.Section("rabbitmq").Key("RabbitMQHost").String()
	RabbitMQPort = file.Section("rabbitmq").Key("RabbitMQPort").String()
}
