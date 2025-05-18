package posts

import "errors"

var (
	ErrorTitleNotEmpty   = errors.New("title khong hop le")
	ErrorContentNotEmpty = errors.New("content khong hop le")
)
