package main

import (
	"math"
	"strconv"
)

func addBinary(a string, b string) string {
	if len(a) < 1 || len(b) > int(math.Pow10(4)) {
		return ""
	}

	dalt := 0
	r := ""
	i,j := len(a)-1,len(b)-1
	for i>=0||j>=0{
		xa,xb:=0,0
		if i>=0{
			xa, _ = strconv.Atoi(string(a[i]))
		}
		if j>=0{
			xb, _ = strconv.Atoi(string(b[j]))
		}
		switch xa+xb+dalt {
		case 3:
			r = "1"+r
			dalt = 1
		case 2:
			r = "0"+r
			dalt = 1
		case 1:
			r = "1"+r
			dalt = 0
		case 0:
			r = "0"+r
			dalt = 0
		}
		i--
		j--
	}
	if dalt == 1 {
		return "1"+r
	}else{
		return r
	}
}
