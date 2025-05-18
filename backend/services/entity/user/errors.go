package user

import (
	"errors"
)

var (

	//Kiểm tra valid

	ErrorEmailNotValid     = errors.New("email khong hop le")
	ErrorFirstNameNotValid = errors.New("lirstName tối thiểu 2 kí tự và nhỏ hơn 20 kí tự")
	ErrorLastNameNotValid  = errors.New("lastname tối thiểu 2 kí tự và nhỏ hơn 20 kí tự")
	ErrorPhoneNotValid     = errors.New("số điện thoại không hợp lệ")

	// Kiểm tra giá trị trống
	ErrorEmailIsNotEmpty     = errors.New("email không được trống")
	ErrorFirstNameIsNotEmpty = errors.New("firstname không được trống")
	ErrorLastNameIsNotEmpty  = errors.New("lastName không được trống")
	ErrorPhoneIsNotEmpty     = errors.New("số điện thoại không được trống")

	//
	ErrorAddressNotValid = errors.New("địa chỉ không hợp lệ")
)
