package main

import "strings"

func isPalindrome(s string) bool {
	i:=0
	j := len(s)-1
	for i<j {
		if !isVaild(s[i]){
			i++
			continue
		}
		if !isVaild(s[j]){
			j--
			continue
		}
		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])){
			return false
		}
		i++
		j--
	}
	return true
}

func isVaild(s byte) bool  {
	return (s >= 'A' && s <= 'Z') || (s >='0'&&s<='9')||(s>='a'&&s<='z')
}