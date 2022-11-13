package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"telegra/database"
	"telegra/ent/account"
	"telegra/ent/page"
	"telegra/model"

	"github.com/gofiber/fiber/v2"
)

// http://localhost:8080/createPage?access_token=eec0b228c9fa28cc7dfd8dfbb84d47ad16a2d39e1ebf8c4f466d1acf9b443840&title=First&author_name=Derek+Wang&content=[%7B%22tag%22:%22p%22,%22children%22:[%22Hello,+world!%22]%7D]&return_content=true
func CreatePage(c *fiber.Ctx) error {
	accessToken := c.Query("access_token")
	title := c.Query("title")
	path := fmt.Sprintf("%v-%v", title, time.Now().UnixNano())
	content := c.Query("content")

	if accessToken == "" {
		return c.JSON(&model.Response{
			Ok:    false,
			Error: "Please provide a valid access token",
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

// http://localhost:8080/getPage/First-1668242327838759000?return_content=true
// http://localhost:8080/getPage/First-1668242327838759000?return_content=false
func GetPage(c *fiber.Ctx) error {
	path := c.Params("path")

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
	path := c.Params("path")
	accessToken := c.Query("access_token")
	title := c.Query("title")
	content := c.Query("content")
	description := c.Query("description")
	authorName := c.Query("author_name")

	if accessToken == "" {
		return c.JSON(&model.Response{
			Ok:    false,
			Error: "Please provide a valid access token",
		})
	}

	if _, err := database.EntClient.Account.Query().Where(account.AccessTokenEQ(accessToken)).Only(context.Background()); err == nil {

		pageEntity, err := database.EntClient.Page.Query().Where(page.PathEQ(path)).Only(context.Background())

		if err != nil {
			return c.JSON(&model.Response{
				Ok:    false,
				Error: "Cannot find the path",
			})
		}

		var nodes []model.NodeElement
		if err := json.Unmarshal([]byte(content), &nodes); err != nil {
			return c.JSON(&model.Response{
				Ok:    false,
				Error: "Fail to provide a valid DOM element",
			})
		}

		pageEntity, err = pageEntity.
			Update().
			SetTitle(title).
			SetAuthorName(authorName).
			SetDescription(description).
			SetContent(content).
			SetViews(pageEntity.Views + 1).
			SetCanEdit(true).
			Save(context.Background())

		if err != nil {
			return c.Status(500).JSON(&model.Response{
				Ok:    false,
				Error: fmt.Sprintf("failed to update a page with path: %v", path),
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
