package user

import (
	"bloghomnay-project/common"
	"fmt"
	"regexp"
	"strings"
)

func CheckEmail(email string) error {
	if len(email) < 1 {
		return ErrorEmailIsNotEmpty
	}
	_, err := regexp.MatchString(`^[0-9a-zA-Z.\-+]+@[0-9a-zA-Z]+\.[0-9a-zA-Z]{2,}$`, email)
	if err != nil {
		return ErrorEmailNotValid
	}
	return nil
}

func CheckFirstName(firstName string) error {
	if len(firstName) < 1 {
		return ErrorFirstNameIsNotEmpty
	}
	if len(firstName) < 2 || len(firstName) > 20 {
		return ErrorFirstNameNotValid
	}
	return nil
}

func CheckLastName(lastName string) error {
	if len(lastName) < 1 {
		return ErrorFirstNameIsNotEmpty
	}
	if len(lastName) < 2 || len(lastName) > 20 {
		return ErrorFirstNameNotValid
	}
	return nil
}

func CheckAddress(address string) error {
	splitAddress := strings.Split(address, "-")
	//ví dụ: Xã Lãng Ngâm - Huyện Gia Bình - Tỉnh Bắc Ninh
	err := common.Checkprovince(splitAddress[2], splitAddress[1], splitAddress[0])
	if err != nil {
		return ErrorAddressNotValid
	}
	return nil
}

func CheckPhone(phone string) error {
	// Regex đầu số các nhà mạng tại Việt Nam
	viettel := `^(0|\+84|84)(32|33|34|35|36|37|38|39|96|97|98|86)[0-9]{7}$`
	vinaphone := `^(0|\+84|84)(81|82|83|84|85|88|91|94)[0-9]{7}$`
	mobifone := `^(0|\+84|84)(70|76|77|78|79|89|90|93)[0-9]{7}$`
	vietnamobile := `^(0|\+84|84)(52|56|58|92)[0-9]{7}$`
	gmobile := `^(0|\+84|84)(59|99)[0-9]{7}$`

	// Kết hợp tất cả regex thành một
	pattern := fmt.Sprintf("(%s)|(%s)|(%s)|(%s)|(%s)", viettel, vinaphone, mobifone, vietnamobile, gmobile)

	match, _ := regexp.MatchString(pattern, phone)
	if !match {
		return ErrorPhoneNotValid
	}
	return nil
}
