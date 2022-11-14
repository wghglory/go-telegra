package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"telegra/database"
	"telegra/ent/account"
	"telegra/model"
	"telegra/utils"

	"github.com/gofiber/fiber/v2"
)

// http://localhost:8080/createAccount?short_name=Derek&author_name=Derek%20Wang&author_url=https://github.com/wghglory/go-telegra
func CreateAccountGetHandler(c *fiber.Ctx) error {
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
		Result: account,
	})
}

func CreateAccountPostHandler(c *fiber.Ctx) error {
	accessToken := utils.GenerateSecureToken(32)

	account := new(model.Account)

	if err := c.BodyParser(account); err != nil {
		return err
	}

	account.AccessToken = accessToken
	account.AuthUrl = fmt.Sprintf("http://localhost:8080/auth/%v", accessToken)

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
		Result: account,
	})
}

// right access token with default fields [“short_name”,“author_name”,“author_url”]: http://localhost:8080/getAccountInfo?access_token=f32f947d6eafe80ca5ea7cd15140bea9214606c83487a520090cd1209e98171f
// right access token with less fields: http://localhost:8080/getAccountInfo?access_token=f32f947d6eafe80ca5ea7cd15140bea9214606c83487a520090cd1209e98171f&fields=[%22short_name%22,%22page_count%22]
// right access token with more fields: http://localhost:8080/getAccountInfo?access_token=f32f947d6eafe80ca5ea7cd15140bea9214606c83487a520090cd1209e98171f&fields=[%22short_name%22,%22page_count%22,%22auth_url%22]
// wrong access token: http://localhost:8080/getAccountInfo?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722&fields=%5B%22short_name%22,%22page_count%22%5D
func GetAccountInfoGetHandler(c *fiber.Ctx) error {
	accessToken := c.Query("access_token")
	// fields (Array of String, default = [“short_name”,“author_name”,“author_url”])
	// List of account fields to return. Available fields: short_name, author_name, author_url, auth_url, page_count.
	fields := c.Query("fields", "[\"short_name\", \"author_name\", \"author_url\"]")

	if accessToken == "" {
		return c.JSON(&model.Response{
			Ok:    false,
			Error: "ACCESS_TOKEN_INVALID",
		})
	}

	if accountEntity, err := database.EntClient.Account.Query().Where(account.AccessTokenEQ(accessToken)).Only(context.Background()); err == nil {
		var props []string // [short_name, author_name, author_url]
		err := json.Unmarshal([]byte(fields), &props)
		if err != nil {
			return c.Status(500).JSON(&model.Response{
				Ok:    false,
				Error: "FIELDS_FORMAT_INVALID",
			})
		}

		tmp := map[string]any{}
		result := map[string]any{}
		tmp["short_name"] = accountEntity.ShortName
		tmp["author_name"] = accountEntity.AuthorName
		tmp["auth_url"] = accountEntity.AuthURL
		tmp["author_url"] = accountEntity.AuthorURL
		tmp["page_count"] = accountEntity.PageCount

		for _, k := range props {
			if value, ok := tmp[k]; ok {
				result[k] = value
			}
		}

		return c.JSON(&model.Response{
			Ok:     true,
			Result: result,
		})
	} else {
		return c.Status(500).JSON(&model.Response{
			Ok:    false,
			Error: fmt.Sprintf("%v", err),
		})
	}
}

func GetAccountInfoPostHandler(c *fiber.Ctx) error {
	payload := new(model.AccountPayload)

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	// fields (Array of String, default = [“short_name”,“author_name”,“author_url”])
	// List of account fields to return. Available fields: short_name, author_name, author_url, auth_url, page_count.

	if payload.AccessToken == "" {
		return c.JSON(&model.Response{
			Ok:    false,
			Error: "ACCESS_TOKEN_INVALID",
		})
	}

	if strings.TrimSpace(payload.Fields) == "" {
		payload.Fields = "[\"short_name\", \"author_name\", \"author_url\"]"
	}

	if accountEntity, err := database.EntClient.Account.Query().Where(account.AccessTokenEQ(payload.AccessToken)).Only(context.Background()); err == nil {
		var props []string // [short_name, author_name, author_url]
		err := json.Unmarshal([]byte(payload.Fields), &props)
		if err != nil {
			return c.Status(500).JSON(&model.Response{
				Ok:    false,
				Error: "FIELDS_FORMAT_INVALID",
			})
		}

		tmp := map[string]any{}
		result := map[string]any{}
		tmp["short_name"] = accountEntity.ShortName
		tmp["author_name"] = accountEntity.AuthorName
		tmp["auth_url"] = accountEntity.AuthURL
		tmp["author_url"] = accountEntity.AuthorURL
		tmp["page_count"] = accountEntity.PageCount

		for _, k := range props {
			if value, ok := tmp[k]; ok {
				result[k] = value
			}
		}

		return c.JSON(&model.Response{
			Ok:     true,
			Result: result,
		})
	} else {
		return c.Status(500).JSON(&model.Response{
			Ok:    false,
			Error: fmt.Sprintf("%v", err),
		})
	}
}

// http://localhost:8080/revokeAccessToken?access_token=f32f947d6eafe80ca5ea7cd15140bea9214606c83487a520090cd1209e98171f
func RevokeAccessToken(c *fiber.Ctx) error {
	accessToken := c.Query("access_token")

	if accessToken == "" {
		return c.JSON(&model.Response{
			Ok:    false,
			Error: "ACCESS_TOKEN_INVALID",
		})
	}

	if data, err := database.EntClient.Account.Query().Where(account.AccessTokenEQ(accessToken)).Only(context.Background()); err == nil {
		newAccessToken := utils.GenerateSecureToken(32)

		tmp, err := data.Update().SetAccessToken(newAccessToken).SetAuthURL(fmt.Sprintf("http://localhost:8080/auth/%v", newAccessToken)).Save(context.Background())

		if err != nil {
			return c.Status(500).JSON(&model.Response{
				Ok:    false,
				Error: fmt.Sprintf("%v", err),
			})
		}

		result := struct {
			AccessToken string `json:"access_token"`
			AuthURL     string `json:"auth_url"`
		}{
			AccessToken: tmp.AccessToken,
			AuthURL:     tmp.AuthURL,
		}

		return c.JSON(&model.Response{
			Ok:     true,
			Result: result,
		})
	} else {
		return c.Status(500).JSON(&model.Response{
			Ok:    false,
			Error: fmt.Sprintf("%v", err),
		})
	}
}
