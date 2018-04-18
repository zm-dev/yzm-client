package distinguish

import (
	"errors"
	"strings"
	"strconv"
	"container/list"
)

type Label2StrFunc func(label Label) string

func GetLabel2StrFunc(category int) (Label2StrFunc, error) {
	switch category {
	case 0:
		return label2StrFuncCategory0, nil
	case 1, 2:
		return label2StrFuncCategory1and2, nil
	case 3:
		return label2StrFuncCategory3, nil
	case 4:
		return label2StrFuncCategory4, nil
	default:
		return nil, errors.New("nonsupport category")
	}
}

func label2StrFuncCategory0(label Label) string {
	level := map[rune]int{
		'*': 2,
		'-': 1,
		'+': 1,
	}
	nums, ops := list.New(), list.New()

	isOp := true

	for _, c := range label.Yzm {
		if c >= '0' && c <= '9' {
			if isOp {
				nums.PushBack(int(c - '0'))
			} else {
				// 前一个字符也是数字
				if num, ok := pop(nums).(int); ok {
					nums.PushBack(num*10 + int(c-'0'))
				}
			}
			isOp = false
		} else {
			isOp = true
			if ops.Len() > 0 {
				op := ops.Back().Value.(rune)
				if needCalc(level, op, c) {
					binaryCalc(ops, nums)
				}
			}
			ops.PushBack(c)
		}
	}

	for ops.Len() > 0 {
		binaryCalc(ops, nums)
	}

	e := nums.Front()
	yzm := label.Yzm

	if e != nil {
		res := e.Value.(int)
		yzm += "=" + strconv.Itoa(res)

	}

	if label.ImageFilename == "" {
		return yzm
	}
	return label.ImageFilename + "," + yzm + "\n"
}

func pop(l *list.List) interface{} {
	e := l.Back()
	if e != nil {
		l.Remove(e)
		return e.Value
	}
	return nil
}

func needCalc(level map[rune]int, opLeft, op2Right rune) bool {
	return level[opLeft]-level[op2Right] >= 0
}

func binaryCalc(ops, nums *list.List) {

	op, ok := pop(ops).(rune)
	if !ok {
		return
	}

	num2, ok := pop(nums).(int)
	if !ok {
		return
	}

	num1, ok := pop(nums).(int)
	if !ok {
		return
	}

	res := 0

	switch op {
	case '+':
		res = num1 + num2
	case '-':
		res = num1 - num2
	case '*':
		res = num1 * num2
	}
	nums.PushBack(res)
}

func label2StrFuncCategory1and2(label Label) string {
	if label.ImageFilename == "" {
		return strings.ToUpper(label.Yzm)
	}
	return label.ImageFilename + "," + strings.ToUpper(label.Yzm) + "\n"
}

func label2StrFuncCategory3(label Label) string {
	yzm := strconv.Itoa(strings.IndexRune(label.Yzm, '1'))
	if label.ImageFilename == "" {
		return yzm
	}
	return label.ImageFilename + "," + yzm + "\n"
}

func label2StrFuncCategory4(label Label) string {
	// todo
	return ""
}
