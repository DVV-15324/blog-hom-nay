package posts

import ()

func CheckTile(content string) error {
	if len(content) < 1 {
		return ErrorContentNotEmpty
	}
	return nil
}
func CheckContent(content string) error {
	if len(content) < 1 {
		return ErrorTitleNotEmpty
	}
	return nil
}
<<<<<<< HEAD
func CheckDescription(content string) error {
	if len(content) < 1 {
		return ErrorDescriptionNotEmpty
	}
	return nil
}
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
