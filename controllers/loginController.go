package controllers

type login struct {
	Username string
	Password string
}

var loginInstance *login

func Init(username, password string) *login {
	if loginInstance == nil {
		loginInstance = &login{
			Username: username,
			Password: password,
		}
	}
	return loginInstance
}

func CheckLogin(username, password string) string {
	if username == "admin" && password == "P@ssw0rd" {
		return "success"
	}
	return "fail"
}
