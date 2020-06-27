package config

import (
	"blog/common"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	ProductionMode = "product"
	LogFile = "/data/www/blog/app.log"
	Port = 9999
	Dialect = "mysql"
)

var (
	Mode string
	PassSalt string
	TokenMaxAge int
	TokenSecret string
)

type dbConfig struct {
	Host 	  string
	Port      int
	User      string
	Password  string
}

type redisConfig struct {
	Host      string
	Port      int
	Password  string
	URL       string
	MaxIdle   int
	MaxActive int
}

var (
	DBConfig dbConfig        // mysql 配置
	RedisConfig redisConfig  // redis相关配置
)

var log = common.Log

func configDB(conf map[string]interface{}) {
	// 解析出来的conf["redis"]是一个map[interface{}]interface{}类型
	config, _ :=  (conf["mysql"]).(map[interface{}]interface{})
	DBConfig.Host, _ = config["host"].(string)
	DBConfig.Port, _ = config["port"].(int)
	DBConfig.User, _ = config["user"].(string)
	DBConfig.Password, _ = config["password"].(string)
}

func configRedis(conf map[string]interface{}) {
	// 解析出来的conf["redis"]是一个map[interface{}]interface{}类型
	config, _ :=  (conf["redis"]).(map[interface{}]interface{})
	RedisConfig.Host, _ = config["host"].(string)
	RedisConfig.Port, _ = config["port"].(int)
	RedisConfig.MaxIdle, _ = config["maxIdle"].(int)
	RedisConfig.MaxActive, _ = config["maxActive"].(int)
	RedisConfig.Password, _ = config["password"].(string)
}

func configEnv(conf map[string]interface{}) {
	config, _ :=  (conf["env"]).(map[interface{}]interface{})
	Mode, _ = config["mode"].(string)
	TokenSecret, _ = config["TokenSecret"].(string)
	TokenMaxAge, _ = config["TokenMaxAge"].(int)
	PassSalt, _ = config["PassSalt"].(string)
}

func parseConfig(configFile string) {
	buffer, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	conf := map[string]interface{}{}
	err = yaml.Unmarshal(buffer, &conf)
	if err != nil {
		log.Fatalf(err.Error())
	}
	configRedis(conf)
	configDB(conf)
	configEnv(conf)
	fmt.Println("配置．．．．", DBConfig, RedisConfig)
	if include, ok := conf["include"]; ok {
		_includes, _ok := include.([]interface{})
		if _ok == false {
			panic("include必须是数组")
		}
		for _, file := range _includes {
			fmt.Println("file=", file)
			filepath := file.(string)
			parseConfig(filepath)
		}
	}
}

func init() {
	fmt.Println("config init...")
	parseConfig("config/deploy.yaml")
}