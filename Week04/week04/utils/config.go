package utils

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"

)

var Conf Config

type Config struct {
	App      App      `yaml:"app,omitempty"`
	Databases Databases `yaml:"databases,omitempty"`
	Redis    Redis    `yaml:"redis,omitempty"`
	Emqtt    Emqtt    `yaml:"emqtt,omitempty"`
}

type App struct {
	Name string `yaml:"name,omitempty"`
	Port string `yaml:"port,omitempty"`
}

type Databases struct {
	Host    string `yaml:"host,omitempty"`
	Port    string `yaml:"port,omitempty"`
	Username     string `yaml:"username,omitempty"`
	Password  string `yaml:"password,omitempty"`
	Dbname  string `yaml:"dbname,omitempty"`
	MaxOpenConns      string `yaml:"maxOpenConns,omitempty"`
	MaxIdleConns string `yaml:"maxIdleConns,omitempty"`
	ConfName string `yaml:"confName,omitempty"`
}

type Emqtt struct {
	Host    string `yaml:"host,omitempty"`
	Port    string `yaml:"port,omitempty"`
	Username     string `yaml:"username,omitempty"`
	Password  string `yaml:"password,omitempty"`
	ClientID  string `yaml:"clientID,omitempty"`
}


type Redis struct {
	Host   string `yaml:"host,omitempty"`
	Port   string `yaml:"port,omitempty"`
	Explre string `yaml:"explre,omitempty"`
}

func InitConfig(ConfName string) (Config){
	if ConfName == ""{
		ConfName = "mysql"
	}
	//读取配置
	yamlFile, err := ioutil.ReadFile("conf/"+ConfName+"_conf.yaml")
	if err == nil {
		err = yaml.Unmarshal(yamlFile, &Conf)
	}else{
		Log(err)
	}
	//日志地址
	//file := "./storage/logs/" + Conf.App.Name + "_" + time.Now().Format("20060102") + ".log"
	//logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	//loger := log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	////SetFlags设置输出选项
	//loger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//
	//json.Marshal(err)

	return Conf
}
