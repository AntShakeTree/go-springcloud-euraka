package eureka

import "github.com/AntShakeTree/go-springcloud-euraka/logger"


var INFO =logger.Info
var Error =logger.Error


//
func Eureka(server, hostname, app, instancesId, ip string, port int) {
	client := NewClient([]string{
		server, //From a spring boot based eureka server
	})
	instance := NewInstanceInfo(hostname, app, instancesId, ip, port, 30, false) //Create a new instance to register
	client.RegisterInstance(app, instance)          // Register new instance in your eureka(s)
	client.GetInstance(instance.App, instancesId)   // retrieve the instance from "test.com" inside "test"" app
	client.LocalInstanceInfo=*instance
	er:=client.SendHeartbeat(instance.App, instancesId) // say to eureka th
	if er!=nil{
		Error.Fatal(er)
	}
}

