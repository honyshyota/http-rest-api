package store

import "github.com/honyshyota/http-rest-api/intenal/app/model"

// UserRepository ...
type UserRepository interface{
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}