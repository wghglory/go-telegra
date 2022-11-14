package main

import (
	"fmt"

	"telegra/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	// api := app.Group("/api/v1")
	api := app

	// http://localhost:8080/createAccount?short_name=Derek&author_name=Derek%20Wang&author_url=https://github.com/wghglory/go-telegra
	api.Get("/createAccount", handlers.CreateAccountGetHandler)
	api.Post("/createAccount", handlers.CreateAccountPostHandler)

	// http://localhost:8080/getAccountInfo?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722&fields=[%22short_name%22,%22page_count%22]
	api.Get("/getAccountInfo", handlers.GetAccountInfoGetHandler)
	api.Post("/getAccountInfo", handlers.GetAccountInfoPostHandler)

	// http://localhost:8080/revokeAccessToken?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722
	api.Get("/revokeAccessToken", handlers.RevokeAccessTokenGetHandler)
	api.Post("/revokeAccessToken", handlers.RevokeAccessTokenPostHandler)

	// http://localhost:8080/createPage?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722&title=Sam%E5%8F%91%E9%A1%BA%E4%B8%B0ple+Page&author_name=Anonymous&content=[%7B%22tag%22:%22p%22,%22children%22:[%22Hello,+world!%22]%7D]&return_content=true
	api.Get("/createPage", handlers.CreatePageGetHandler)
	api.Post("/createPage", handlers.CreatePagePostHandler)

	// http://localhost:8080/getPage/Sample-Page-12-15?return_content=true
	api.Get("/getPage/:path", handlers.GetPage)

	// http://localhost:8080/editPage/Sample-Page-12-15?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722&title=Sample+Page&author_name=Derek+Wang&content=[%7B%22tag%22:%22div%22,%22children%22:[%22Hello,+New+world!%22]%7D]&return_content=true
	// http://localhost:8080/editPage/Sample-Page-12-15?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722&title=Sample+Page&author_name=Derek+Wang&content=[%7B%22tag%22:%22div%22,%22children%22:[%22Hello,+New+world!%22]%7D]&return_content=false
	api.Get("/editPage/:path", handlers.EditPage)
}

func main() {
	app := fiber.New()

	Routes(app)

	err := app.Listen("localhost:8080")
	if err != nil {
		fmt.Println("Unable to start server")
	}

	fmt.Println("Server is up on port 8080")
}
