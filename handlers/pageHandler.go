package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"telegra/model"

	"github.com/gofiber/fiber/v2"
)

// http://localhost:8080/createPage?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722&title=Sam%E5%8F%91%E9%A1%BA%E4%B8%B0ple+Page&author_name=Anonymous&content=[%7B%22tag%22:%22p%22,%22children%22:[%22Hello,+world!%22]%7D]&return_content=true
func CreatePage(c *fiber.Ctx) error {
	accessToken := c.Query("access_token")
	title := c.Query("title")
	path := fmt.Sprintf("%v-%v", title, time.Now())
	content := c.Query("content")

	return pageHandler(c, accessToken, path, title, content)
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
