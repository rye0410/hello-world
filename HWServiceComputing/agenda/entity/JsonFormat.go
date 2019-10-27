package entity

import (
	"encoding/json"
	"os"
	"fmt"
	"bufio"
	"io"
)

func User_JsonDecode(js []byte) User{
	var trans User
	err := json.Unmarshal(js, &trans)
	if err != nil {
		fmt.Println("error2")
	}
	return trans
}



func User_JsonEncode(m User) []byte {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error1")
		os.Exit(1)
	}
	return data
}

func User_ReadFromFile() []User{
	var userList []User
	file, err := os.Open("entity/data/User.txt")

    if err != nil {
        panic(err)
    }
    defer file.Close()
    read := bufio.NewReader(file)

    for {
        line, err := read.ReadString('\n') //ending with'\n'
        if err != nil || io.EOF == err {
            break
        }
        userList = append(userList, User_JsonDecode([]byte(line)))
    }
    return userList
}

func User_WriteToFile(My_User []User) {//**RegisterUser
	file, err := os.OpenFile("entity/data/User.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	os.Truncate("entity/data/User.txt", 0)

    if err != nil {
        fmt.Println("open file failed.", err.Error())
        os.Exit(1)
    }
    defer file.Close()

    for i := 0; i < len(My_User); i++ {
        file.WriteString(string(User_JsonEncode(My_User[i])[:]))
        file.WriteString("\n")
	}
}

func FromHost() []string{//**Init
	var tmp []string
	f, err := os.Open("entity/data/Host.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    rd := bufio.NewReader(f)
    for {
        line, err := rd.ReadString('\n') //ending with'\n'
        if err != nil || io.EOF == err {
            break
        }
        tmp = append(tmp, line)
    }
    return tmp
}

func ToHost(name string) {//**logIn
	file, err := os.OpenFile("entity/data/Host.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	os.Truncate("entity/data/Host.txt", 0)
    if err != nil {
        fmt.Println("open file failed.", err.Error())
        os.Exit(1)
    }
    defer file.Close()
        file.WriteString(name)
        file.WriteString("\n")
}

func Empty_login() {//**Empty_login
    os.Truncate("entity/data/Host.txt", 0)
}