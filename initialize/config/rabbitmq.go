package config

type RabbitMQ struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             //服务器地址：端口
	Exchange string `mapstructure:"exchange" json:"exchange" yaml:"exchange"` //所属交换机
}
