package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"telegra/database"
	"telegra/ent/account"
	"telegra/model"

	"github.com/gofiber/fiber/v2"
)

// http://localhost:8080/createPage?access_token=eec0b228c9fa28cc7dfd8dfbb84d47ad16a2d39e1ebf8c4f466d1acf9b443840&title=First+Page&author_name=Derek+Wang&content=[%7B%22tag%22:%22p%22,%22children%22:[%22Hello,+world!%22]%7D]&return_content=true
func CreatePage(c *fiber.Ctx) error {
	accessToken := c.Query("access_token")
	title := c.Query("title")
	path := fmt.Sprintf("%v-%v", title, time.Now().UnixNano())
	content := c.Query("content")

	if accessToken == "" {
		return c.JSON(&model.Response{
			Ok:    false,
			Error: "Not a valid access token",
		})
	}

	if accountEntity, err := database.EntClient.Account.Query().Where(account.AccessTokenEQ(accessToken)).Only(context.Background()); err == nil {

		page := &model.Page{
			Path:        path,
			Title:       title,
			AuthorName:  c.Query("author_name"),
			Url:         fmt.Sprintf("http://localhost:8080/%v", path),
			Description: c.Query("description"),
			Content:     nil,
			Views:       1,
			CanEdit:     false,
		}

		var nodes []model.NodeElement
		if err := json.Unmarshal([]byte(content), &nodes); err != nil {
			return c.JSON(&model.Response{
				Ok:    false,
				Error: "Fail to provide a valid DOM element",
			})
		}

		if c.Query("return_content") == "true" {
			page.Content = nodes
			page.CanEdit = true
		}

		pageEntity, err := database.EntClient.Page.
			Create().
			SetPath(page.Path).
			SetTitle(page.Title).
			SetAuthorName(page.AuthorName).
			SetURL(page.Url).
			SetDescription(page.Description).
			SetContent(content).
			SetViews(page.Views).
			SetCanEdit(page.CanEdit).
			Save(context.Background())

		if err != nil {
			return c.Status(500).JSON(&model.Response{
				Ok:    false,
				Error: fmt.Sprintf("failed creating a page: %v", err),
			})
		}

		accountEntity.Update().AddPages(pageEntity).Save(context.Background())

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

// http://localhost:8080/getPage/Sample-Page-12-15?return_content=true
func GetPage(c *fiber.Ctx) error {
	path := c.Params("path")
	title := "New Title"
	content := "[{\"tag\":\"div\",\"children\":[\"Hello, New world!\"]}]"

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

// http://localhost:8080/editPage/Sample-Page-12-15?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722&title=Sample+Page&author_name=Derek+Wang&content=[%7B%22tag%22:%22div%22,%22children%22:[%22Hello,+New+world!%22]%7D]&return_content=true
// http://localhost:8080/editPage/Sample-Page-12-15?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722&title=Sample+Page&author_name=Derek+Wang&content=[%7B%22tag%22:%22div%22,%22children%22:[%22Hello,+New+world!%22]%7D]&return_content=false
func EditPage(c *fiber.Ctx) error {
	path := c.Params("path")
	accessToken := c.Query("access_token")
	title := c.Query("title")
	content := c.Query("content")

	return pageHandler(c, accessToken, path, title, content)
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
