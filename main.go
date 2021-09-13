package main

import (
	"fmt"
	"regexp"
	"strings"
)

type User struct {
	PhoneNumber, Email string
}

func main() {
	var u = User{}
	fmt.Println("Input your Phone Number: ")
	fmt.Scanln(&u.PhoneNumber)
	fmt.Println("Input your Email Adress: ")
	fmt.Scanln(&u.Email)
	DataPhoneNumber := u.ChangePhoneNumberFormat()
	DataEmail := u.ValidateEmail()
	//fmt.Println("Your Phone Number is ", DataPhoneNumber)
	//fmt.Println("Your Email Adress is ", DataEmail)

	//Creating error checking message for phone number
	if ValidatePhoneNumber(DataPhoneNumber) {
		fmt.Println("The format that you inputted is correct", DataPhoneNumber)
	} else {
		fmt.Println("The phone number format that you inputted is wrong")
	}

	//Creating error checking message for email
	if DataEmail {
		fmt.Println("The format that you inputted is correct", DataEmail)
	} else {
		fmt.Println("The email format that you inputted is wrong")
	}

}

func ValidatePhoneNumber(input string) bool {
	//The input only accept 0-9, with min of 9 and max of 14
	CheckConstraint := regexp.MustCompile(`^[0-9]{9,14}$`).MatchString
	return CheckConstraint(input)
}

func (u User) ChangePhoneNumberFormat() string {
	//Eliminate any space in phone number
	Phonenumber := strings.Replace(u.PhoneNumber, " ", "", -1)
	//Eliminate any "+" character in front of the Phone number
	if strings.HasPrefix(Phonenumber, "+") {
		Phonenumber = strings.Replace(Phonenumber, "+", "", 1)
		//Eliminate number "0" in front of the phone number and change it to "62"
	} else if strings.HasPrefix(Phonenumber, "0") {
		Phonenumber = strings.Replace(Phonenumber, "0", "62", 1)
	}
	return Phonenumber
}

func (u User) ValidateEmail() bool {
	//Change format of the inputted email and validate email. Only accept with @domain.com/id/etc
	Email := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString
	return Email(u.Email)
}
