package userservice

import (
	"fmt"
	"game/entity"
	"strconv"
)

type Service struct {
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
}

type RegisterResponse struct {
	User entity.User
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// 1. validate phone number
	if !isPhoneNumberValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}
	// 2. check uniqueness of phone number

	// 3. validate name

	// 4. create new user in storage

	// 5. return created user

}

func isPhoneNumberValid(phoneNumber string) bool {
	// TODO - we can use regular expression to support +98 pattern
	if len(phoneNumber) != 11 {
		return false
	}

	if phoneNumber[0:2] != "09" {
		return false
	}

	if _, err := strconv.Atoi(phoneNumber[2:]); err != nil {
		return false
	}

	return true
}
