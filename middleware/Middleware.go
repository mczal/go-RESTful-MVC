package middleware

import (
	"fmt"
	"go-restful-mvc/model"
	"go-restful-mvc/utility"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-siris/siris/context"
)

// ctx.Next is required to move forward to the chain of handlers,
// if missing then it stops the execution at this handler.

func LogMiddleware(ctx context.Context) {
	start := time.Now()
	ctx.Application().Log("%s\t%s\t%s \t| IP: %s", ctx.Method(), ctx.Path(), time.Since(start), ctx.RemoteAddr())
	ctx.Next()
}

func ValidateTokenUser(ctx context.Context) {
	auth := ctx.GetHeader("Authorization")
	if len(auth) <= 0 {
		err := model.BaseResponse{
			Status:  401,
			Message: "Unauthorized | Token is not present on header",
		}
		ctx.StatusCode(401)
		ctx.JSON(err)
		return
	}

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
		if claims["role"] == "user" || claims["role"] == "admin" {
			ctx.Next()
			return
		} else {
			errClaim := model.BaseResponse{
				Status:  401,
				Message: "Unauthorized | claim",
			}
			ctx.StatusCode(401)
			ctx.JSON(errClaim)
			return
		}

	} else {
		errElseClaim := model.BaseResponse{
			Status:  401,
			Message: "Unauthorized | errElseClaim",
		}
		ctx.StatusCode(401)
		ctx.JSON(errElseClaim)
		return
	}
}
