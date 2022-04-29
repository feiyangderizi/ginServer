package config

type Mysql struct {
	Conn          string `mapstructure:"conn" json:"conn" yaml:"conn"`                               //主机地址:端口
	MaxIdleConns  int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` //最大空闲连接
	MaxOpenConns  int    `mapstructure:"max-open-conns" son:"max-open-conns" yaml:"max-open-conns"`  //数据库最大连接数
	SlowThreshold int    `mapstructure:"slow-threshold" son:"slow-threshold" yaml:"slow-threshold"`  //慢 SQL 阈值
	LogMode       string `mapstructure:"log-mode" son:"log-mode" yaml:"log-mode"`                    //日志模式（Silent、Error、Warn、Info）
}
