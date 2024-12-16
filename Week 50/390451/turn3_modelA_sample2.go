// Example of retrieving a secret from Azure Key Vault using Azure SDK
package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/secrets"
)

func main() {
	vaultName := "<Your-Key-Vault-Name>"
	secretName := "<Your-Secret-Name>"

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Println("Failed to create credential: ", err)
		return
	}

	client, err := secrets.NewClient("https://"+vaultName+".vault.azure.net/", cred, nil)
	if err != nil {
		fmt.Println("Failed to create secret client: ", err)
		return
	}

	secretResp, err := client.GetSecret(context.TODO(), secretName, nil)
	if err != nil {
		fmt.Println("Failed to get secret: ", err)
		return
	}

	fmt.Println("Secret value:", *secretResp.Value)
}
