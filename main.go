package main

import (
	"errors"
	"fmt"
	"go-with-cognito/admin"
	"go-with-cognito/auth"
	"go-with-cognito/client"
	"go-with-cognito/models"
	"go-with-cognito/register"
	"go-with-cognito/utils"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	c := client.NewClient()

	//adm := admin.NewAdmin(c)

	//AdminCreateUser(adm)
	//AdminInitAuth(adm)
	//AdminUpdateUser(adm)
	// AdminGetUser(adm)
	//AdminGetUsers(adm)
	//AdminFilterUsers(adm)
	//AdminDeleteUser(adm)

	//r := register.NewRegister(c)

	//Register(r)

	a := auth.NewAuth(c)

	Login(a)
	//ForgotPassword(a)
}

func AdminCreateUser(adm admin.Admin) {

	user := &models.User{
		Email:       "eric12364597-test@email.com",
		Name:        "John Doe",
		Nickname:    "johndoe",
		PhoneNumber: "+55123456789",
	}

	out, err := adm.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}

func AdminInitAuth(adm admin.Admin) {

	user := &models.User{
		Email:       "eric.devtt@gmail.com",
		Name:        "Eric Lau",
		Nickname:    "ericlau1",
		PhoneNumber: "+551111111111",
		Password:    "Qwer@1234",
	}

	tmpPass := "F9ofi8f;"

	out, err := adm.InitAuth(user.Nickname, tmpPass)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(out)
	out2, err := adm.RespondChallenge(out, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out2)
}

func AdminUpdateUser(adm admin.Admin) {

	user := &models.User{
		Email:       "eric.devtt@gmail.com",
		Name:        "John Doe",
		Nickname:    "ericlau1",
		PhoneNumber: "+55222255511",
	}

	out, err := adm.UpdateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}

func AdminGetUser(adm admin.Admin) {

	user := &models.User{
		Email:       "eric.devtt@gmail.com",
		Name:        "John Doe",
		Nickname:    "ericlau1",
		PhoneNumber: "+55222255511",
	}

	out, err := adm.GetUser(user.Nickname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}

func AdminGetUsers(adm admin.Admin) {

	out, err := adm.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}

func AdminFilterUsers(adm admin.Admin) {

	out, err := adm.FilterUser(utils.AddAttr("sub", "4d6e82b9-b1d1-44e8-9432-842a79f6acab"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}

func AdminDeleteUser(adm admin.Admin) {

	out, err := adm.DeleteUser("johndoe")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}

func Register(r register.Register) {

	user := &models.User{
		Email:       "eric.devtt@gmail.com",
		Name:        "John Doe",
		Nickname:    "ericlau1",
		PhoneNumber: "+55222255511",
		Password:    "Qwert1234@",
	}

	out, err := r.SignUp(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)

	var code string
	fmt.Print("Confirmation Code? ")
	fmt.Scan(&code)

	out2, err := r.Confirm(user.Nickname, code)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out2)
}

func Login(a auth.Auth) {

	user := &models.User{
		Email:       "eric.devtt@gmail.com",
		Name:        "John Doe",
		Nickname:    "ericlau1",
		PhoneNumber: "+55222255511",
		Password:    "1234Aabc@",
	}

	out, err := a.Login(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}

func ForgotPassword(a auth.Auth) {

	out, err := a.ForgotPassword("ericlau1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)

	var code string
	var newPass string
	var confirmPass string

	fmt.Print("New Password? ")
	fmt.Scan(&newPass)
	fmt.Print("Confirm New Password? ")
	fmt.Scan(&confirmPass)

	if newPass != confirmPass {
		log.Fatal(errors.New("Different passwords"))
	}

	fmt.Print("Confirmation Code? ")
	fmt.Scan(&code)

	out2, err := a.ConfirmForgotPassword("ericlau1", newPass, code)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out2)
}
