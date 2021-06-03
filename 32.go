package main


// TODO 32
func longestValidParentheses(s string) int {
	stack := make([]rune, 0)
	ret := 0
	i:=0
	for _, v := range s{
		if v == '('{
			if len(stack) >0{
				i=0
				continue
			}else{
				stack = append(stack, v)
			}
		}else{
			if len(stack) > 0{
				stack = stack[:len(stack)-1]
				i=i +2
			}else{
				i=0
				continue
			}
		}
		ret = maxS(ret, i)
	}
	return ret
}

