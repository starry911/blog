package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Conf struct {
	Mysql     MySql
	Redis     Redis
	Server    Server
	Logs      Logs
	Coroutine Coroutine
	Jwt       Jwt
}
type Jwt struct {
	Secret string `yaml:"Secret"`
	TTL    uint   `yaml:"TTL"`
}

type Coroutine struct {
	PollNum int `yaml:"PollNum"`
}

type MySql struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Database string `yaml:"Database"`
	Pool     string `yaml:"Pool"`
	Charset  string `yaml:"Charset"`
}

type Redis struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Password string `yaml:"Password"`
	Database int    `yaml:"Database"`
	Pool     int    `yaml:"Pool"`
	Conn     int    `yaml:"Conn"`
}

type Server struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
	Mode string `yaml:"Mode"`
}

type Logs struct {
	Level     string `yaml:"Level"`
	Type      string `yaml:"Type"`
	FileName  string `yaml:"FileName"`
	MaxSize   int    `yaml:"MaxSize"`
	MaxBackup int    `yaml:"MaxBackup"`
	MaxAge    int    `yaml:"MaxAge"`
	Compress  bool   `yaml:"Compress"`
}

func GetConf() Conf {
	var conf Conf
	// 加载文件
	yamlFile, err := os.ReadFile("etc/config.yaml")
	if err != nil {
		panic(err)
	}
	// 将读取的yaml文件解析为响应的 struct
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
