package test

import (
	"github.com/AntShakeTree/go-springcloud-euraka/eureka"
	"github.com/AntShakeTree/go-springcloud-euraka/logger"
	"github.com/AntShakeTree/go-springcloud-euraka/tools"
	"testing"
)

var Err = logger.Error

func TestEureka(t *testing.T) {
	client := eureka.NewClient([]string{
		"http://192.168.103.111:7070/eureka", //From a spring boot based eureka server
	})

	instance := eureka.NewInstanceInfo(tools.Hostname(), "test", tools.Hostname()+":"+":8080", tools.IpAddr(), 8000, 30, false) //Create a new instance to register
	instanceId := tools.Hostname() + ":" + ":8080"
	client.RegisterInstance("test", instance)    // Register new instance in your eureka(s)
	client.GetInstance(instance.App, instanceId) // retrieve the instance from "test.com" inside "test"" app
	client.LocalInstanceInfo = *instance
	er := client.SendHeartbeat(instance.App, instanceId) // say to eureka th
	if er != nil {
		t.Fatal(er)
	}

}
