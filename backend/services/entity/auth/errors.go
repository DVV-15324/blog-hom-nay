package auth

import (
	"errors"
)

var (

	//Kiểm tra valid
	ErrorPasswordNotValid  = errors.New("password phải có độ dài hơn 2 kí tự")
	ErrorEmailNotValid     = errors.New("email khong hop le")
	ErrorFirstNameNotValid = errors.New("lirstName tối thiểu 2 kí tự và nhỏ hơn 20 kí tự")
	ErrorLastNameNotValid  = errors.New("lastname tối thiểu 2 kí tự và nhỏ hơn 20 kí tự")
	ErrorPhoneNotValid     = errors.New("số điện thoại không hợp lệ")
	ErrorGenderNotValid    = errors.New("gender khong hop le")

	// Kiểm tra giá trị trống
	ErrorEmailIsNotEmpty     = errors.New("email không được trống")
	ErrorPasswordIsNotEmpty  = errors.New("password không được trống")
	ErrorFirstNameIsNotEmpty = errors.New("firstname không được trống")
	ErrorLastNameIsNotEmpty  = errors.New("lastName không được trống")
	ErrorPhoneIsNotEmpty     = errors.New("số điện thoại không được trống")

	// Kiểm tra xác thực
	ErrorEmailIsExisted    = errors.New("email đã tồn tại")                 // kiểm tra lúc đăng kí
	ErrorEmailAndPassword  = errors.New("tai Khoan va mat khau khong dung") // kiểm tra lúc đăng kí
	ErrorEmailIsNotExisted = errors.New("email không tồn tại")              // kiểm tra lúc đăng nhập
)
