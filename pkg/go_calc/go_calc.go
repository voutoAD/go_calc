package go_calc

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func validate(expression string) (bool, error) {
	if len(expression) == 0 {
		return false, fmt.Errorf("length of zero")
	}
	count := 0
	for k, v := range expression {
		if string(v) == "(" {
			count++
		} else if string(v) == ")" {
			count--
		}
		if k > 0 && strings.Contains("+-*/", string(expression[k])) && strings.Contains("+-*/", string(expression[k-1])) {
			return false, fmt.Errorf("two symbols")
		}
		if v == '0' && expression[k-1] == '/' {
			return false, fmt.Errorf("divide by zero")
		}
	}
	if count != 0 {
		return false, fmt.Errorf("brackets")
	}
	if isOperand(rune(expression[len(expression)-1])) {
		return false, fmt.Errorf("ww")
	}
	return true, nil
}

func infixToPostfix(expression string) []rune {
	precedence := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}
	var stack []rune
	var queue []rune
	for _, v := range expression {
		if isNum(v) {
			queue = append(queue, v)
		} else if isOperand(v) {
			n := precedence[v]
			if n == 1 {
				if len(stack) == 0 || stack[len(stack)-1] == '(' {
					stack = append(stack, v)
				} else if stack[len(stack)-1] == '*' || stack[len(stack)-1] == '/' {
					stack, queue = pop(stack, queue)
					stack = append(stack, v)
				}
			} else {
				if len(stack) != 0 && (stack[len(stack)-1] == '*' || stack[len(stack)-1] == '/') {
					stack, queue = pop(stack, queue)
				}
				stack = append(stack, v)
			}
		} else if v == '(' {
			stack = append(stack, v)
		} else if v == ')' {
			stack, queue = pop(stack, queue)
		}
	}
	if len(stack) != 0 {
		for i := len(stack) - 1; i >= 0; i-- {
			queue = append(queue, stack[i])
		}
	}
	return queue
}

func isOperand(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func isNum(char rune) bool {
	return char >= '0' && char <= '9'
}

func pop(stack []rune, queue []rune) ([]rune, []rune) {
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] == '(' {
			stack = stack[:i]
			break
		}
		queue = append(queue, stack[i])
	}
	return stack, queue
}

func Calc(expression string) (float64, error) {
	res, err := validate(expression)
	if !res {
		return 0.0, err
	}
	postfix := infixToPostfix(expression)
	var result []interface{}
	for _, v := range postfix {
		if isNum(v) {
			n, _ := strconv.ParseFloat(string(v), 64)
			result = append(result, n)
		} else {
			result = append(result, string(v))
		}
	}
	var stack []float64
	for _, r := range result {
		if reflect.TypeOf(r) == reflect.TypeOf(0.0) {
			x := r.(float64)
			stack = append(stack, x)
		} else {
			switch r {
			case "+":
				stack = append(stack[:len(stack)-2], stack[len(stack)-2]+stack[len(stack)-1])
			case "-":
				stack = append(stack[:len(stack)-2], stack[len(stack)-2]-stack[len(stack)-1])
			case "*":
				stack = append(stack[:len(stack)-2], stack[len(stack)-2]*stack[len(stack)-1])
			case "/":
				stack = append(stack[:len(stack)-2], stack[len(stack)-2]/stack[len(stack)-1])
			}
		}
	}
	return stack[0], nil
}
