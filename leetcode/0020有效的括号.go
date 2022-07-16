package leetcode

import (
	// "fmt"
	"strings"
)

//栈 解决

/**

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。


示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false
示例 4：

输入：s = "([)]"
输出：false
示例 5：

输入：s = "{[]}"
输出：true


提示：
1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/

// func main() {
// 	//fmt.Println(stringReplace("{]}[](){}"))
// 	//fmt.Println(stringReplace("{}[](){}"))

// 	fmt.Println(stackResolve("{]}[](){}"))
// 	fmt.Println(stackResolve("{}[](){}"))
// 	//fmt.Println(stackResolve("{}[](){}"))

// 	fmt.Println(']')

// }

func stackResolve(str string) bool {
	stack := NewStack()
	n := 0
	for _, v := range str {
		if v == '{' || v == '(' || v == '[' {
			stack.data = append(stack.data, v)
			n++
		} else {
			if v == '}' {
				if n == 0 || stack.data[n-1] != '{' {
					return false
				} else {
					stack.data = stack.data[:n-1]
					n--
				}
			}
			if v == ']' {
				if n == 0 || stack.data[n-1] != '[' {
					return false
				} else {
					stack.data = stack.data[:n-1]
					n--
				}
			}
			if v == ')' {
				if n == 0 || stack.data[n-1] != '(' {
					return false
				} else {
					stack.data = stack.data[:n-1]
					n--
				}
			}
		}
	}
	return len(stack.data) == 0
}

// 字符串 方案1
func stringReplace(str string) bool {
	str = strings.Replace(str, "{}", "", -1)
	str = strings.Replace(str, "[]", "", -1)
	str = strings.Replace(str, "()", "", -1)
	return len(str) == 0
}

type Stack struct {
	data  []rune
	index int
}

func NewStack() *Stack {
	return &Stack{
		data:  make([]rune, 0),
		index: 0,
	}

}
