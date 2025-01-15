package pkg

import (
	"fmt"
	"strconv"
	"unicode"
)

func Calc(expression string) (float64, error) {
	output := []float64{}
	operators := []rune{}
	var lastWasOperator = true // Флаг для отслеживания последнего символа

	for i := 0; i < len(expression); {
		ch := rune(expression[i])

		// Игнорируем пробелы
		if unicode.IsSpace(ch) {
			i++
			continue
		}

		// Обработка унарного минуса
		if ch == '-' && lastWasOperator {
			i++
			start := i
			for i < len(expression) && (unicode.IsDigit(rune(expression[i])) || expression[i] == '.') {
				i++
			}
			num, err := strconv.ParseFloat(expression[start-1:i], 64) // Включаем минус в число
			if err != nil {
				return 0, fmt.Errorf("некорректное число")
			}
			output = append(output, num)
			lastWasOperator = false
			continue
		}

		// Проверка на число
		if unicode.IsDigit(ch) || ch == '.' {
			start := i
			for i < len(expression) && (unicode.IsDigit(rune(expression[i])) || expression[i] == '.') {
				i++
			}
			num, err := strconv.ParseFloat(expression[start:i], 64)
			if err != nil {
				return 0, fmt.Errorf("некорректное число")
			}
			output = append(output, num)
			lastWasOperator = false
			continue
		}

		// Проверка на открывающую скобку
		if ch == '(' {
			operators = append(operators, ch)
			lastWasOperator = true
			i++
			continue
		}

		// Проверка на закрывающую скобку
		if ch == ')' {
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				result, err := applyOperator(operators[len(operators)-1], &output)
				if err != nil {
					return 0, err
				}
				output = append(output, result)
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return 0, fmt.Errorf("недостаточно открывающих скобок")
			}
			operators = operators[:len(operators)-1] // Удаление '('
			lastWasOperator = false
			i++
			continue
		}

		// Проверка на оператор
		if isOperator(ch) {
			if lastWasOperator && ch != '-' {
				return 0, fmt.Errorf("неправильное использование оператора: %c", ch)
			}
			for len(operators) > 0 && precedence(ch) <= precedence(operators[len(operators)-1]) {
				result, err := applyOperator(operators[len(operators)-1], &output)
				if err != nil {
					return 0, err
				}
				output = append(output, result)
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, ch)
			lastWasOperator = true
			i++
			continue
		}

		return 0, fmt.Errorf("недопустимый символ: %c", ch)
	}

	// Применяем оставшиеся операторы
	for len(operators) > 0 {
		result, err := applyOperator(operators[len(operators)-1], &output)
		if err != nil {
			return 0, err
		}
		output = append(output, result)
		operators = operators[:len(operators)-1]
	}

	if len(output) != 1 {
		return 0, fmt.Errorf("некорректное выражение")
	}

	return output[0], nil
}

func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

func applyOperator(op rune, output *[]float64) (float64, error) {
	if len(*output) < 2 {
		return 0, fmt.Errorf("недостаточно операндов для оператора %c", op)
	}
	b, a := (*output)[len(*output)-1], (*output)[len(*output)-2]
	*output = (*output)[:len(*output)-2]

	switch op {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	}

	return 0, fmt.Errorf("неизвестный оператор: %c", op)
}