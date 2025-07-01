package categories

import "errors"

var (
	ErrorDiscribeNotEmpty = errors.New("mieu ta khong duoc trong")
	ErrorNameNotEmpty     = errors.New("name khong duoc trong")
	ErrorImgNotEmpty      = errors.New("img khong duoc trong")
)
