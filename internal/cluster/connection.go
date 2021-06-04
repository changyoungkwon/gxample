package cluster

import (
	"sync"
	"time"

	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/changyoungkwon/gxample/internal/logging"
)

// Connection declares gateway and instance information
type Connection struct {
	gateway  *eureka.Client
	instance *eureka.InstanceInfo
}

var (
	once sync.Once
	conn *Connection
)

// GetConn register the instance to gateway(ensures registered state)
func GetConn() *Connection {
	once.Do(func() {
		conn = &Connection{
			gateway: eureka.NewClient([]string{gatewayURL}),
			instance: eureka.NewInstanceInfo(
				instanceHostname, appID, instanceIPAddress, instancePort, eurekaTTL, false,
			),
		}
		err := conn.gateway.RegisterInstance(conn.instance.App, conn.instance)
		if err != nil {
			logging.Logger.Errorf("error on register instance")
			panic(err)
		}
	})
	return conn
}

// SendHeartbeatForever send heartbests forever
func (c *Connection) SendHeartbeatForever() error {
	_, err := c.gateway.GetApplication(c.instance.App)
	if err != nil {
		logging.Logger.Errorf("error before send heartbeat, %v", err)
		return err
	}
	go func() {
		for {
			<-time.After(time.Second * 30)
			err := c.gateway.SendHeartbeat(c.instance.App, c.instance.HostName)
			if err != nil {
				logging.Logger.Debugf("error while sending heatbeat, %v", err)
			}
		}
	}()
	return nil
}
