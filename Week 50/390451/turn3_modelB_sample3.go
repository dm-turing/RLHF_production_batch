package main

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
)

func GetSecret(projectID, secretID string) (string, error) {
	ctx := context.Background()
	client, err := secretmanager.NewSecretVersionsClient(ctx, option.WithCredentialsFile("path/to/service-account.json"))
	if err != nil {
		return "", err
	}

	req := &secretmanager.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretID),
	}

	resp, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", err
	}

	return string(resp.Payload.Data), nil
}
