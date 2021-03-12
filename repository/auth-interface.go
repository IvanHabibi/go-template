package repository

import "github.com/Nerzal/gocloak/v8"

type AuthRepository interface {
	Login(username, password string) (*gocloak.JWT, error)
}
