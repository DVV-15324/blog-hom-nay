package categories

import "strings"

type CreateCategory struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Img         string `json:"img"`
}

func (c *CreateCategory) Validate() error {
	c.Description = strings.TrimSpace(c.Description)
	c.Name = strings.TrimSpace(c.Name)
	c.Img = strings.TrimSpace(c.Img)
	err_des := CheckDiscribe(c.Description)
	if err_des != nil {
		return err_des
	}
	err_name := CheckName(c.Name)
	if err_name != nil {
		return err_name
	}
	err_img := CheckImg(c.Description)
	if err_img != nil {
		return err_img
	}
	return nil
}

type UpdateCategory struct {
	Description *string `json:"description"`
	Name        *string `json:"name"`
	Img         *string `json:"img"`
}

func (u *UpdateCategory) Validate() error {
	*u.Description = strings.TrimSpace(*u.Description)
	*u.Name = strings.TrimSpace(*u.Name)
	*u.Img = strings.TrimSpace(*u.Img)

	err_des := CheckDiscribe(*u.Description)
	if err_des != nil {
		return err_des
	}
	err_img := CheckImg(*u.Description)
	if err_img != nil {
		return err_img
	}
	err_name := CheckName(*u.Name)
	if err_name != nil {
		return err_name
	}
	return nil
}
