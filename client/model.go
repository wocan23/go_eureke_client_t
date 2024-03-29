package client

type ServiceSupportType int

type AppConfig struct{
	Profile *string `yaml:"profile"`
	AppName *string `yaml:"appName"`
	Host *string `yaml:"host"`
	Port *int `yaml:"port"`

	EurekaUrls []string `yaml:"eurekaUrls"`
	ServiceSupport ServiceSupportType `yaml:"serviceSupport"`
}
