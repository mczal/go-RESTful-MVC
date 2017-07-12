package controller

import (
	"go-restful-mvc/model"

	"go-restful-mvc/utility"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-siris/siris/context"
)

func PostAuth(ctx context.Context) {
	var userInput model.User
	ctx.ReadJSON(&userInput)

	if len(userInput.Email) == 0 || len(userInput.Password) == 0 {
		errWrongPass := model.BaseResponse{
			Status:  400,
			Message: "Bad Request | Email and Password must be exist as request body as JSON encoded.",
		}
		ctx.StatusCode(400)
		ctx.JSON(errWrongPass)
		return
	}

	// Do check password

	// See password hashing here and reCheck using bcrypt
	hashedPass, _ := utility.HashPassword(userInput.Password)

	if !utility.CheckPasswordHash(userInput.Password, hashedPass) {
		errWrongPass := model.BaseResponse{
			Status:  401,
			Message: "Unauthorized | Incorrect password",
		}
		ctx.StatusCode(401)
		ctx.JSON(errWrongPass)
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	// PleaseRead RFC 7519 for JWT Registered and Private Claim Name
	// link:  https://tools.ietf.org/html/rfc7519#section-4.1
	nowAddWeek := time.Now().AddDate(0, 0, 7).UnixNano() / 1000000
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "MY_ID",
		"role": "user",
		"exp":  nowAddWeek,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(utility.Configuration.Secret))
	if err != nil {
		errSignedString := model.BaseResponse{
			Status:  500,
			Message: "Internal Server Error | ErrTokenSignedString: " + err.Error(),
		}
		ctx.StatusCode(500)
		ctx.JSON(errSignedString)
		return
	}

	res := model.BaseSingleResponse{
		Status:  200,
		Message: "Success",
		Value:   tokenString,
	}
	ctx.StatusCode(200)
	ctx.JSON(res)
}
