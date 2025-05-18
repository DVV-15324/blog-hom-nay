package posts

import (
	"bloghomnay-project/common"
	"strings"
)

type CreatePost struct {
	FakeCategoryId string               `json:"categories_id"`
	CategoryId     int                  `json:"-"`
	UserId         int                  `json:"-"`
	FakeUserId     string               `json:"user_id"`
	Title          string               `json:"title"`
	Content        string               `json:"content"`
	Tag            []common.TagFormBase `json:"tags"`
}

func (c *CreatePost) Validate() error {
	c.Title = strings.TrimSpace(c.Title)
	c.Content = strings.TrimSpace(c.Content)
	err_content := CheckContent(c.Content)
	if err_content != nil {
		return err_content
	}
	err_title := CheckTile(c.Title)
	if err_title != nil {
		return err_title
	}
	return nil
}

type UpdatePost struct {
	FakeCategoryId *string              `json:"categories_id"`
	CategoryId     *int                 `json:"-"`
	Title          *string              `json:"title"`
	Content        *string              `json:"content"`
	Tag            []common.TagFormBase `json:"tags"`
}

func (u *UpdatePost) Validate() error {
	*u.Title = strings.TrimSpace(*u.Title)
	*u.Content = strings.TrimSpace(*u.Content)
	err_content := CheckContent(*u.Content)
	if err_content != nil {
		return err_content
	}
	err_title := CheckTile(*u.Title)
	if err_title != nil {
		return err_title
	}
	return nil
}
