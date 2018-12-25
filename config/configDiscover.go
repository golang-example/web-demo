package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"strings"
	"gopkg.in/afmobi-QSee/service-manager.v3"
)

var Cfg *ConfigMap
var configFile = "config/config.yml"
var configETCDPrefix string = "web-demo"

type ConfigMap struct {
	ListenPort          string		`json:"ListenPort" yaml:"listen_port"`

	LogPath				string		`json:"LogPath" yaml:"log_path"`

	RedisAddr			string		`json:"RedisAddr" yaml:"redis_addr"`
	RedisPassword		string		`json:"RedisPassword" yaml:"redis_password"`
	RedisMaxConn		int		    `json:"RedisMaxConn" yaml:"redis_max_conn"`

	MysqlUser			string		`json:"MysqlUser" yaml:"mysql_user"`
	MysqlPwd		    string		`json:"MysqlPwd" yaml:"mysql_pwd"`
	MysqlHost		    string		`json:"MysqlHost" yaml:"mysql_host"`
	MysqlPort		    int		    `json:"MysqlPort" yaml:"mysql_port"`
	MysqlDb		        string		`json:"MysqlDb" yaml:"mysql_db"`
	MysqlMaxConn		int		    `json:"MysqlMaxConn" yaml:"mysql_max_conn"`

	Liang		        string		`json:"Liang" yaml:"liang"`
}

func ConfigRead() {
	config := new(ConfigMap)
	config.readConfigFromYaml()
	config.readConfigFromEtcd()
	fmt.Println("Final config: ", config)

	if config.ListenPort == "" {
		panic("init config fail")
	}
	Cfg = config
}

func (this *ConfigMap)readConfigFromYaml() {
	if err := configor.Load(this, configFile); err != nil {
		fmt.Println("read config yaml err: " + err.Error())
	}
	fmt.Println("Config from yaml file: ", this)
}

func (this *ConfigMap)readConfigFromEtcd(){
	etcdArr := strings.Split(*EtcdAddr, ",")
	fmt.Println("Etcd Address: ", etcdArr)
	_, err := service_manager.InitConfig(configETCDPrefix, this, etcdArr)
	if err != nil {
		fmt.Println("Get Etcd Config Error: ", err.Error())
	}
	fmt.Println("Config from Etcd: ", this)
}
