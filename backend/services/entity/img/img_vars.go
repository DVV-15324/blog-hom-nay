package img

type CreateImg struct {
	UserId int    `json:"-"`
	Img    string `json:"img"`
}

func (c *CreateImg) Validate() error {
	err := CheckImg(c.Img)
	if err != nil {
		return err
	}
	return nil
}

type UpdateImg struct {
	UserId *string `json:"-"`
	Img    *string `json:"img"`
}

func (c *UpdateImg) Validate() error {
	err := CheckImg(*c.Img)
	if err != nil {
		return err
	}
	return nil
}
