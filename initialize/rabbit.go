package initialize

import (
	"errors"

	"github.com/socifi/jazz"

	"github.com/feiyangderizi/ginServer/global"
)

type RabbitMQClient struct{}

func (rabbitClient *RabbitMQClient) init() {
	if global.Config.RabbitMQ.Addr == "" {
		panic(errors.New("RabbitMQ连接串配置"))
	}

	client, err := jazz.Connect(global.Config.RabbitMQ.Addr)
	if err != nil {
		global.Logger.Error("RabbitMQ连接错误:" + err.Error())
	} else {
		global.RABBITMQ = client
	}
}

func (rabbitClient *RabbitMQClient) close() {
	if global.RABBITMQ != nil {
		global.RABBITMQ.Close()
		global.RABBITMQ = nil
	}
}

func (rabbitClient *RabbitMQClient) Send(queueName string, msg string) error {
	err := global.RABBITMQ.SendMessage(global.Config.RabbitMQ.Exchange, queueName, msg)
	if err != nil {
		global.Logger.Error("RabbitMQ发送消息错误:" + err.Error())
		return err
	}
	return nil
}

func (rabbitClient *RabbitMQClient) Listener(queueName string, listener func(msg []byte)) {
	//侦听之前先创建队列
	rabbitClient.CreateQueue(queueName)
	//启动侦听消息处理线程
	go global.RABBITMQ.ProcessQueue(queueName, listener)
}

func (rabbitClient *RabbitMQClient) CreateQueue(queueName string) error {
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
	err := global.RABBITMQ.CreateScheme(*setting)
	if err != nil {
		global.Logger.Error("RabbitMQ创建队列失败:" + err.Error())
		return err
	}
	return nil
}
