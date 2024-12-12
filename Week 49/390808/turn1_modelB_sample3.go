package main

import (
	"log"

	"github.com/hashicorp/consul/api"
)

func main() {
	config := &api.Config{Address: "consul:8500"}
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}

	svc := &api.AgentServiceRegistration{
		ID:      "my-service",
		Name:    "my-service",
		Address: "127.0.0.1",
		Port:    8080,
		Check: &api.AgentServiceCheck{
			TCP:      "127.0.0.1:8080",
			Interval: "10s",
		},
	}

	if err := client.Agent().ServiceRegister(svc); err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}

	defer client.Agent().ServiceDeregister(svc.ID)

	select {} // Wait forever
}
