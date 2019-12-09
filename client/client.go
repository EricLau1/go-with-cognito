package client

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func NewClient() *cognitoidentityprovider.CognitoIdentityProvider {

	sess, err := createSessionWithStaticCredentials()
	if err != nil {
		log.Fatal(err)
	}

	return cognitoidentityprovider.New(sess)
}

func createSessionWithStaticCredentials() (*session.Session, error) {

	creds := credentials.NewStaticCredentials(getAccessKeyId(), getSecretKeyId(), "")

	return session.NewSession(&aws.Config{
		Region:      aws.String(GetRegion()),
		Credentials: creds,
	})
}
