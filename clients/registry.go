package clients

import (
	"order-service/clients/config"
	userClient "order-service/clients/user"
	configApp "order-service/config"
)

type ClientRegistry struct{}

type IClientRegistry interface {
	GetUser() userClient.IUserClient
}

func NewClientRegistry() IClientRegistry {
	return &ClientRegistry{}
}

func (c *ClientRegistry) GetUser() userClient.IUserClient {
	return userClient.NewUserClient(
		config.NewClientConfig(
			config.WithBaseURL(configApp.Config.InternalService.User.Host),
			config.WithSignatureKey(configApp.Config.InternalService.User.SignatureKey),
		))
}
