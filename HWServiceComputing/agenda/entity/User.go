package entity

type User struct {
	Name string
	Password string
	Email string
	Phone string
}
func GetName(a User)string {//queryUser
	return a.Name
}

func GetPhone(a User) string{//queryUser
	return a.Phone
}
func GetEmail(a User) string{//queryUser
	return a.Email
}
func GetPassword(a User) string{//**logIn
	return a.Password
}