package comment

import "strings"

type CreateComment struct {
	PostID  string `json:"-"`
	UserID  string `json:"-"`
	Content string `json:"content"`
}

func (c *CreateComment) Validate() error {
	c.Content = strings.TrimSpace(c.Content)
	er := CheckContent(c.Content)
	if er != nil {
		return er
	}
	return nil
}

type UpdateComment struct {
	Content *string `json:"content"`
}

func (c *UpdateComment) Validate() error {
	*c.Content = strings.TrimSpace(*c.Content)
	er := CheckContent(*c.Content)
	if er != nil {
		return er
	}
	return nil
}
