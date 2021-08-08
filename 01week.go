// 20. 有效的括号
// https://leetcode-cn.com/problems/valid-parentheses/

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


// 70. 爬楼梯
// https://leetcode-cn.com/problems/climbing-stairs/

func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    f1, f2 := 1, 2
    for i := 3; i <= n; i++ {
        f3 := f2 + f1
        f1 = f2
        f2 = f3
    }
    return f2
}

// 11. 盛最多水的容器
// https://leetcode-cn.com/problems/container-with-most-water/
func maxArea(height []int) int {
    max := 0
    for i, j := 0, len(height) - 1; i < j; {
        minHeight := height[i] 
        if minHeight > height[j] {
            minHeight = height[j]
            j = j - 1
        } else {
            i = i + 1
        }
        area := (j-i+1)*minHeight
        if max < area {
            max = area
        }
    }
    return max
}
