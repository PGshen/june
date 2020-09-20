package setting

import (
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	Config *Conf
)

type Conf struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Redis    Redis    `yaml:"redis"`
	App      App      `yaml:"app"`
}

type Server struct {
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read-timeout"`
	WriteTimeout time.Duration `yaml:"write-timeout"`
}

type App struct {
	RunMode     string `yaml:"run-mode"`
	PageSize    int    `yaml:"page-size"`
	IdentityKey string `yaml:"identity-key"`
	LogPath     string `yaml:"log-path"`
	AesKey      string `yaml:"aes-key"`
}

type Database struct {
	Type        string `yaml:"type"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Name        string `yaml:"name"`
	TablePrefix string `yaml:"table-prefix"`
}

type Redis struct {
	Addr        string        `yaml:"addr"`
	Pass        string        `yaml:"pass"`
	DB          int           `yaml:"db"`
	Timeout     time.Duration `yaml:"timeout"`
	ExpiredTime int           `yaml:"expired-time"`
}

func init() {
	Config = getConf()
	log.Println("[Setting] Config init success")
}

func getConf() *Conf {
	var c *Conf
	file, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		log.Println("[Setting] config error: ", err)
	}
	err = yaml.UnmarshalStrict(file, &c)
	if err != nil {
		log.Println("[Setting] yaml unmarshal error: ", err)
	}
	return c
}
