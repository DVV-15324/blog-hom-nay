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
