package middleware

import (
	"Walter0697/GinBackend/utility"
	"Walter0697/GinBackend/service"
)

type LoginService interface {
	LoginUser(username string, password string) bool
}

type loginInformation struct {
	username  string
	password  string
}

//for testing purpose
// func StaticLoginService() LoginService {
// 	return &loginInformation{
// 		username:    "testing",
// 		password: 	"testing",
// 	}
// }

func FetchingLoginService(username string) LoginService {
	user, err := service.FindUserByName(username)
	if err != nil {
		return &loginInformation{}
	}
	return &loginInformation{
		username: user.Username,
		password: user.Password,
	}
}

func (info *loginInformation) LoginUser(username string, password string) bool {
	return info.username == username && utility.ComparePassword(info.password, password)
}