package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Nerzal/gocloak/v8"
)

type authRepo struct{}

func NewAuthRepository() AuthRepository {
	return &authRepo{}
}

func (a *authRepo) Login(username, password string) (*gocloak.JWT, error) {
	client := gocloak.NewClient(os.Getenv("AUTH_SERVER_URL"))
	ctx := context.Background()
	// token, err := client.LoginAdmin(ctx, "admin", "admin", "master")
	token, err := client.Login(ctx, os.Getenv("CLIENT_NAME"), os.Getenv("CLIENT_CREDENTIAL"), os.Getenv("REALM"), "fearzsen3", "Mavericks.23")
	if err != nil {
		fmt.Print(err.Error())
		log.Fatal(err)
		fmt.Print("Something wrong with the credentials or url")
		fmt.Print(err.Error())
	}

	if token != nil {
		fmt.Print("success login")
	}

	return token, nil
}
