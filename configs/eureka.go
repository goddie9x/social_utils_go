package configs

import (
	"fmt"
	"post_service/pkg/dotenv"
	"strconv"

	"github.com/hudl/fargo"
)

type DiscoveryServerConnect struct {
	instance fargo.Instance
	conn     fargo.EurekaConnection
}

func (d *DiscoveryServerConnect) ConnectToEurekaDiscoveryServer() {
	portString := dotenv.GetEnvOrDefaultValue("API_PORT", "3005")
	port, _ := strconv.Atoi(portString)
	app_ip_address := dotenv.GetEnvOrDefaultValue("APP_IP_ADDRESS", "127.0.0.1")
	eureka_app_name := dotenv.GetEnvOrDefaultValue("EUREKA_APP_NAME", "post-service")
	host_name := dotenv.GetEnvOrDefaultValue("HOST_NAME", "localhost")
	d.instance = fargo.Instance{
		InstanceId:        fmt.Sprintf("%s:%s:%d", app_ip_address, eureka_app_name, port),
		HostName:          host_name,
		IPAddr:            app_ip_address,
		App:               eureka_app_name,
		VipAddress:        eureka_app_name,
		DataCenterInfo:    fargo.DataCenterInfo{Name: fargo.MyOwn, Class: "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo"},
		HealthCheckUrl:    fmt.Sprintf("http://%s:%d/health", host_name, port),
		StatusPageUrl:     fmt.Sprintf("http://%s:%d/status", host_name, port),
		HomePageUrl:       fmt.Sprintf("http://%s:%d/", host_name, port),
		LeaseInfo:         fargo.LeaseInfo{RenewalIntervalInSecs: 90, DurationInSecs: 120},
		Status:            fargo.UP,
		SecurePortEnabled: false,
		Port:              port,
		PortEnabled:       true,
	}
	var fargoConfig fargo.Config
	fargoConfig.Eureka.ServiceUrls = []string{
		dotenv.GetEnvOrDefaultValue("EUREKA_DISCOVERY_SERVER_URL", "http://localhost:8761/eureka"),
	}
	fargoConfig.Eureka.PollIntervalSeconds = 30
	d.conn = fargo.NewConnFromConfig(fargoConfig)
	d.conn.RegisterInstance(&d.instance)
}
