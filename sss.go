package main

import (
	"github.com/pkg/errors"
)

func test() error {
	var err error

	if err != nil {
		if err != nil {
			if err != nil { return errors.WithMessage(err, "") }
			return errors.WithMessage(err, "")
		}
		if err != nil { return errors.WithMessage(err, "") }
	}
	if err != nil {
		return errors.WithMessage(err, "") //
	}
	return nil
}
