package config

type ServerConfig struct {
	Application Application `mapstructure:"application" json:"application" yaml:"application"`
	Redis       Redis       `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql       Mysql       `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	MongoDB     MongoDB     `mapstructure:"mongoDB" json:"mongoDB" yaml:"mongoDB"`
	RabbitMQ    RabbitMQ    `mapstructure:"rabbitmq" json:"rabbitmq" yaml:"rabbitmq"`
}
