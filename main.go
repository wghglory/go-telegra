package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

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

	// http://localhost:8080/createPage?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722&title=Sam%E5%8F%91%E9%A1%BA%E4%B8%B0ple+Page&author_name=Anonymous&content=[%7B%22tag%22:%22p%22,%22children%22:[%22Hello,+world!%22]%7D]&return_content=true
	app.Get("/createPage", func(c *fiber.Ctx) error {
		accessToken := c.Query("access_token")
		title := c.Query("title")
		path := fmt.Sprintf("%v-%v", title, time.Now())
		content := c.Query("content")

		return pageHandler(c, accessToken, path, title, content)
	})

	// http://localhost:8080/editPage/Sample-Page-12-15?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722&title=Sample+Page&author_name=Derek+Wang&content=[%7B%22tag%22:%22div%22,%22children%22:[%22Hello,+New+world!%22]%7D]&return_content=true
	// http://localhost:8080/editPage/Sample-Page-12-15?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722&title=Sample+Page&author_name=Derek+Wang&content=[%7B%22tag%22:%22div%22,%22children%22:[%22Hello,+New+world!%22]%7D]&return_content=false
	app.Get("/editPage/:path", func(c *fiber.Ctx) error {
		path := c.Params("path")
		accessToken := c.Query("access_token")
		title := c.Query("title")
		content := c.Query("content")

		return pageHandler(c, accessToken, path, title, content)
	})

	log.Fatal(app.Listen("localhost:8080"))
}

func pageHandler(c *fiber.Ctx, accessToken string, path string, title string, content string) error {
	if accessToken == "" {
		return c.JSON(&model.Response{
			Ok:    false,
			Error: "Not a valid access token",
		})
	}

	page := &model.Page{
		Path:        path,
		Title:       title,
		AuthorName:  c.Query("author_name"),
		Url:         fmt.Sprintf("http://localhost:8080/%v", path),
		Description: c.Query("description"),
		Content:     nil,
		Views:       1,
		CanEdit:     true,
	}

	var nodes []model.NodeElement
	err := json.Unmarshal([]byte(content), &nodes)

	if err != nil {
		return c.JSON(&model.Response{
			Ok:    false,
			Error: "Fail to provide a valid DOM element",
		})
	}

	if c.Query("return_content") == "true" {
		page.Content = nodes
	}

	return c.JSON(&model.Response{
		Ok:     true,
		Result: page,
	})
}
