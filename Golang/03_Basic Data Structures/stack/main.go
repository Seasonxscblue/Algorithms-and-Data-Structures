package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"cn.edu.xtu/Stack"
)

func main() {
	fmt.Println(parChecker("{()[]{}[]"))
	fmt.Println(baseConverter(258, 8))
	s := infix2Postfix(" ( 1 + 2 ) * ( 3 + 4 ) ")
	fmt.Println(postfixEval(s))
}

func parChecker(symbolString string) bool {
	m := map[byte]byte{}
	m['('] = ')'
	m['['] = ']'
	m['{'] = '}'

	s := Stack.Stack[byte]{}
	bytes := []byte(symbolString)

	for _, c := range bytes {
		if _, ok := m[c]; ok {
			s.Push(c)
		} else {
			if s.IsEmpty() {
				return false
			}

			if c != m[s.Pop()] {
				return false
			}
		}
	}

	return s.IsEmpty()
}

func baseConverter(decNumber, base int) string {
	digits := "0123456789ABCDEF"

	remStack := Stack.Stack[int]{}

	for decNumber > 0 {
		rem := decNumber % base
		remStack.Push(rem)
		decNumber /= base
	}

	rslt := ""
	for !remStack.IsEmpty() {
		rslt += string(digits[remStack.Pop()])
	}

	return rslt
}

func infix2Postfix(infixExpr string) string {
	prec := map[string]int{
		"*": 3,
		"/": 3,
		"+": 2,
		"-": 2,
		"(": 1,
		")": 1,
	}

	postfixList := []string{}
	opStack := Stack.Stack[string]{}

	infixExprArray := strings.Split(strings.Trim(infixExpr, " "), " ")
	for _, token := range infixExprArray {
		if _, ok := prec[token]; ok {
			if token == "(" {
				opStack.Push(token)
			} else if token == ")" {
				topToken := opStack.Pop() // 栈顶元素
				for topToken != "(" {
					postfixList = append(postfixList, topToken)
					topToken = opStack.Pop()
				}
			} else {
				for (!opStack.IsEmpty()) && (prec[opStack.Peek()] >= prec[token]) {
					postfixList = append(postfixList, opStack.Pop())
				}
				opStack.Push(token)
			}
		} else {
			postfixList = append(postfixList, token)
		}
	}

	rslt := strings.Join(postfixList, " ")
	for !opStack.IsEmpty() {
		rslt += " "
		rslt += opStack.Pop()
	}
	return rslt
}

func postfixEval(postfixExpr string) int {
	operandStack := Stack.Stack[int]{}

	postfixExprArray := strings.Split(postfixExpr, " ")
	for _, token := range postfixExprArray {
		if s, err := strconv.Atoi(token); err == nil {
			operandStack.Push(s)
		} else {
			b := operandStack.Pop()
			a := operandStack.Pop()
			if tmp, err := doMath(token, a, b); err == nil {
				operandStack.Push(tmp)
			}
		}
	}
	return operandStack.Pop()
}

func doMath(operator string, a, b int) (int, error) {
	if operator == "*" {
		return a * b, nil
	} else if operator == "/" {
		return a / b, nil
	} else if operator == "+" {
		return a + b, nil
	} else if operator == "-" {
		return a - b, nil
	}

	return -1, errors.New("the operator does not exist")
}
