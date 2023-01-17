package services

type LoginService interface {
	LogInUser(email string, password string) bool
}

type LoginInfo struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &LoginInfo{
		email:    "lahirupr471@gmail.com",
		password: "lahiru12@",
	}
}
func (info *LoginInfo) LogInUser(email string, password string) bool {
	return info.email == email && info.password == password
}
