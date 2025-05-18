package comment

import ()

func CheckContent(content string) error {
	if len(content) < 1 {
		return ErrorContentNotEmpty
	}
	return nil
}
