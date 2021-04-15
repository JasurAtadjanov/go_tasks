package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func decode(str string) (res string) {
	rx := regexp.MustCompile(`\#\d+\#`)
	fmt.Println(str, ":", rx.MatchString(str))
	nums := rx.FindAllString(str, -1)

	if len(nums) < 1 {
		return str
	}

	items := rx.Split(str, -1)
	res = items[0]
	for i := 0; i < len(nums); i++ {
		n, _ := strconv.ParseInt(strings.Replace(nums[i], "#", "", -1), 10, 0)
		ch := string(items[i+1][0])
		for j := 0; j < int(n); j++ {
			res += ch
		}
		res += items[i+1]
	}

	return
}

func encode(str string) (string, error) {
	rx := regexp.MustCompile(`\#\d+\#`)
	if rx.MatchString(str) {
		return str, errors.New("This string can not be Compressed by this Algorithm")
	}

	n := len(str)
	p := 0
	cnt := 1
	ch := string(str[p])
	res := ""

	for i := 1; i < n; i++ {
		if ch == string(str[i]) {
			cnt++
			continue
		}

		if cnt > 4 {
			res += "#" + strconv.Itoa(cnt) + "#" + ch
		} else {
			res += str[p:i]
		}

		ch = string(str[i])
		cnt = 1
		p = i
	}

	if cnt > 4 {
		res += "#" + strconv.Itoa(cnt) + "#" + ch
	} else if p < n {
		res += str[p:n]
	}
	return res, nil
}

func main() {
	var str string = "as1111112assssssssss"
	var s, er = encode(str)
	if er == nil {
		fmt.Println("Encoding: ", str, " is ", s)
		fmt.Println("Decoding: ", s, " is ", decode(s))
	}

}
