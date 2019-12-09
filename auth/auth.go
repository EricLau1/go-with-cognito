package auth

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"go-with-cognito/client"
	"go-with-cognito/models"
)

type Auth interface {
	Login(*models.User) (*cognitoidentityprovider.InitiateAuthOutput, error) 
	ForgotPassword(string) (*cognitoidentityprovider.ForgotPasswordOutput, error)
	ConfirmForgotPassword(string, string, string) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error) 
}

type authImpl struct {
	cli *cognitoidentityprovider.CognitoIdentityProvider
}

func NewAuth(cli *cognitoidentityprovider.CognitoIdentityProvider) *authImpl {
	return &authImpl{cli}
}

func (a *authImpl) Login(user *models.User) (*cognitoidentityprovider.InitiateAuthOutput, error) {

	params := map[string]*string{
		"PASSWORD": &user.Password,
		"USERNAME": &user.Nickname,
	}

	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: params,
		ClientId:       aws.String(client.GetClientId()),
	}

	return a.cli.InitiateAuth(input)
}

func (a *authImpl) ForgotPassword(username string) (*cognitoidentityprovider.ForgotPasswordOutput, error) {

	input := &cognitoidentityprovider.ForgotPasswordInput{
		Username: aws.String(username),
		ClientId: aws.String(client.GetClientId()),
	}

	return a.cli.ForgotPassword(input)
}

// username, newPassword, confirmationCode
func (a *authImpl) ConfirmForgotPassword(username, newPassword, confirmationCode string) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error) {

	input := &cognitoidentityprovider.ConfirmForgotPasswordInput{
		Username: aws.String(username),
		ClientId: aws.String(client.GetClientId()),
		ConfirmationCode: aws.String(confirmationCode),
		Password: aws.String(newPassword),
	}

	return a.cli.ConfirmForgotPassword(input)
}