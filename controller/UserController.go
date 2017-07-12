package controller

import (
	"fmt"
	"go-restful-mvc/model"
	"go-restful-mvc/service"
	"go-restful-mvc/utility"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-siris/siris/context"
)

func GetUserBydIDWithToken(ctx context.Context) {
	auth := ctx.GetHeader("Authorization") // Or convert directly using: .Values().GetInt/GetInt64 etc...

	token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(utility.Configuration.Secret), nil
	})
	if err != nil {
		errParse := model.BaseResponse{
			Status:  401,
			Message: "Unauthorized | parse: " + err.Error(),
		}
		ctx.StatusCode(401)
		ctx.JSON(errParse)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["sub"].(string)
		statQuery, resQuery := service.FindUserByID(userID)
		ctx.StatusCode(statQuery)
		ctx.JSON(resQuery)
	} else {
		errElseClaim := model.BaseResponse{
			Status:  401,
			Message: "Unauthorized | errElseClaim",
		}
		ctx.StatusCode(401)
		ctx.JSON(errElseClaim)
	}
}

func GetScanAllUser(ctx context.Context) {
	statScan, resScan := service.ScanUser()
	ctx.StatusCode(statScan)
	ctx.JSON(resScan)
}

func GetUserByID(ctx context.Context) {
	userID := ctx.Params().Get("id") // Or convert directly using: .Values().GetInt/GetInt64 etc...
	statQuery, resQuery := service.FindUserByID(userID)
	ctx.StatusCode(statQuery)
	ctx.JSON(resQuery)
}

func GetUserByEmail(ctx context.Context) {
	userEmail := ctx.Params().Get("email")
	if !strings.Contains(userEmail, "@") || !strings.Contains(userEmail, ".") {
		badReq := model.BaseResponse{
			Status:  400,
			Message: "Bad Request | Email validation format constraint",
		}
		ctx.StatusCode(400)
		ctx.JSON(badReq)
		return
	}
	statQuery, resQuery := service.FindUserByEmail(userEmail)
	ctx.StatusCode(statQuery)
	ctx.JSON(resQuery)
}
