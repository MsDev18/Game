package main

import (
	"fmt"
	"game/entity"
	"game/repository/mysql"
)

func main() {
	mysqlRepo := mysql.New()

	createdUser, err := mysqlRepo.Register(entity.User{
		ID:          0,
		Name:        "MsDev18",
		PhoneNumber: "09351721415",
	})
	fmt.Println(createdUser)
	if err != nil  {
		fmt.Printf("register user: %v \n", err)
	} else {
		fmt.Printf("created user: %v\n" , createdUser)
	}

	fmt.Println("t",createdUser.PhoneNumber)
	isUnique , err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber)
	fmt.Println(isUnique, err)
	if err != nil {
		fmt.Printf("unique err: %v", err)
	}
	fmt.Println("isUnique -> ", isUnique)
}
