package discovery

import (
	"errors"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
)

// Consul服务注册与发现客户端
type ConsulClient struct {
	client *api.Client
}

// 初始化Consul客户端
func NewConsulClient(address string) (*ConsulClient, error) {
	config := api.DefaultConfig()
	config.Address = address

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ConsulClient{client: client}, nil
}

// 注册服务
func (c *ConsulClient) RegisterService(serviceID, name, address string, port int, tags []string) error {
	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    name,
		Address: address,
		Port:    port,
		Tags:    tags,
		Check: &api.AgentServiceCheck{
			HTTP:                           "http://" + address + ":" + strconv.Itoa(port) + "/health",
			Interval:                       "10s",
			Timeout:                        "5s",
			DeregisterCriticalServiceAfter: "30s",
		},
	}

	return c.client.Agent().ServiceRegister(registration)
}

// 发现服务
func (c *ConsulClient) DiscoverService(serviceName string) ([]*api.ServiceEntry, error) {
	queryOptions := &api.QueryOptions{
		WaitTime: 15 * time.Second, // 长轮询等待时间
	}

	services, _, err := c.client.Health().Service(serviceName, "", true, queryOptions)
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, errors.New("未发现服务实例")
	}

	return services, nil
}

// 注销服务
func (c *ConsulClient) DeregisterService(serviceID string) error {
	return c.client.Agent().ServiceDeregister(serviceID)
}
