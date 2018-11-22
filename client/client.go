package client

func init(){
	appConfig := ParseConfig()
	RegisterEureka(appConfig)
}
