package validator_test

import (
	"testing"

	"github.com/dgibs750/lynx/util/validator"
)

type testCase struct {
	name     string
	input    interface{}
	expected string
}

var tests = []*testCase{
	{
		name: `required`,
		input: struct {
			Title string `json:"title" validate:"required"`
		}{},
		expected: "Title is a required field",
	},
	{
		name: `email`,
		input: struct {
			Email string `json:"email" validate:"email"`
		}{Email: "Abc.example.com"},
		expected: "Email must be a valid email address",
	},
	{
		name: `email`,
		input: struct {
			Email string `json:"email" validate:"email"`
		}{Email: "A@b@c@example.com"},
		expected: "Email must be a valid email address",
	},
	{
		name: `email`,
		input: struct {
			Email string `json:"email" validate:"email"`
		}{Email: `a"b(c)d,e:f;g<h>i[j\'k]l@example.com`},
		expected: "Email must be a valid email address",
	},
	{
		name: `email`,
		input: struct {
			Email string `json:"email" validate:"email"`
		}{Email: `just"not"right@example.com`},
		expected: "Email must be a valid email address",
	},
	{
		name: `email`,
		input: struct {
			Email string `json:"email" validate:"email"`
		}{Email: `this is"not\allowed@example.com`},
		expected: "Email must be a valid email address",
	},
	{
		name: `email`,
		input: struct {
			Email string `json:"email" validate:"email"`
		}{Email: `this\ still\"notallowed@example.com`},
		expected: "Email must be a valid email address",
	},
	{
		name: `min`,
		input: struct {
			Password string `json:"password" validate:"min=6"`
		}{Password: "test"},
		expected: "Password must be a minimum of 6 in length",
	},
	{
		name: `max`,
		input: struct {
			Course string `json:"course" validate:"max=7"`
		}{Course: "CS-0001."},
		expected: "Course must be a maximum of 7 in length",
	},
	{
		name: `url`,
		input: struct {
			Image string `json:"image" validate:"url"`
		}{Image: "image.png"},
		expected: "Image must be a valid URL",
	},
	{
		name: `alpha_space`,
		input: struct {
			Name string `json:"name" validate:"alpha_space"`
		}{Name: "Some Name 2"},
		expected: "Name can only contain alphabetic and space characters",
	},
	{
		name: `pwd`,
		input: struct {
			Password string `json:"password" validate:"pwd"`
		}{Password: "password123!@#$%^&*()"},
		expected: "Password invalid",
	},
}

func TestToErrReaponse(t *testing.T) {
	vr := validator.New()

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := vr.Struct(tc.input)
			if errResp := validator.ToErrResponse(err); errResp == nil || len(errResp.Errors) != 1 {
				t.Fatalf(`Expected:"{[%v]}", Got:"%v"`, tc.expected, errResp)
			} else if errResp.Errors[0] != tc.expected {
				t.Fatalf(`Expected:"%v", Got:"%v"`, tc.expected, errResp.Errors[0])
			}
		})
	}
}
