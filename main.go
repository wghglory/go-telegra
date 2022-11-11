package main

import (
	"fmt"
	"log"

	"telegra/model"
	"telegra/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	// http://localhost:8080/createAccount?short_name=Derek&author_name=Derek%20Wang
	app.Get("/createAccount", func(c *fiber.Ctx) error {
		accessToken := utils.GenerateSecureToken(32)

		// TODO: update host
		account := &model.Account{
			ShortName:   c.Query("short_name"),
			AuthorName:  c.Query("author_name"),
			AuthorUrl:   c.Query("author_url"),
			AccessToken: accessToken,
			AuthUrl:     fmt.Sprintf("http://localhost:8080/auth/%v", accessToken)}

		return c.JSON(&model.Response{
			Ok:     true,
			Result: account,
		})
	})

	log.Fatal(app.Listen("localhost:8080"))
}
