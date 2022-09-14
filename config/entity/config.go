package entity

var CONFIG Config

type Config struct {
	App   App   `mapstructrue:"app" json:"app" yaml:"app"`
	Mysql Mysql `mapstructrue:"mysql" json:"mysql" yaml:"mysql"`
	Redis Redis `mapstructrue:"redis" json:"redis" yaml:"redis"`
	Zap   Zap   `mapstructrue:"zap" json:"zap" yaml:"zap"`
}
