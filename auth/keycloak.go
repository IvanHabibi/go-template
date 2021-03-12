package auth

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v8"
	"github.com/dgrijalva/jwt-go"
)

func Authorization() {

	client := gocloak.NewClient("http://localhost:8080/")
	ctx := context.Background()
	// token, err := client.LoginAdmin(ctx, "admin", "admin", "master")
	token, err := client.Login(ctx, "service-oms", "07b23e14-0ab7-422d-b88b-0d5d6ba447ab", "ged", "fearzen3", "Mavericks.23")
	if err != nil {
		fmt.Print(err.Error())
		// log.Fatal(err)
		fmt.Print("Something wrong with the credentials or url")
		// fmt.Print(err.Error())
	}

	if token != nil {
		fmt.Println("success login")
		fmt.Println(token.AccessToken)
	}

	// user := gocloak.User{
	// 	FirstName: gocloak.StringP("Bob"),
	// 	LastName:  gocloak.StringP("Uncle"),
	// 	Email:     gocloak.StringP("something@really.wrong"),
	// 	Enabled:   gocloak.BoolP(true),
	// 	Username:  gocloak.StringP("CoolGuy"),
	// }
	// fmt.Printf(token.AccessToken)
	// client := gocloak.NewClient("http://localhost:8080/")
	//  (*jwt.Token, *jwt.MapClaims, error)
	// jwtToken, _, err := client.DecodeAccessToken(ctx, token.AccessToken, "ged", "")
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	// log.Fatal(err)
	// 	fmt.Print("Something wrong with the credentials or url")
	// 	// fmt.Print(err.Error())
	// }
	// if !strings.Contains(jwtToken.Signature, "Barang") {
	// fmt.Println(jwtToken.Raw)
	// }
	claims := jwt.MapClaims{}
	token2, err := jwt.ParseWithClaims(token.AccessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("07b23e14-0ab7-422d-b88b-0d5d6ba447ab"), nil
	})
	// ... error handling

	// do something with decoded claims
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}

	fmt.Println(token2)
	// jwt.ParseClaimStrings(mapClaims)
	// fmt.Println("Header :", jwtToken.Header)
	// fmt.Println("Method :", jwtToken.Method)
	// fmt.Println("Signature :", jwtToken.Signature)
	// fmt.Println("Valid :", jwtToken.Valid)
	// fmt.Println(jwtToken.)
	// fmt.Println(jwtToken.Signature)
	// _, err = client.CreateUser(ctx, token.AccessToken, "realm", user)
	// if err != nil {
	// 	panic("Oh no!, failed to create user :(")
	// }
}

func InstrospectToken() {
	client := gocloak.NewClient("http://localhost:8080/")
	ctx := context.Background()
	token, err := client.LoginClient(ctx, "oms", "6323500e-4e20-4fc5-b1f0-d34e4c3f103a", "ged")
	if err != nil {
		panic("Login failed:" + err.Error())
	}

	rptResult, err := client.RetrospectToken(ctx, token.AccessToken, "oms", "6323500e-4e20-4fc5-b1f0-d34e4c3f103a", "ged")
	if err != nil {
		panic("Inspection failed:" + err.Error())
	}

	fmt.Print(rptResult.Active)
	// if !rptResult. {
	// 	panic("Token is not active")
	// }

	permissions := rptResult.Permissions
	fmt.Print(permissions)
	// Do something with the permissions ;)

}
