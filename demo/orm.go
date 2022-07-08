package web

import (
	"errors"
)

var (
	NoResult = errors.New("no result")
	NoField  = errors.New("no field")
)

type Model map[string]interface{}

// Do 模拟依赖库，通常自定义异常
func Do(id int) (Model, error) {
	switch id {
	case 1:
		return Model{"blog": "https://darr-en1.top"}, nil
	case 2:
		return Model{}, NoField
	default:
		return Model{}, NoResult
	}
}
