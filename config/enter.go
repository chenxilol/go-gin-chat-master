package config

type Config struct {
	Mysql  Mysql  `json:"mysql" yaml:"mysql"`
	App    App    `json:"app" yaml:"app"`
	System System `json:"conSystem" yaml:"conSystem"`
}
