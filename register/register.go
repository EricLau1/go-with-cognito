package register

import (
	"go-with-cognito/client"
	"go-with-cognito/models"
	"go-with-cognito/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type Register interface {
	SignUp(*models.User) (*cognitoidentityprovider.SignUpOutput, error)
	Confirm(string, string) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
}

type registerImpl struct {
	cli *cognitoidentityprovider.CognitoIdentityProvider
}

func NewRegister(cli *cognitoidentityprovider.CognitoIdentityProvider) *registerImpl {
	return &registerImpl{cli}
}

func (r *registerImpl) SignUp(user *models.User) (*cognitoidentityprovider.SignUpOutput, error) {

	input := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(client.GetClientId()),
		Password: aws.String(user.Password),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			utils.AddAttr("email", user.Email),
			utils.AddAttr("name", user.Name),
			utils.AddAttr("nickname", user.Nickname),
			utils.AddAttr("phone_number", user.PhoneNumber),
		},
		Username: &user.Nickname,
	}

	return r.cli.SignUp(input)
}

func (r *registerImpl) Confirm(username, confirmationCode string) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {

	input := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(client.GetClientId()),
		ConfirmationCode: aws.String(confirmationCode),
		Username:         aws.String(username),
	}

	return r.cli.ConfirmSignUp(input)
}
