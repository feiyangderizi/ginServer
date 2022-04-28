package config

type Application struct {
	Name          string `mapstructure:"name" json:"name" yaml:"name"`                                  //应用名
	Port          string `mapstructure:"port" json:"addr" yaml:"port"`                                  // 端口值
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                         // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                   // 使用redis
	UseMongodb    bool   `mapstructure:"use-mongodb" json:"use-mongodb" yaml:"use-mongodb"`             // 使用mongodb
	UseRabbitMQ   bool   `mapstructure:"use-rabbitmq" json:"use-rabbitmq" yaml:"use-rabbitmq"`          // 使用rabbitmq
	AutoCheckTime int    `mapstructure:"auto-check-time" json:"auto-check-time" yaml:"auto-check-time"` // 自动检查时间
}
