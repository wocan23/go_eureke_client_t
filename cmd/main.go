package main

import (
	"fmt"
	"github.com/ContainX/go-springcloud/discovery/eureka/model"
	"github.com/ContainX/go-springcloud/discovery/eureka"
	"net/http"
	"time"
	"../constant"
	"github.com/itgeniusshuai/go_common/common"
	"git.youxinpai.com/golang/uxtools/yxpconfig"
	"os/exec"
	"os"
	"reflect"
	"runtime"
	"strings"
	"strconv"
)


func main(){

}


var eurekaClient eureka.EurekaClient

var goConfig *yxpconfig.GoConfigServer


func mainFunc(){
	path,err := exec.LookPath(os.Args[0])
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(path)

	var appName = "go-client"
	ip,_ := common.GetLocalIp()
	var host = ip
	var port = 9878
	var serviceUrls = []string{"http://10.70.93.52:9876"}
	eurekaConfig := model.NewConfigFromArgs(appName,host,port,serviceUrls...)
	eurekaClient = eureka.NewClient(eurekaConfig)
	apps,err := eurekaClient.GetApplications()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(apps)
	eurekaClient.Register(true)

	mux := http.NewServeMux()
	mux.HandleFunc("/test/go",test)
	mux.HandleFunc("/test/goService",testGoService)
	http.ListenAndServe(":9878",mux)

	time.Sleep(10000*time.Second)
}

func test(w http.ResponseWriter, r *http.Request){
}

var pids = make(map[string]string)

func testGoService(w http.ResponseWriter, r *http.Request){
	fmt.Println("receivced java request")
	pid := common.IntToStr(common.GetGoroutineId())

	fmt.Println("pid:"+pid)
	pids[pid] = "a"
	goService(pid)
	time.Sleep(500*time.Millisecond)
	w.Write([]byte("go service return "))
}

func goService(pid string){
	fmt.Println(pid+":pid:"+common.IntToStr(common.GetGoroutineId()))
	fmt.Println("pids len:"+common.IntToStr(len(pids)))
}






