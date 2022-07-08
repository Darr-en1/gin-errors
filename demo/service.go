package web

import (
	"github.com/Darr-en1/gin-errors/errors"
)

func getBlogService(id int) (Model, error) {
	result, err := getBlog(id)
	if err != nil {
		switch err {
		case NoField:
			err = errors.MySQLNoFieldError.Wrap(err)
		case NoResult:
			err = errors.MySQLNoQueryError.Wrap(err)
		}

		return nil, err
	}
	return result, nil
}
