package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Addr          string `yaml:"addr"`
	SMSFile       string `yaml:"sms_file"`
	MMSAddr       string `yaml:"mms_addr"`
	VoiceCallFile string `yaml:"voice_call_file"`
	EmailFile     string `yaml:"email_file"`
	BillingFile   string `yaml:"billing_file"`
	SupportAddr   string `yaml:"support_addr"`
	IncidentAddr  string `yaml:"incident_addr"`
}

var GlobalConfig Config

func NewConfig(file string) Config {
	fmt.Println("Get config.yaml")
	var config Config

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return GetDefaultConfig()
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err.Error())
		return GetDefaultConfig()
	}

	return config
}

func GetDefaultConfig() Config {
	fmt.Println("Get default config")
	const dir = "/"

	const addr = "localhost:8080"

	var config Config

	config.Addr = addr
	config.SMSFile = dir + "sms.data"
	config.MMSAddr = "http://" + addr + "/mms"
	config.VoiceCallFile = dir + "voice.data"
	config.EmailFile = dir + "email.data"
	config.BillingFile = dir + "billing.data"
	config.SupportAddr = "http://" + addr + "/support"
	config.IncidentAddr = "http://" + addr + "/accendent"

	return config
}
