// Example of retrieving a secret from AWS Secrets Manager using the SDK
package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func main() {
	svc := secretsmanager.New(session.Must(session.NewSession()))

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("mySecretID"),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}

	fmt.Println("Secret value:", *result.SecretString)
}
