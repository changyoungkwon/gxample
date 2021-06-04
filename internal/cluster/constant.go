package cluster

import "github.com/changyoungkwon/gxample/internal/config"

var (
	gatewayURL        = config.Get().Eureka.GatewayURL
	appID             = config.Get().Eureka.AppID
	instanceID        = config.Get().Eureka.InstanceID
	instanceHostname  = config.Get().Eureka.HostName
	instanceIPAddress = config.Get().Eureka.IPAdress
	instancePort      = config.Get().Eureka.Port
	eurekaTTL         = config.Get().Eureka.TTL
)
