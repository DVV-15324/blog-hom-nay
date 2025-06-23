package posts

import (
	"bloghomnay-project/common"
	"strings"
)

type CreatePost struct {
	FakeCategoryId string               `json:"categories_id"`
	CategoryId     int                  `json:"-"`
	UserId         int                  `json:"-"`
	Title          string               `json:"title"`
	Content        string               `json:"content"`
	Description    string               `json:"description"`
	Tag            []common.TagFormBase `json:"tags"`
}

func (c *CreatePost) Validate() error {
	c.Title = strings.TrimSpace(c.Title)
	c.Content = strings.TrimSpace(c.Content)
	c.Description = strings.TrimSpace(c.Description)
	err_content := CheckContent(c.Content)
	if err_content != nil {
		return err_content
	}
	err_description := CheckDescription(c.Description)
	if err_description != nil {
		return err_description
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
	Description    *string              `json:"description"`
	Tag            []common.TagFormBase `json:"tags"`
}

func (u *UpdatePost) Validate() error {
	*u.Title = strings.TrimSpace(*u.Title)
	*u.Content = strings.TrimSpace(*u.Content)
<<<<<<< HEAD
	*u.Description = strings.TrimSpace(*u.Description)
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	err_content := CheckContent(*u.Content)
	if err_content != nil {
		return err_content
	}
<<<<<<< HEAD
	err_description := CheckDescription(*u.Description)
	if err_description != nil {
		return err_description
	}
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	err_title := CheckTile(*u.Title)
	if err_title != nil {
		return err_title
	}
	return nil
}
