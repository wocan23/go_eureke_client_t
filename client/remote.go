package client

import (
	"fmt"
	"net/http"
	"strings"
	"encoding/json"
	"io/ioutil"
	"github.com/kataras/iris/core/errors"
	"time"
)


/**
	调用其他微服务方法
 */
func ExecRemoteFunc(appName string,urlPath string,paramObj interface{},resultObjPtr interface{}) error{
	url := getRemoteClientUrl(appName,urlPath)
	resBytes,err := PostWithTimeout(url,paramObj,defaultTimeout)
	if err != nil{
		return err
	}
	json.Unmarshal(resBytes,resultObjPtr)
	return nil
}

/**
	随机选取客户端方法
 */
func getRemoteClientUrl( appName string,urlPath string) string{
	instance := realServiceSupport.GetSupportIndex(apps,appName)
	ip := instance.IpAddr
	port := instance.Port.Number
	url := fmt.Sprintf(remoteUrl,ip,port,urlPath)
	return url
}

func Post(url string,paramObj interface{})([]byte,error){
	// 超时处理
	var res *http.Response
	var err error
	paramByte, _ := json.Marshal(paramObj)
	res,err = http.Post(url,"application/json",strings.NewReader(string(paramByte)))
	if err != nil{
		return nil,err
	}
	if res.StatusCode != 200{
		return nil,errors.New("返回状态异常")
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func PostWithTimeout(url string,paramObj interface{},timeout time.Duration)([]byte,error){
	var resBytes []byte
	var err error
	var do = make(chan int)

	timeoutChan := time.After(timeout)
	go func() {
		resBytes,err = Post(url,paramObj)
		do<-1
	}()
	for {
		select{
		case <-timeoutChan:
			return nil,errors.New("调用超时")
		case <-do:
			return resBytes,err
		default:
		}
	}
}

