package service

import (
	"demo/first/Yd/pkg"
	"errors"
	"strings"
)

func CalculateExpression(expr string) (float64, error) {
	expr = strings.TrimSpace(expr)

	for _, char := range expr {
		if !strings.ContainsRune("0123456789+-*/. ()", char) {
			return 0, errors.New("Expression is not valid")
		}
	}

	result, err := pkg.Calc(expr)
	if err != nil {
		return 0, err
	}

	return result, nil
}