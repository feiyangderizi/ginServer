package config

type Mysql struct {
	Conn         string `mapstructure:"conn" json:"conn" yaml:"conn"`                               //主机地址:端口
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` //最大空闲连接
	MinOpenConns int    `mapstructure:"min-open-conns" son:"min-open-conns" yaml:"min-open-conns"`  //数据库最小连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" son:"max-open-conns" yaml:"max-open-conns"`  //数据库最大连接数
	IdleTimeOut  int    `mapstructure:"idle-time-out" son:"idle-time-out" yaml:"idle-time-out"`     //连接超时时间
	LogMode      bool   `mapstructure:"log-mode" son:"log-mode" yaml:"log-mode"`                    //是否打开日志
}
