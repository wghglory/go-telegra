package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"telegra/database"
	"telegra/model"
	"telegra/utils"

	"github.com/gofiber/fiber/v2"
)

// http://localhost:8080/createAccount?short_name=Derek&author_name=Derek%20Wang&author_url=https://github.com/wghglory/go-telegra
func CreateAccount(c *fiber.Ctx) error {
	accessToken := utils.GenerateSecureToken(32)

	account := &model.Account{
		ShortName:   c.Query("short_name"),
		AuthorName:  c.Query("author_name"),
		AuthorUrl:   c.Query("author_url"),
		AccessToken: accessToken,
		AuthUrl:     fmt.Sprintf("http://localhost:8080/auth/%v", accessToken),
	}

	data, err := database.EntClient.Account.
		Create().
		SetShortName(account.ShortName).
		SetAuthorName(account.AuthorName).
		SetAuthorURL(account.AuthorUrl).
		SetAccessToken(account.AccessToken).
		SetAuthURL(account.AuthUrl).
		Save(context.Background())

	if err != nil {
		return c.Status(500).JSON(&model.Response{
			Ok:    false,
			Error: fmt.Sprintf("failed creating account: %v", err),
		})
	}
	log.Println("Account was created: ", data)

	return c.JSON(&model.Response{
		Ok:     true,
		Result: data,
	})
}

// http://localhost:8080/getAccountInfo?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722&fields=[%22short_name%22,%22page_count%22]
func GetAccountInfo(c *fiber.Ctx) error {
	accessToken := c.Query("access_token")
	// fields (Array of String, default = [“short_name”,“author_name”,“author_url”])
	// List of account fields to return. Available fields: short_name, author_name, author_url, auth_url, page_count.
	fields := c.Query("fields")

	var props []string
	err := json.Unmarshal([]byte(fields), &props)
	result := map[string]any{}

	if err == nil {
		// TODO: from DB
		for _, p := range props {
			result[p] = "Derek"
		}
	}

	// TODO: middleware
	if accessToken == "" {
		return c.JSON(&model.Response{
			Ok:    false,
			Error: "Not a valid access token",
		})
	}

	return c.JSON(&model.Response{
		Ok:     true,
		Result: result,
	})
}

// http://localhost:8080/revokeAccessToken?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722
func RevokeAccessToken(c *fiber.Ctx) error {
	accessToken := c.Query("access_token")

	// TODO: middleware
	if accessToken == "" {
		return c.JSON(&model.Response{
			Ok:    false,
			Error: "Not a valid access token",
		})
	}

	newAccessToken := utils.GenerateSecureToken(32)

	account := &model.Account{
		AccessToken: newAccessToken,
		AuthUrl:     fmt.Sprintf("http://localhost:8080/auth/%v", newAccessToken)}

	return c.JSON(&model.Response{
		Ok:     true,
		Result: account,
	})
}
