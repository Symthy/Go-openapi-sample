package yaml

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type Twitter struct {
	ConsumerApiKey       string `json:"consumerApiKey" yaml:"consumerApiKey"`
	ConsumerApiSecretKey string `json:"consumerApiSecretKey" yaml:"consumerApiSecretKey"`
}

type Config struct {
	Twitter Twitter `json:"twitter" yaml:"twitter"`
}

func LoadConfigForYaml() (*Config, error) {
	f, err := os.Open("../conf/twitter-conf.yaml")
	if err != nil {
		//log.Fatal("load Yaml os.Open err:", err)
		log.Print("load Yaml os.Open err:", err)
		return nil, err
	}
	defer f.Close()

	var conf Config
	err = yaml.NewDecoder(f).Decode(&conf)
	return &conf, err
}
