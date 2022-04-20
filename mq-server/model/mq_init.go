package model

import "github.com/streadway/amqp"

/*
* @Author: hh
* @Date:   2022/4/15 9:12
 */

var MQ *amqp.Connection

func RabbitMQ(connString string){
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	MQ=conn
}
