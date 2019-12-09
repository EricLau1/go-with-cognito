package admin

import (
	"fmt"
	"go-with-cognito/client"
	"go-with-cognito/models"
	"go-with-cognito/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

const (
	flow_admin_initiate_auth = "ADMIN_NO_SRP_AUTH"
)

type Admin interface {
	CreateUser(*models.User) (*cognitoidentityprovider.AdminCreateUserOutput, error)
	InitAuth(string, string) (*cognitoidentityprovider.AdminInitiateAuthOutput, error)
	RespondChallenge(*cognitoidentityprovider.AdminInitiateAuthOutput, *models.User) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error)
	UpdateUser(*models.User) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error)
	DeleteUser(string) (*cognitoidentityprovider.AdminDeleteUserOutput, error)
	GetUser(string) (*cognitoidentityprovider.AdminGetUserOutput, error)
	GetUsers() (*cognitoidentityprovider.ListUsersOutput, error)
	FilterUser(*cognitoidentityprovider.AttributeType) (*cognitoidentityprovider.ListUsersOutput, error)
}

type adminImpl struct {
	cli *cognitoidentityprovider.CognitoIdentityProvider
}

func NewAdmin(cli *cognitoidentityprovider.CognitoIdentityProvider) *adminImpl {
	return &adminImpl{cli}
}

func (adm *adminImpl) CreateUser(user *models.User) (*cognitoidentityprovider.AdminCreateUserOutput, error) {

	input := &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:             aws.String(client.GetUserPoolId()),
		DesiredDeliveryMediums: []*string{aws.String("EMAIL")},
		Username:               &user.Nickname,
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			utils.AddAttr("email", user.Email),
			utils.AddAttr("name", user.Name),
			utils.AddAttr("nickname", user.Nickname),
			utils.AddAttr("phone_number", user.PhoneNumber),
		},
	}

	return adm.cli.AdminCreateUser(input)
}

func (adm *adminImpl) InitAuth(username, pass string) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {

	params := map[string]*string{
		"USERNAME": &username,
		"PASSWORD": &pass,
	}

	input := &cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow:       aws.String(flow_admin_initiate_auth),
		AuthParameters: params,
		ClientId:       aws.String(client.GetClientId()),
		UserPoolId:     aws.String(client.GetUserPoolId()),
	}

	return adm.cli.AdminInitiateAuth(input)
}

func (adm *adminImpl) RespondChallenge(auth *cognitoidentityprovider.AdminInitiateAuthOutput, user *models.User) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error) {

	challengeResponses := map[string]*string{
		"NEW_PASSWORD": aws.String(user.Password),
		"USERNAME":     aws.String(user.Nickname),
	}

	input := &cognitoidentityprovider.AdminRespondToAuthChallengeInput{
		ChallengeName:      auth.ChallengeName,
		ChallengeResponses: challengeResponses,
		ClientId:           aws.String(client.GetClientId()),
		UserPoolId:         aws.String(client.GetUserPoolId()),
		Session:            auth.Session,
	}

	return adm.cli.AdminRespondToAuthChallenge(input)
}

func (adm *adminImpl) UpdateUser(user *models.User) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error) {
	input := &cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			utils.AddAttr("name", user.Name),
			utils.AddAttr("phone_number", user.PhoneNumber),
		},
		UserPoolId: aws.String(client.GetUserPoolId()),
		Username:   aws.String(user.Nickname),
	}

	return adm.cli.AdminUpdateUserAttributes(input)
}

func (adm *adminImpl) DeleteUser(username string) (*cognitoidentityprovider.AdminDeleteUserOutput, error) {

	input := &cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: aws.String(client.GetUserPoolId()),
		Username:   aws.String(username),
	}

	return adm.cli.AdminDeleteUser(input)
}

func (adm *adminImpl) GetUser(username string) (*cognitoidentityprovider.AdminGetUserOutput, error) {

	input := &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: aws.String(client.GetUserPoolId()),
		Username:   aws.String(username),
	}

	return adm.cli.AdminGetUser(input)
}

func (adm *adminImpl) GetUsers() (*cognitoidentityprovider.ListUsersOutput, error) {

	input := &cognitoidentityprovider.ListUsersInput{
		Limit:      aws.Int64(50),
		UserPoolId: aws.String(client.GetUserPoolId()),
	}

	return adm.cli.ListUsers(input)
}

func (adm *adminImpl) FilterUser(attr *cognitoidentityprovider.AttributeType) (*cognitoidentityprovider.ListUsersOutput, error) {
	filter := fmt.Sprintf("%s=\"%s\"", *attr.Name, *attr.Value)

	input := &cognitoidentityprovider.ListUsersInput{
		Filter:     &filter,
		UserPoolId: aws.String(client.GetUserPoolId()),
	}

	return adm.cli.ListUsers(input)
}
