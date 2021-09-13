package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	phonenumber       string = "081291755506"
	phonenumberwith62 string = "6281291755506"
)

var user = User{
	PhoneNumber: "081291755506",
	Email:       "eldhoral@gmai.com",
}

func TestValidatePhoneNumber(t *testing.T) {
	assert := assert.New(t)
	TestPhoneNumber := ValidatePhoneNumber("081291755506")
	//checking the result of the ValidatePhoneNumber function
	assert.Equal(true, TestPhoneNumber, "PASS")
	assert.NotEqual(false, TestPhoneNumber, "PASS")
}

func TestChangePhoneNumberFormat(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(phonenumberwith62, user.ChangePhoneNumberFormat(), "PASS")
	assert.NotEqual(phonenumber, user.ChangePhoneNumberFormat(), "PASS")
}

func TestValidateEmail(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(true, user.ValidateEmail(), "PASS")
	assert.NotEqual(false, user.ValidateEmail(), "PASS")
}
