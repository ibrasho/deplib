package deplib

import "github.com/pkg/errors"

func Test() error {
	return errors.New("test lib err")
}
