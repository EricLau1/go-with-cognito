package client

import "os"

func getAccessKeyId() string {
	return os.Getenv("AWS_ACCESS_KEY_ID")
}

func getSecretKeyId() string {
	return os.Getenv("AWS_SECRET_KEY_ID")
}

func GetUserPoolId() string {
	return os.Getenv("AWS_USER_POOL_ID")
}

func GetRegion() string {
	return os.Getenv("AWS_USER_POOL_REGION")
}

func GetClientId() string {
	return os.Getenv("AWS_CLIENT_ID")
}
