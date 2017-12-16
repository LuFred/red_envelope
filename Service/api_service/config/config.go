package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

// Config struct defines the config structure
type Config struct {
	ServiceName         string        `json:"service_name"`
	Debug               bool          `json:"debug"`
	TracingTransportURL string        `json:"tracing_transport_url"`
	Microservice        *Microservice `json:"microservice"`
	Listen              string        `json:"listen"`
	DefaultToken        string        `json:"default_token"`
}

//Microservice 微服务配置类
type Microservice struct {
	RedEnvelopeHost string `json:"red_envelope_service_host"`
}

var (
	ProConfig = Config{}
	env       = flag.String("env", "local", "运行环境")
	conf      = flag.String("conf", "nil", "配置文件路径")
)

//RegisterConfig 初始化config
func RegisterConfig() {
	var file []byte
	var err error
	if !flag.Parsed() {
		os.Stderr.Write([]byte("ERROR: config before flag.Parse"))
		os.Exit(1)
		return
	}
	if *env == "test" {
		file = []byte(test_config)
	} else {
		cp := ""
		if *conf == "nil" {
			cp, err = os.Getwd()
			if err != nil {
				log.Fatalln(err.Error())
			}
			cp = cp + "/config/local_conf.json"
		} else {
			cp = *conf
		}
		log.Println(cp)
		file, err = ioutil.ReadFile(cp)
	}
	if err != nil {
		os.Stderr.Write([]byte("ERROR: Read config file error"))
		os.Exit(1)
	}
	err = json.Unmarshal(file, &ProConfig)
	if err != nil {
		panic(err)
	}
}
