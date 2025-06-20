package comment

import "strings"

type CreateComment struct {
<<<<<<< HEAD
	PostID     int    `json:"-"`
	UserID     int    `json:"-"`
	FakePostID string `json:"post_id"`
	Content    string `json:"content"`
=======
	PostID  string `json:"-"`
	UserID  string `json:"-"`
	Content string `json:"content"`
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
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
<<<<<<< HEAD
	PostID     *int    `json:"-"`
	UserID     *int    `json:"-"`
	FakePostID *string `json:"post_id"`
	Content    *string `json:"content"`
=======
	Content *string `json:"content"`
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
}

func (c *UpdateComment) Validate() error {
	*c.Content = strings.TrimSpace(*c.Content)
	er := CheckContent(*c.Content)
	if er != nil {
		return er
	}
	return nil
}
