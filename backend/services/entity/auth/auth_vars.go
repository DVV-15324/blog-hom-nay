package auth

import "strings"

type RegisterForm struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

func (r *RegisterForm) Validate() error {
	r.Email = strings.TrimSpace(r.Email)
	r.Password = strings.TrimSpace(r.Password)
	r.FirstName = strings.TrimSpace(r.FirstName)
	r.LastName = strings.TrimSpace(r.LastName)
	r.Phone = strings.TrimSpace(r.Phone)
	//checks
	err_email := CheckEmail(r.Email)
	if err_email != nil {
		return err_email
	}
	err_password := CheckPasword(r.Password)
	if err_password != nil {
		return err_password
	}
	err_firstname := CheckFirstName(r.FirstName)
	if err_firstname != nil {
		return err_firstname
	}
	err_lastname := CheckLastName(r.LastName)
	if err_lastname != nil {
		return err_lastname
	}

	err_phone := CheckPhone(r.Phone)
	if err_phone != nil {
		return err_phone
	}
	return nil
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginForm) Validate() error {
	r.Email = strings.TrimSpace(r.Email)
	r.Password = strings.TrimSpace(r.Password)

	err_email := CheckEmail(r.Email)
	if err_email != nil {
		return err_email
	}
	err_password := CheckPasword(r.Password)
	if err_password != nil {
		return err_password
	}

	return nil
}
