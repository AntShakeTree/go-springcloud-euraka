package eureka

import (
	"encoding/json"
	"strings"

	
)

func (c *Client) RegisterInstance(appId string, instanceInfo *InstanceInfo) error {
	values := []string{"apps", appId}
	path := strings.Join(values, "/")
	instance := &Instance{
		Instance: instanceInfo,
		//Eureka:EurekaInstanceConfigBean{PreferIpAddress:true,IpAddress:instanceInfo.IpAddr},
	}
	body, err := json.Marshal(instance)
	if err != nil {
		return err
	}
	//common.Corn("1/30 * * * * *", func() {
	_, err = c.Post(path, body)
	//})
	//fmt.Println(string(res.Body))
	return err
}
