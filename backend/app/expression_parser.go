package app

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type ExpressionParser struct {
}

func NewExpressionParser() ExpressionParser {
	return ExpressionParser{}
}

func (e ExpressionParser) Calculate(expression string) (float64, error) {
	// 预处理表达式：移除空格，统一乘号
	expr := strings.ReplaceAll(expression, " ", "")
	expr = strings.ReplaceAll(expr, "×", "*")

	// 使用两个栈：操作数栈和运算符栈
	var nums []float64
	var ops []rune

	for i := 0; i < len(expr); {
		c := rune(expr[i])

		if unicode.IsDigit(c) || c == '.' {
			// 处理数字（包括小数）
			j := i
			for j < len(expr) && (unicode.IsDigit(rune(expr[j])) || expr[j] == '.') {
				j++
			}
			num, err := strconv.ParseFloat(expr[i:j], 64)
			if err != nil {
				return 0, fmt.Errorf("无效数字: %s", expr[i:j])
			}
			nums = append(nums, num)
			i = j
		} else if c == '(' {
			ops = append(ops, c)
			i++
		} else if c == ')' {
			// 计算到匹配的左括号
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				if err := applyOp(&nums, &ops); err != nil {
					return 0, err
				}
			}
			if len(ops) == 0 {
				return 0, fmt.Errorf("括号不匹配")
			}
			ops = ops[:len(ops)-1] // 弹出左括号
			i++
		} else if isOperator(c) {
			// 处理运算符优先级
			for len(ops) > 0 && ops[len(ops)-1] != '(' &&
				precedence(ops[len(ops)-1]) >= precedence(c) {
				if err := applyOp(&nums, &ops); err != nil {
					return 0, err
				}
			}
			ops = append(ops, c)
			i++
		} else {
			return 0, fmt.Errorf("无效字符: %c", c)
		}
	}

	// 处理剩余的运算符
	for len(ops) > 0 {
		if ops[len(ops)-1] == '(' {
			return 0, fmt.Errorf("括号不匹配")
		}
		if err := applyOp(&nums, &ops); err != nil {
			return 0, err
		}
	}

	if len(nums) != 1 {
		return 0, fmt.Errorf("表达式不完整")
	}

	return nums[0], nil
}

func isOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
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

func applyOp(nums *[]float64, ops *[]rune) error {
	if len(*nums) < 2 || len(*ops) < 1 {
		return fmt.Errorf("表达式不完整")
	}

	op := (*ops)[len(*ops)-1]
	*ops = (*ops)[:len(*ops)-1]

	b := (*nums)[len(*nums)-1]
	a := (*nums)[len(*nums)-2]
	*nums = (*nums)[:len(*nums)-2]

	var result float64
	switch op {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		if b == 0 {
			return fmt.Errorf("除零错误")
		}
		result = a / b
	default:
		return fmt.Errorf("未知运算符: %c", op)
	}

	*nums = append(*nums, result)
	return nil
}
