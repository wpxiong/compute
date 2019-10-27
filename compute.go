// test project main.go
package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	if result, err := computeStr("1233+3*(4*(6-(2*(3-2))))+(9-3)*(4-2)"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("result is :")
		fmt.Println(result)
	}

}

func compute(input string) (int, error) {
	var inputOperatorStack []string
	var inputDataStack []int
	var temStr string = ""
	var isOper bool = false
	for index, str := range input {
		if str == '+' || str == '-' || str == '*' || str == '/' {
			if temInt, err := strconv.Atoi(temStr); err != nil {
				return -1, errors.New("input string is not illegal")
			} else {
				if isOper {
					var value int = inputDataStack[len(inputDataStack)-1]
					inputDataStack = inputDataStack[0 : len(inputDataStack)-1]
					var mark string = inputOperatorStack[len(inputOperatorStack)-1]
					inputOperatorStack = inputOperatorStack[0 : len(inputOperatorStack)-1]
					temInt = computeTwoEval(mark, value, temInt)
				}
				inputDataStack = append(inputDataStack, temInt)
				inputOperatorStack = append(inputOperatorStack, string(str))
				temStr = ""
				isOper = false
				if str == '*' || str == '/' {
					isOper = true
				}
			}
		} else {
			temStr = temStr + string(str)
			if index == len(input)-1 {
				if temInt, err := strconv.Atoi(temStr); err != nil {
					return -1, errors.New("input string is not illegal")
				} else {
					if isOper {
						var value int = inputDataStack[len(inputDataStack)-1]
						inputDataStack = inputDataStack[0 : len(inputDataStack)-1]
						var mark string = inputOperatorStack[len(inputOperatorStack)-1]
						inputOperatorStack = inputOperatorStack[0 : len(inputOperatorStack)-1]
						var result = computeTwoEval(mark, value, temInt)
						temInt = result
					}
					inputDataStack = append(inputDataStack, temInt)
				}
			}
		}
	}

	inputDataStack = reverseArrayInt(inputDataStack)
	inputOperatorStack = reverseArrayString(inputOperatorStack)
	for {
		if len(inputDataStack) > 1 {
			var v1 int = inputDataStack[len(inputDataStack)-1]
			var v2 int = inputDataStack[len(inputDataStack)-2]
			inputDataStack = inputDataStack[0 : len(inputDataStack)-2]
			var markNew string = inputOperatorStack[len(inputOperatorStack)-1]
			inputOperatorStack = inputOperatorStack[0 : len(inputOperatorStack)-1]
			inputDataStack = append(inputDataStack, computeTwoEval(markNew, v1, v2))
		} else {
			break
		}
	}
	return inputDataStack[0], nil
}

func reverseArrayInt(l []int) []int {
	for i := 0; i < int(len(l)/2); i++ {
		li := len(l) - i - 1
		l[i], l[li] = l[li], l[i]
	}
	return l
}

func reverseArrayString(l []string) []string {
	for i := 0; i < int(len(l)/2); i++ {
		li := len(l) - i - 1
		l[i], l[li] = l[li], l[i]
	}
	return l
}

func computeTwoEval(mark string, value1 int, value2 int) int {
	if mark == "+" {
		return value1 + value2
	} else if mark == "-" {
		return value1 - value2
	} else if mark == "*" {
		return value1 * value2
	} else if mark == "/" {
		return value1 / value2
	}
	return 0
}

func findLastQuto(input string, start int) (int, error) {
	var end, i int = 1, 0
	for i, str := range input[start:] {
		if str == '(' {
			end += 1
		} else if str == ')' {
			end -= 1
		}
		if end == 0 {
			return i, nil
		}
	}
	if end == 0 {
		return i, nil
	} else {
		return -1, errors.New("input string is illegal")
	}
}

func computeStr(inputStr string) (int, error) {
	inputStr = strings.ReplaceAll(inputStr, " ", "")
	if len(inputStr) == 0 {
		return 0, errors.New("input string is illegal")
	}
	r := regexp.MustCompile(`^[0-9\+\-\*\/\(\)]+$`)
	if !r.MatchString(inputStr) {
		return 0, errors.New("input string is illegal")
	}
	return computeQuoto(inputStr)
}

func computeQuoto(input string) (int, error) {
	var index int = strings.Index(input, "(")
	if index >= 0 {
		if quotoIndex, err := findLastQuto(input, index+1); err != nil {
			return 0, err
		} else {
			var laststr string = ""
			if quotoIndex+index+2 != len(input) {
				laststr = input[quotoIndex+index+2:]
			}
			fmt.Println(input[index+1 : quotoIndex+index+1])
			if resint, err := computeQuoto(input[index+1 : quotoIndex+index+1]); err == nil {
				return computeQuoto(input[0:index] + strconv.Itoa(resint) + laststr)
			} else {
				return -1, errors.New("input string is not illegal")
			}
		}
	} else {
		return compute(input)
	}
}

