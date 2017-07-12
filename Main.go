package main

import (
	"go-restful-mvc/controller"
	"go-restful-mvc/middleware"
	"go-restful-mvc/routes"
	"go-restful-mvc/service"

	"go-restful-mvc/utility"

	"github.com/go-siris/siris"
	"github.com/go-siris/siris/context"
)

func main() {
	utility.InitializeConfig()
	// fmt.Printf("Config: %v", utility.Configuration) uncomment this for see externalizing configuration works.
	service.GenerateDynamoDBSvc()

	app := siris.New()

	// Regster custom handler for specific http errors.
	app.OnErrorCode(siris.StatusInternalServerError, func(ctx context.Context) {
		// .Values are used to communicate between handlers, middleware.
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}

		ctx.Writef("(Unexpected) internal server error")
	})

	// Applying Log Middleware To All Request
	app.Use(middleware.LogMiddleware)

	// API open for public
	app.Post("/register", controller.PostNewUser)
	app.Post("/auth", controller.PostAuth)

	routes.UserParty(app.Party("/user", middleware.ValidateTokenUser))
	// routes.AdminParty(app.Party("/admin", middleware.ValidateTokenAdmin))

	app.StaticWeb("/static", "./public")

	// Listen for incoming HTTP/1.x & HTTP/2 clients on localhost port defined in Configuration from env.json.
	app.Run(siris.Addr(":"+utility.Configuration.Port), siris.WithCharset("UTF-8"))
}
