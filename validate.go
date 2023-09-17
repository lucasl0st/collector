package main

import (
	"errors"
	"regexp"
)

const validateEntityRegexp = "^[a-zA-Z0-9]*$"

func ValidateEntity(entity string) error {
	exp, err := regexp.Compile(validateEntityRegexp)
	if err != nil {
		return err
	}

	if !exp.MatchString(entity) {
		return errors.New("contains forbidden characters")
	}

	return nil
}
