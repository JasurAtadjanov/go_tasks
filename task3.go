package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := "Are you from UK? No, I am from Uzb. It is in Middle Asia. You should know about Uzb";
	if (strings.Trim(str, " ") == ""){
		return;
	}
	str = strings.ToLower(str)
	b := regexp.MustCompile(`,|( )|-|!|\.|:|\?`)
	items := b.Split(str, -1)
	mItems := map[string]int{}
	word := ""
	for i := 0; i < len(items); i++{
	        word = strings.Trim(items[i], " ")
		if len(word) > 0 {
		  mItems[word]++  
		}
	}
	fmt.Println(mItems)
}
