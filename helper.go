package main

import "errors"

func RefString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func RefInt(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

func DerefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func DerefInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

var (
	ErrInvalidRegistrationNumber = errors.New("invalid number plate")
	ErrCarNotFound               = errors.New("car not found")
)
