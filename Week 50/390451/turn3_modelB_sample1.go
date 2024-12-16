package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func GetSecret(secretName string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		return "", err
	}

	svc := secretsmanager.New(sess)
	params := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	resp, err := svc.GetSecretValue(params)
	if err != nil {
		return "", err
	}

	secret := *resp.SecretString
	if err != nil {
		return "", err
	}

	return secret, nil
}

func main() {
	fmt.Println(GetSecret("secretName"))
}
