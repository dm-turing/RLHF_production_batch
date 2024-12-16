// Example of retrieving a secret from Google Cloud Secret Manager using GCP SDK
package main

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	secretmanager "google.golang.org/api/secretmanager/v1"
)

func main() {
	ctx := context.Background()
	client, err := secretmanager.NewService(ctx, option.WithCredentialsFile("path/to/credentials.json"))
	if err != nil {
		fmt.Println("Failed to create secret manager client:", err)
		return
	}

	secretName := "projects/my-project/secrets/my-secret/versions/latest"
	resp, err := client.Projects.Secrets.Versions.Access(secretName).Do()
	if err != nil {
		fmt.Println("Failed to access secret:", err)
		return
	}

	fmt.Println("Secret value:", string(resp.Payload.Data))
}
