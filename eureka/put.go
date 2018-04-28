package eureka

import (
	"time"
	"bytes"
)

func (c *Client) SendHeartbeat(appName, instanceId string) error {
	var buf bytes.Buffer

	buf.WriteString("apps/")
	buf.WriteString(appName)
	buf.WriteString("/")
	buf.WriteString(instanceId)
	buf.WriteString("?")
	buf.WriteString("status=UP")
	path := buf.String()
	//
	go func() {
		for {
			time.Sleep(time.Duration(30 * time.Second))
			r, _ := c.Put(path, nil)
			if r.StatusCode == 404 {
				c.RegisterInstance(c.LocalInstanceInfo.App, &c.LocalInstanceInfo)
			}
		}
	}()

	return nil
}
