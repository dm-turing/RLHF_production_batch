package main

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/tracing"
)

func GetSecret(vaultURL, secretName string) (string, error) {
	authConfig := autorest.NewBearerAuthorizerFromEnvironment()
	client := keyvault.SecretsClient{BaseURI: vaultURL, Authorizer: authConfig}
	tracing.Start(context.Background(), "keyvault.SecretsClient.Get")
	defer tracing.End(context.Background(), "keyvault.SecretsClient.Get")
	resp, err := client.Get(context.Background(), secretName, "")
	if err != nil {
		return "", err
	}
	return *resp.Value, nil
}
