package categories

import "strings"

type CreateCategory struct {
	Description string `json:"description"`
	Name        string `json:"name"`
<<<<<<< HEAD
	Img         string `json:"img"`
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
}

func (c *CreateCategory) Validate() error {
	c.Description = strings.TrimSpace(c.Description)
	c.Name = strings.TrimSpace(c.Name)
<<<<<<< HEAD
	c.Img = strings.TrimSpace(c.Img)
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	err_des := CheckDiscribe(c.Description)
	if err_des != nil {
		return err_des
	}
	err_name := CheckName(c.Name)
	if err_name != nil {
		return err_name
	}
<<<<<<< HEAD
	err_img := CheckImg(c.Description)
	if err_img != nil {
		return err_img
	}
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	return nil
}

type UpdateCategory struct {
	Description *string `json:"description"`
	Name        *string `json:"name"`
<<<<<<< HEAD
	Img         *string `json:"img"`
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
}

func (u *UpdateCategory) Validate() error {
	*u.Description = strings.TrimSpace(*u.Description)
	*u.Name = strings.TrimSpace(*u.Name)
<<<<<<< HEAD
	*u.Img = strings.TrimSpace(*u.Img)

=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	err_des := CheckDiscribe(*u.Description)
	if err_des != nil {
		return err_des
	}
<<<<<<< HEAD
	err_img := CheckImg(*u.Description)
	if err_img != nil {
		return err_img
	}
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	err_name := CheckName(*u.Name)
	if err_name != nil {
		return err_name
	}
	return nil
}
