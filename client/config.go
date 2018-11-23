package client

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"net"
	"fmt"
	"errors"
)



func parseConfig() *AppConfig{
	defaultAppConfig := new(AppConfig)
	err := parseYml(configPath+".yaml",defaultAppConfig)
	if err != nil{
		panic("配置文件conf/eureka.yaml没有找到")
	}
	if defaultAppConfig.Profile != nil{
		profileAppConfig := new(AppConfig)
		err := parseYml(configPath+*defaultAppConfig.Profile+".yaml",defaultAppConfig)
		if err !=  nil{
			defaultAppConfig.Port = profileAppConfig.Port
			defaultAppConfig.AppName = profileAppConfig.AppName
			defaultAppConfig.EurekaUrls = profileAppConfig.EurekaUrls
		}
	}
	// ip
	ip,_ := getLocalIp()
	defaultAppConfig.Host = &ip
	return defaultAppConfig
}




func parseYml(filePath string,configObj interface{}) error{
	configByte,err := ioutil.ReadFile(filePath)
	if err != nil{
		return err
	}
	return  yaml.Unmarshal(configByte,configObj)
}

func getLocalIp() (string,error){
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return "",err
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(),nil
			}
		}
	}
	return "",errors.New("can't get ip")
}
