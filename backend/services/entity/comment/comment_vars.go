package comment

import "strings"

type CreateComment struct {
	PostID     int    `json:"-"`
	UserID     int    `json:"-"`
	FakePostID string `json:"post_id"`
	Content    string `json:"content"`
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
	PostID     *int    `json:"-"`
	UserID     *int    `json:"-"`
	FakePostID *string `json:"post_id"`
	Content    *string `json:"content"`
}

func (c *UpdateComment) Validate() error {
	*c.Content = strings.TrimSpace(*c.Content)
	er := CheckContent(*c.Content)
	if er != nil {
		return er
	}
	return nil
}
