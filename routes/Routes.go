package routes

import (
	"go-restful-mvc/controller"

	"github.com/go-siris/siris/core/router"
)

// Route-Group called from Main.go.
// Applying route-group with middleware were very helpful for securing endpoints
// See Main.go for applying jwt middleware for this userRoutes

func UserParty(userRoutes router.Party) {
	userRoutes.Get("", controller.GetUserBydIDWithToken)
	userRoutes.Get("/users/{id:string}", controller.GetUserByID)
	userRoutes.Get("/users/email/{email:string}", controller.GetUserByEmail)
	userRoutes.Get("/users", controller.GetScanAllUser)
}
