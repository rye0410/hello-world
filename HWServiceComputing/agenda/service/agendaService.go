package service

import (
	"fmt"  
    "os"
    "strings"
    "log"
    "github.com/spf13/agenda/entity"
)

var my_name, my_password string
var Login_flag bool 
var All_name []string

var log_file *os.File



func GetFlag() bool {
	return Login_flag
}

func BeginWithLog() {
	entity.Init()

    logFile,err := os.OpenFile("service/agenda.log",os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
    log_file = logFile
    if err != nil {
        log.Fatalln("Fail to open file!")
    }
	tmp := entity.FromHost()
	if (len(tmp)==0) {
		Login_flag = false
	} else {
		Login_flag = true
		my_name = strings.Replace(tmp[0],"\n","",-1)
	}
	
}

func RegisterUser(name string, password string, email string, phone string) {
	BeginWithLog()
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	i := entity.RegisterUser(name, password, email, phone)
	if (i) {
		debugLog.Println(name, " register successfully!")
	} else {
		debugLog.Println(name, " register failed!")
	}
	defer log_file.Close()
}

func LogIn(name string, password string) {
	BeginWithLog()
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	tmp_u, flag, _:= entity.QueryUser(name)
	if flag == true {
		my_name = name
		my_password = password
		if (entity.GetPassword(tmp_u) != password) {
			debugLog.Println(name, " log in failed!")
			fmt.Println("Password is wrong!")
		} else {
			debugLog.Println(name, " log in successfully!")
			fmt.Println("Log in successfully!\nWelcome to Agenda!")
		}
	} else {
		debugLog.Println(name, " Log in failed!")
		fmt.Println("Haven't register!")
	}
	entity.ToHost(name)
	defer log_file.Close()
}

func LogOut() {
	BeginWithLog()
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	debugLog.Println(my_name, " Log out successfully!")
	fmt.Println("Log out successfully!")
	entity.Empty_login()
	defer log_file.Close()
}

func queryUser(name string) {
	BeginWithLog()
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	tmp_u, flag, _ := entity.QueryUser(name)
	if !flag {
		debugLog.Println(my_name, " query user ", name, " failed!")
		fmt.Println(name," doesn't exists!")
	} else {
		debugLog.Println(my_name, " query user ", name, " successfully!")
		fmt.Println("Name : ", entity.GetName(tmp_u))
		fmt.Println("Email : ", entity.GetEmail(tmp_u))
		fmt.Println("Phone : ", entity.GetPhone(tmp_u))
	}
	defer log_file.Close()
}
