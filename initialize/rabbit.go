package initialize

import (
	"errors"

	"github.com/socifi/jazz"

	"github.com/feiyangderizi/ginServer/global"
)

var rabbit *jazz.Connection

type RabbitMQClient struct{}

func (client *RabbitMQClient) init() {
	if global.Config.RabbitMQ.Addr == "" {
		panic(errors.New("RabbitMQ连接串配置"))
	}

	if rabbit == nil {
		client, err := jazz.Connect(global.Config.RabbitMQ.Addr)
		if err != nil {
			global.Logger.Error("RabbitMQ连接错误:" + err.Error())
		} else {
			rabbit = client
		}
	}
}

func (client *RabbitMQClient) close() {
	if rabbit != nil {
		rabbit.Close()
		rabbit = nil
	}
}

func (client *RabbitMQClient) Send(queueName string, msg string) {
	err := rabbit.SendMessage(global.Config.RabbitMQ.Exchange, queueName, msg)
	if err != nil {
		global.Logger.Error("RabbitMQ发送消息错误:" + err.Error())
	}
}

func (client *RabbitMQClient) Listener(queueName string, listener func(msg []byte)) {
	//侦听之前先创建队列
	client.CreateQueue(queueName)
	//启动侦听消息处理线程
	go rabbit.ProcessQueue(queueName, listener)
}

func (client *RabbitMQClient) CreateQueue(queueName string) {
	queues := make(map[string]jazz.QueueSpec)
	binding := &jazz.Binding{
		Exchange: global.Config.RabbitMQ.Exchange,
		Key:      queueName,
	}
	queueSpec := &jazz.QueueSpec{
		Durable:  true,
		Bindings: []jazz.Binding{*binding},
	}
	queues[queueName] = *queueSpec
	setting := &jazz.Settings{
		Queues: queues,
	}
	err := rabbit.CreateScheme(*setting)
	if err != nil {
		global.Logger.Error("RabbitMQ创建队列失败:" + err.Error())
	}
}
