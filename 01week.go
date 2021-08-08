20. 有效的括号
https://leetcode-cn.com/problems/valid-parentheses/

func isValid(s string) bool {
    n := len(s)
    if n % 2 == 1 {
        return false
    }
    pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    stack := []byte{}
    for i := 0; i < n; i++ {
        if leftChar, ok := pairs[s[i]]; ok {
            if len(stack) == 0 || stack[len(stack)-1] != leftChar {
                return false
            }
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, s[i])
        }
    }
    return len(stack) == 0
}
