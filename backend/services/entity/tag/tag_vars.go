package tag

import "strings"

type CreateTagForm struct {
	Name string `json:"name"`
}

func (t *CreateTagForm) Validate() error {
	t.Name = strings.TrimSpace(t.Name)
	err_name := CheckName(t.Name)
	if err_name != nil {
		return err_name
	}
	return nil
}

type UpdateTagForm struct {
	Name *string `json:"name"`
}

func (t *UpdateTagForm) Validate() error {
	*t.Name = strings.TrimSpace(*t.Name)
	err_name := CheckName(*t.Name)
	if err_name != nil {
		return err_name
	}
	return nil
}
