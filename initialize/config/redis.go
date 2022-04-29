package config

type Redis struct {
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`                               //服务器地址：端口
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   //密码
	DB           int    `mapstructure:"db" json:"db" yaml:"db"`                                     //库号
	MinIdleConns int    `mapstructure:"min-idle-conns" json:"min-idle-conns" yaml:"max-idle-conns"` //最小空闲连接
	PoolSize     int    `mapstructure:"pool-size" son:"pool-size" yaml:"pool-size"`                 //连接池最大socket连接数，默认为4倍CPU数
	IdleTimeOut  int    `mapstructure:"idle-time-out" son:"idle-time-out" yaml:"idle-time-out"`     //连接超时时间
}
