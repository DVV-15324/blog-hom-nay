package user

import "strings"

type CreateUserForm struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

func (c *CreateUserForm) Validate() error {
	c.FirstName = strings.TrimSpace(c.FirstName)
	c.LastName = strings.TrimSpace(c.LastName)
	c.Phone = strings.TrimSpace(c.Phone)
	c.Email = strings.TrimSpace(c.Email)
	//Checks
	err_email := CheckEmail(c.Email)
	if err_email != nil {
		return err_email
	}
	err_phone := CheckPhone(c.Phone)
	if err_phone != nil {
		return err_phone
	}
	err_firstname := CheckFirstName(c.FirstName)
	if err_firstname != nil {
		return err_firstname
	}
	err_lastname := CheckLastName(c.LastName)
	if err_lastname != nil {
		return err_lastname
	}
	return nil
}

type UpdateUserForm struct {
	Avatar    *string `json:"avatar"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Phone     *string `json:"phone"`
	Address   *string `json:"address"`
}

func (c *UpdateUserForm) Validate() error {
	if c.FirstName != nil {
		*c.FirstName = strings.TrimSpace(*c.FirstName)
		if err := CheckFirstName(*c.FirstName); err != nil {
			return err
		}
	}

	if c.LastName != nil {
		*c.LastName = strings.TrimSpace(*c.LastName)
		if err := CheckLastName(*c.LastName); err != nil {
			return err
		}
	}

	if c.Phone != nil {
		*c.Phone = strings.TrimSpace(*c.Phone)
		if err := CheckPhone(*c.Phone); err != nil {
			return err
		}
	}

	if c.Address != nil {
		*c.Address = strings.TrimSpace(*c.Address)
		if err := CheckAddress(*c.Address); err != nil {
			return err
		}
	}

	if c.Avatar != nil {
		*c.Avatar = strings.TrimSpace(*c.Avatar)
	}

	return nil
}
