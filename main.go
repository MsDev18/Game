package main

import (
	"fmt"
	"game/entity"
	"game/repository/mysql"
	"net/http"
)

func main() {

	http.HandleFunc("/users/register" , userRegisterHandler)
	err := http.ListenAndServe(":8081", nil)
	fmt.Println("err: ", err)
}



// ServeHTTP(ResponseWriter, *Request)
func userRegisterHandler (writer http.ResponseWriter , req *http.Request) {
	fmt.Println("Here..")
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer , "invalid method")
	}
}

func tetsUserMysqlRepo() {
	mysqlRepo := mysql.New()

	createdUser, err := mysqlRepo.Register(entity.User{
		ID:          0,
		Name:        "MsDev18",
		PhoneNumber: "09351721410",
	})

	if err != nil {
		fmt.Printf("register user: %v \n", err)
	} else {
		fmt.Printf("created user: %v\n", createdUser)
	}
	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber + "5")

	if err != nil {
		fmt.Printf("unique err: %v", err)
	}
	fmt.Println("isUnique -> ", isUnique)
}
