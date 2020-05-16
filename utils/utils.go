package utils

import (
	"errors"
	"strconv"
)

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
