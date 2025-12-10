package mocks

import "github.com/DuduAlmeida/dudu.web.go.jwt-poc/domain/user"

var UsersMocked = user.UserList{
	{ID: "user123", Username: "testeuser", Password: "testepassword"},
	{ID: "user321", Username: "testeuser321", Password: "testepassword"},
}
