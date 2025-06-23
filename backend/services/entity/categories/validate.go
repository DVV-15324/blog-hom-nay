package categories

import ()

func CheckDiscribe(describe string) error {
	if len(describe) < 1 {
		return ErrorDiscribeNotEmpty
	}
	return nil
}

func CheckName(name string) error {
	if len(name) < 1 {
		return ErrorNameNotEmpty
	}
	return nil
}
func CheckImg(name string) error {
	if len(name) < 1 {
		return ErrorImgNotEmpty
	}
	return nil
}
