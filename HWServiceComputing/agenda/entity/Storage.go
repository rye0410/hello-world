package entity

import (
	"fmt"  
    "regexp"
) 

var usersArray []User

func IsEmail(str string) bool {  
    var email bool  
    email, _ = regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", str)  
    if false == email {  
        return email  
    }
    return email  
}

func IsCellphone(str string) bool {  
    var phone bool  
    phone, _ = regexp.MatchString("^1[0-9]{10}$", str)  
    if false == phone {  
        return phone  
    }  
    return phone  
}  

func Init() {
	tmp_u := User_ReadFromFile()
	for i := 0; i < len(tmp_u); i++ {
		usersArray = append(usersArray, tmp_u[i])
	}
}


func RegisterUser(name string, password string, email string, phone string) bool {//**RegisterUser
	var user User
	err := false
	if (IsEmail(email) == false) {
		fmt.Println("Email is error!")
		err = true
	}
	if (IsCellphone(phone) == false) {
		fmt.Println("Phone is error!")
		err = true
	}
	if (len(password) < 6) {
		fmt.Println("The length of password can't be less than 6!")
		err = true
	}


	_, isExit, _:= QueryUser(name)

	if (isExit) {
		fmt.Println("This username exits, please use another username!")
		err = true
	}
	if (err) {
		return false
	}
	user.Name = name
	user.Password = password
	user.Email = email
	user.Phone = phone
	usersArray = append(usersArray,user)
	
	User_WriteToFile(usersArray)
	fmt.Println("Register successfully!")
	return true
}


func QueryUser(name string) (User,bool, int){//**QueryUser，logIn
	//Show()
	Init()
	for i := 0; i< len(usersArray); i++ {
		//Show()
		if usersArray[i].Name == name {
			return usersArray[i], true, i
		}
	}
	return User{"1","2","3","4"}, false, 0
}