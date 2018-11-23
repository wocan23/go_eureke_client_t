package client

import (
	"math/rand"
	"sync"
	"github.com/ContainX/go-springcloud/discovery/eureka/model"
)

type ServiceSupport interface{
	GetSupportIndex(apps map[string]*model.Application,appName string) *model.Instance
}

/**
 随机提供
 */
type RandServiceSupport struct {

}

func (this RandServiceSupport)GetSupportIndex(apps map[string]*model.Application,appName string) *model.Instance{
	application := apps[appName]
	instances := application.Instances
	// 随机选取
	instanceIndex := rand.Intn(len(instances))
	instance := instances[instanceIndex]
	return instance
}

type CycleServiceSupport struct{
	lock sync.RWMutex
	appIndexMap map[string]int
}

func (this CycleServiceSupport)GetSupportIndex(apps map[string]*model.Application,appName string) *model.Instance{
	defer this.lock.Unlock()
	this.lock.Lock()
	application := apps[appName]
	instances := application.Instances
	totalInstanceNum := len(instances)

	lastIndex := this.appIndexMap[appName]
	var currentIndex = lastIndex + 1
	if lastIndex > totalInstanceNum{
		currentIndex = 0
	}
	this.appIndexMap[appName] = currentIndex
	return instances[currentIndex]
}

func getServiceSupport() ServiceSupport{
	switch appConfig.ServiceSupport {
	case randServiceSupport:
		return RandServiceSupport{}
	case cycleServiceSupport:
		return CycleServiceSupport{}
	default:
		return RandServiceSupport{}
	}
}

