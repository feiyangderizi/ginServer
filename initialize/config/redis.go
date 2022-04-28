package config

type Redis struct {
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`                               //服务器地址：端口
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   //密码
	DB           int    `mapstructure:"db" json:"db" yaml:"db"`                                     //库号
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` //最大空闲连接
	MinOpenConns int    `mapstructure:"min-open-conns" son:"min-open-conns" yaml:"min-open-conns"`  //数据库最小连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" son:"max-open-conns" yaml:"max-open-conns"`  //数据库最大连接数
	IdleTimeOut  int    `mapstructure:"idle-time-out" son:"idle-time-out" yaml:"idle-time-out"`     //连接超时时间
}
