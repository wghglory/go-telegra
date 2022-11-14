package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"telegra/database"
	"telegra/ent/account"
	"telegra/ent/page"
	"telegra/model"

	"github.com/gofiber/fiber/v2"
)

// http://localhost:8080/createPage?access_token=eec0b228c9fa28cc7dfd8dfbb84d47ad16a2d39e1ebf8c4f466d1acf9b443840&title=First&author_name=Derek+Wang&content=[%7B%22tag%22:%22p%22,%22children%22:[%22Hello,+world!%22]%7D]&return_content=true
func CreatePageGetHandler(c *fiber.Ctx) error {
	accessToken := c.Query("access_token")
	title := c.Query("title")
	path := fmt.Sprintf("%v-%v", title, time.Now().UnixNano())
	content := c.Query("content")

	if accessToken == "" {
		return c.JSON(&model.Response{
			Ok:    false,
			Error: "ACCESS_TOKEN_INVALID",
		})
	}

	if accountEntity, err := database.EntClient.Account.Query().Where(account.AccessTokenEQ(accessToken)).Only(context.Background()); err == nil {
		var nodes []model.NodeElement
		if err := json.Unmarshal([]byte(content), &nodes); err != nil {
			return c.JSON(&model.Response{
				Ok:    false,
				Error: "Fail to provide a valid DOM element",
			})
		}

		var canEdit bool = false
		if c.Query("return_content") == "true" {
			canEdit = true
		}

		pageEntity, err := database.EntClient.Page.
			Create().
			SetPath(path).
			SetTitle(title).
			SetURL(fmt.Sprintf("http://localhost:8080/%v", path)).
			SetAuthorName(c.Query("author_name")).
			SetAuthorURL(c.Query("author_url")).
			SetDescription(c.Query("description")).
			SetContent(content).
			SetViews(1).
			SetCanEdit(canEdit).
			Save(context.Background())

		if err != nil {
			return c.Status(500).JSON(&model.Response{
				Ok:    false,
				Error: fmt.Sprintf("failed creating a page: %v", err),
			})
		}
		accountEntity.Update().AddPages(pageEntity).Save(context.Background())

		page := &model.Page{
			Path:        pageEntity.Path,
			Url:         pageEntity.URL,
			Title:       pageEntity.Title,
			AuthorName:  pageEntity.AuthorName,
			AuthorUrl:   pageEntity.AuthorURL,
			Description: pageEntity.Description,
			Views:       pageEntity.Views,
			CanEdit:     pageEntity.CanEdit,
		}

		if c.Query("return_content") == "true" {
			page.Content = nodes
		}

		return c.JSON(&model.Response{
			Ok:     true,
			Result: page,
		})
	} else {
		return c.Status(500).JSON(&model.Response{
			Ok:    false,
			Error: fmt.Sprintf("%v", err),
		})
	}
}

func CreatePagePostHandler(c *fiber.Ctx) error {
	payload := new(model.PagePayload)

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	path := fmt.Sprintf("%v-%v", payload.Title, time.Now().UnixNano())

	if payload.AccessToken == "" {
		return c.JSON(&model.Response{
			Ok:    false,
			Error: "ACCESS_TOKEN_INVALID",
		})
	}

	if accountEntity, err := database.EntClient.Account.Query().Where(account.AccessTokenEQ(payload.AccessToken)).Only(context.Background()); err == nil {

		var nodes []model.NodeElement
		content, _ := url.PathUnescape(payload.Content)
		if err := json.Unmarshal([]byte(content), &nodes); err != nil {
			return c.JSON(&model.Response{
				Ok:    false,
				Error: "DOM_NODE_INVALID",
			})
		}

		var canEdit bool = false
		if payload.ReturnContent == "true" {
			canEdit = true
		}

		pageEntity, err := database.EntClient.Page.
			Create().
			SetPath(path).
			SetTitle(payload.Title).
			SetURL(fmt.Sprintf("http://localhost:8080/%v", path)).
			SetAuthorName(payload.AuthorName).
			SetAuthorURL(payload.AuthorUrl).
			SetDescription(payload.Description).
			SetContent(string(payload.Content)).
			SetViews(1).
			SetCanEdit(canEdit).
			Save(context.Background())

		if err != nil {
			return c.Status(500).JSON(&model.Response{
				Ok:    false,
				Error: fmt.Sprintf("failed creating a page: %v", err),
			})
		}
		accountEntity.Update().AddPages(pageEntity).Save(context.Background())

		page := &model.Page{
			Path:        pageEntity.Path,
			Url:         pageEntity.URL,
			Title:       pageEntity.Title,
			AuthorName:  pageEntity.AuthorName,
			AuthorUrl:   pageEntity.AuthorURL,
			Description: pageEntity.Description,
			Views:       pageEntity.Views,
			CanEdit:     pageEntity.CanEdit,
		}

		if payload.ReturnContent == "true" {
			page.Content = nodes
		}

		return c.JSON(&model.Response{
			Ok:     true,
			Result: page,
		})
	} else {
		return c.Status(500).JSON(&model.Response{
			Ok:    false,
			Error: fmt.Sprintf("%v", err),
		})
	}
}

// http://localhost:8080/getPage/First-1668242327838759000?return_content=true
// http://localhost:8080/getPage/First-1668242327838759000?return_content=false
func GetPage(c *fiber.Ctx) error {
	path, _ := url.PathUnescape(c.Params("path"))

	if pageEntity, err := database.EntClient.Page.Query().Where(page.PathEQ(path)).Only(context.Background()); err != nil {
		return c.Status(500).JSON(&model.Response{
			Ok:    false,
			Error: fmt.Sprintf("failed to find a page with path: %v", path),
		})
	} else {
		var nodes []model.NodeElement
		result := model.Page{
			Path:        pageEntity.Path,
			Url:         pageEntity.URL,
			Title:       pageEntity.Title,
			Description: pageEntity.Description,
			AuthorName:  pageEntity.AuthorName,
			ImageUrl:    pageEntity.ImageURL,
			Views:       pageEntity.Views,
			CanEdit:     pageEntity.CanEdit,
		}

		if c.Query("return_content") == "true" {
			err := json.Unmarshal([]byte(pageEntity.Content), &nodes)

			if err != nil {
				return c.JSON(&model.Response{
					Ok:    false,
					Error: "Fail to provide a valid DOM element",
				})
			}

			result.Content = nodes
		}

		return c.JSON(&model.Response{
			Ok:     true,
			Result: result,
		})
	}

}

// Not a valid path: http://localhost:8080/editPage/Sample-Page-12-15?access_token=eec0b228c9fa28cc7dfd8dfbb84d47ad16a2d39e1ebf8c4f466d1acf9b443840&title=Sample+Page&author_name=Derek+Wang&content=[%7B%22tag%22:%22div%22,%22children%22:[%22Hello,+New+world!%22]%7D]&return_content=true
// valid but not content: http://localhost:8080/editPage/First-1668242327838759000?access_token=eec0b228c9fa28cc7dfd8dfbb84d47ad16a2d39e1ebf8c4f466d1acf9b443840&title=Sample+Page&author_name=Derek+Wang&content=[%7B%22tag%22:%22div%22,%22children%22:[%22Hello,+New+world!%22]%7D]
// valid with content: http://localhost:8080/editPage/First-1668242327838759000?access_token=eec0b228c9fa28cc7dfd8dfbb84d47ad16a2d39e1ebf8c4f466d1acf9b443840&title=Sample+Page&author_name=Derek+Wang&content=[%7B%22tag%22:%22div%22,%22children%22:[%22Hello,+New+world!%22]%7D]&return_content=true
func EditPage(c *fiber.Ctx) error {
	accessToken := c.Query("access_token")
	title := c.Query("title")
	path, _ := url.PathUnescape(c.Params("path"))
	content := c.Query("content")
	description := c.Query("description")
	authorName := c.Query("author_name")

	if accessToken == "" {
		return c.JSON(&model.Response{
			Ok:    false,
			Error: "ACCESS_TOKEN_INVALID",
		})
	}

	if _, err := database.EntClient.Account.Query().Where(account.AccessTokenEQ(accessToken)).Only(context.Background()); err == nil {

		var nodes []model.NodeElement
		if err := json.Unmarshal([]byte(content), &nodes); err != nil {
			return c.JSON(&model.Response{
				Ok:    false,
				Error: "Fail to provide a valid DOM element",
			})
		}

		var canEdit bool = false
		if c.Query("return_content") == "true" {
			canEdit = true
		}

		pageEntity, err := database.EntClient.Page.Query().Where(page.PathEQ(path)).Only(context.Background())

		if err != nil {
			return c.JSON(&model.Response{
				Ok:    false,
				Error: "Cannot find the path",
			})
		}

		pageEntity, err = pageEntity.
			Update().
			SetTitle(title).
			SetAuthorName(authorName).
			SetDescription(description).
			SetContent(content).
			SetViews(pageEntity.Views + 1).
			SetCanEdit(canEdit).
			Save(context.Background())

		if err != nil {
			return c.Status(500).JSON(&model.Response{
				Ok:    false,
				Error: fmt.Sprintf("failed to update a page: %v", err),
			})
		}

		page := &model.Page{
			Path:        pageEntity.Path,
			Title:       pageEntity.Title,
			AuthorName:  pageEntity.AuthorName,
			Url:         pageEntity.URL,
			Description: pageEntity.Description,
			Content:     nil,
			Views:       pageEntity.Views,
			CanEdit:     pageEntity.CanEdit,
		}

		if c.Query("return_content") == "true" {
			page.Content = nodes
		}

		return c.JSON(&model.Response{
			Ok:     true,
			Result: page,
		})
	} else {
		return c.Status(500).JSON(&model.Response{
			Ok:    false,
			Error: fmt.Sprintf("%v", err),
		})
	}

}
