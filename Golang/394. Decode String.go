package leetcode

// 字符串解码
// https://leetcode-cn.com/problems/decode-string/

// 基本解压：3[a]2[bc]
// 嵌套解压：3[a2[c]]
// 混合解压：2[abc]3[cd]ef
// 部分压缩：abc3[cd]xyz

// 栈
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func decodeString(s string) string {
	// 使用两个栈：数字栈和字符串栈
	countStack := []int{}
	strStack := []string{}

	currentStr := "" // 当前累积的字符串
	currentNum := 0  // 当前解析的数字

	for _, char := range s {
		if char >= '0' && char <= '9' {
			// 处理数字：累积多位数字
			currentNum = currentNum*10 + int(char-'0')
		} else if char == '[' {
			// 遇到左括号，将当前状态压入栈
			countStack = append(countStack, currentNum) // 当前数字入栈
			strStack = append(strStack, currentStr)     // 当前字符串入栈

			// 重置当前状态，准备处理括号内的内容
			currentNum = 0
			currentStr = ""
		} else if char == ']' {
			// 遇到右括号，弹出栈顶状态并处理
			// 弹出重复次数
			count := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			// 弹出之前的字符串
			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			// 将当前字符串重复count次，并拼接到之前字符串后面
			repeatedStr := ""
			for i := 0; i < count; i++ {
				repeatedStr += currentStr
			}
			currentStr = prevStr + repeatedStr
		} else {
			// 普通字符，直接添加到当前字符串
			currentStr += string(char)
		}
	}

	return currentStr
}

// 递归
// 遇到括号就递归，遇到字母就累积
func decodeStringRec(s string) string {
	// 主函数调用递归辅助函数
	ans, _ := decode(s, 0)
	return ans
}

// 递归辅助函数
// 参数：s：输入字符串，index：当前处理位置
// 返回值：解码后的字符串，下一个要处理的位置
func decode(s string, index int) (string, int) {
	res := "" // 存储当前层级的解码结果

	// 遍历字符串直到结束或遇到右括号
	for index < len(s) && s[index] != ']' {
		// 如果是字母，直接添加到结果
		if s[index] >= 'a' && s[index] <= 'z' {
			res += string(s[index])
			index++
		} else { // 遇到数字，需要处理重复模式
			// 1. 解析重复次数
			count := 0
			for index < len(s) && s[index] >= '0' && s[index] <= '9' {
				count = count*10 + int(s[index]-'0')
				index++
			}

			// 2. 跳过左括号'['
			index++

			// 3. 递归解码括号内的内容
			innerStr, nextIndex := decode(s, index)

			// 4. 重复innerStr count次
			for i := 0; i < count; i++ {
				res += innerStr
			}

			// 5. 更新索引位置
			index = nextIndex
		}
	}

	// 返回结果和下一个位置（如果是遇到']'，需要跳过它）
	if index < len(s) && s[index] == ']' {
		index++ // 跳过右括号
	}
	return res, index
}
