package service

import (
	"go-restful-mvc/model"
)

func NewUser(user *model.User) (int, interface{}) {
	// Your db call service implementation instead of resSuccess:=

	resSuccess := model.BaseResponse{
		Status:  200,
		Message: "Success create new user, email: " + user.Email,
	}
	return 200, resSuccess
}

func FindUserByEmail(email string) (int, interface{}) {
	// Your db call service implementation instead of user:=

	user := model.UserDetailDtoResponse{
		Email:    email,
		FullName: "mczal",
		Role:     "user",
	}
	result := model.BaseSingleResponse{
		Status:  200,
		Message: "Success",
		Value:   user,
	}
	return 200, result
}

func ScanUser() (int, interface{}) {
	// Your db call service implementation instead of users:=

	users := []model.UserSimpleDtoResponse{
		{
			Email:    "mczal1@example.com",
			FullName: "mczal1",
			Role:     "user",
		},
		{
			Email:    "mczal2@example.com",
			FullName: "mczal",
			Role:     "user",
		},
		{
			Email:    "mczal3@example.com",
			FullName: "mczal3",
			Role:     "user",
		},
	}

	userInts := make([]interface{}, len(users))
	for i, v := range users {
		userInts[i] = interface{}(v)
	}

	result := model.BaseListResponse{
		Status:  200,
		Message: "Success",
		Content: userInts,
	}
	return 200, result
}

func FindUserByID(userID string) (int, interface{}) {
	// Your db call service implementation instead of user:=

	user := model.UserDetailDtoResponse{
		Email:       "mczal@example.com",
		FullName:    "mczal fullname",
		Address:     "Mczal Street 313",
		Role:        "user",
		PhoneNumber: "1111111",
	}

	succRes := model.BaseSingleResponse{
		Status:  200,
		Message: "Success get user by id: " + userID,
		Value:   user,
	}

	return 200, succRes
}
