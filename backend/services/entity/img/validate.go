package img

func CheckImg(img string) error {
	if len(img) < 1 {
		return ErrorImgNotValid
	}
	return nil
}
