package main

import (
	"github.com/pkg/errors"
)

func test() error {
	var err error
	if err != nil { return errors.WithMessage(err, "") }
	return nil
}
