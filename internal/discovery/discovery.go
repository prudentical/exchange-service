package discovery

import (
	"exchange-service/internal/configuration"

	"fmt"
	"log/slog"
	"math/rand"
	"os"

	"github.com/hashicorp/consul/api"
)

type service struct {
	name string
	host string
	port int
}

type ServiceDiscovery interface {
	Register() error
	Discover(name string) (string, error)
}

type consulServiceDiscovery struct {
	logger  *slog.Logger
	service service
	client  *api.Client
}

func NewServiceDiscovery(logger *slog.Logger, config configuration.Config) ServiceDiscovery {
	consulConfig := api.DefaultConfig()
	address := fmt.Sprintf("%s:%d", config.Discovery.Server.Host, config.Discovery.Server.Port)
	consulConfig.Address = address

	client, err := api.NewClient(consulConfig)

	if err != nil {
		panic(err)
	}
	return consulServiceDiscovery{
		logger: logger,
		service: service{
			name: config.App.Name,
			host: config.Server.Host,
			port: config.Server.Port,
		},
		client: client,
	}
}

func (c consulServiceDiscovery) Register() error {
	c.logger.Debug("Registering to service discovery", "host", c.service.host, "port", c.service.port)
	host := c.service.host
	if host == "" {
		hostname, err := os.Hostname()
		if err != nil {
			return err
		}
		host = hostname
	}
	id := fmt.Sprintf("%s::%s:%d", c.service.name, host, c.service.port)

	service := &api.AgentServiceRegistration{
		ID:      id,
		Name:    c.service.name,
		Address: host,
		Port:    c.service.port,
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/health", c.service.host, c.service.port),
			Interval: "10s",
			Timeout:  "2s",
		},
	}
	return c.client.Agent().ServiceRegister(service)
}

func (c consulServiceDiscovery) Discover(name string) (string, error) {
	services, _, err := c.client.Health().Service(name, "", true, nil)
	if err != nil {
		return "", err
	}
	if len(services) == 0 {
		return "", NoInstanceAvailable{name}
	}
	instance := services[rand.Intn(len(services))]
	result := fmt.Sprintf("%s:%d", instance.Node.Address, instance.Service.Port)
	return result, nil
}
