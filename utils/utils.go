package utils

import (
	"errors"
	"strconv"
)

type Paginator struct {
	Page int `form:"page" binding:"numeric,gt=0"`
	Limit int `form:"limit" binding:"numeric,gt=0"`
}

func GetInt(param interface{}) (int, error) {
	switch v := param.(type) {
	case int:
		return v, nil
	case string:
		if i, err := strconv.Atoi(v); err != nil || i <= 0 {
			return 0, errors.New("invalid id")
		} else {
			return i, nil
		}
	default:
		return 0, errors.New("invalid id")
	}
}
