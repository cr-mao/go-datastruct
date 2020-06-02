package demo

import (
	"datastruct/stack"
)

// 栈模拟递归
func StackMockRecurse(first int) int {
	mystatck := stack.NewStack()
	mystatck.Push(first)
	var res int
	for !mystatck.IsEmpty() {
		val := mystatck.Pop()
		if val == 0 {
			res += 0
		} else {
			res += val.(int)
			mystatck.Push((val.(int) - 1))
		}
	}
	return res
}

//斐波那契数列 栈模拟递归
// 1,1, 2,3,5,8,13
func Feibo(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return Feibo(num-1) + Feibo(num-2)
	}
}

func StackFeibo(num int) int {
	mystatck := stack.NewStack()
	mystatck.Push(num)
	var res int
	for !mystatck.IsEmpty() {
		num = mystatck.Pop().(int)
		if num == 1 || num == 2 {
			res += 1
		} else {
			mystatck.Push((num - 1))
			mystatck.Push((num - 2))
		}
	}
	return res
}


