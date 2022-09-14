package config

type System struct {
	Env  string `json:"env" yaml:"env"`
	Port int    `json:"port" yaml:"port"`
}
