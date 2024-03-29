package client

import (
	"github.com/ContainX/go-springcloud/discovery/eureka"
	"github.com/ContainX/go-springcloud/discovery/eureka/model"
	"time"
	"fmt"
)



func registerEureka(appConfig *AppConfig) eureka.EurekaClient{
	eurekaConfig := model.NewConfigFromArgs(*appConfig.AppName,*appConfig.Host,*appConfig.Port,appConfig.EurekaUrls...)
	eurekaClient := eureka.NewClient(eurekaConfig)
	eurekaClient = eurekaClient
	err := eurekaClient.Register(true)
	if err !=  nil{
		panic("注册eureka出现错误"+err.Error())
	}
	go flushAppInfos()
	return eurekaClient
}

func flushAppInfos(){
	for t := range time.Tick(30*time.Second){
		fmt.Println(t.Format("2006-01-02 15:04:05")+"刷新应用实例")
		apps,_ := eurekaClient.GetApplications()
		apps = apps
	}
}
