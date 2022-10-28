package main

import (
	"math"

	"github.com/pkg/errors"
)

func Test() error {
	var err error = errors.Errorf("")
	if math.Max(1, 2) == 1 {
		err = errors.New("")
	}
	if err != nil { return err }
	err = errors.Errorf("")
	if err != nil {
		err = errors.Errorf("")
		if err != nil {
			err = errors.Errorf("")
			if err != nil {
				return errors.Wrap(err, "")
			}
			return errors.WithMessage(err, "")
		}
		err = errors.Errorf("")
		if err != nil {
			return errors.WithMessage(err, "")
		}
	}
	if err != nil {
		return errors.Wrap(err, "") //
	}
	a, b := 1, 2
	if a/b+b-a > 1 { return nil }

	if a > 0 || b > 0 {
		return nil
	}
	return nil
}
