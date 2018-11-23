package client

import (
	"github.com/ContainX/go-springcloud/discovery/eureka"
	"github.com/ContainX/go-springcloud/discovery/eureka/model"
	"time"
)
func init(){
	appConfig = parseConfig()
	realServiceSupport = getServiceSupport()
	registerEureka(appConfig)
}

// 配置文件路径
const configPath  = "../conf/eureka"
// 远端服务调用
const remoteUrl = "http://%s:%d/%s"
const defaultTimeout = 5*time.Second
var appConfig *AppConfig
var eurekaClient eureka.EurekaClient
var apps map[string]*model.Application

var realServiceSupport ServiceSupport
const(
	randServiceSupport  ServiceSupportType = iota
	cycleServiceSupport
)



