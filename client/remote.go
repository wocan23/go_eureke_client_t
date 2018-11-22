package client

import (
	"math/rand"
	"fmt"
	"net/http"
	"strings"
	"encoding/json"
	"io/ioutil"
)

const URL = "http://%s:%d/%s"
/**
	调用其他微服务方法
 */
func ExecRemoteFunc(appName string,urlPath string,paramObj interface{},resultObjPtr interface{}) error{
	url := getRemoteClientUrl(appName,urlPath)
	paramByte, _ := json.Marshal(paramObj)
	res,err := http.Post(url,"application/json",strings.NewReader(string(paramByte)))
	if err != nil{
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body,resultObjPtr)
	return nil
}

/**
	随机选取客户端方法
 */
func getRemoteClientUrl( appName string,urlPath string) string{
	application := Apps[appName]
	instances := application.Instances
	// 随机选取
	instanceIndex := rand.Intn(len(instances))
	instance := instances[instanceIndex]
	ip := instance.IpAddr
	port := instance.Port.Number
	url := fmt.Sprintf(URL,ip,port,urlPath)
	return url
}