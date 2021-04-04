package main

import (
	"fmt"
	"strings"
)
var pos string = "AZaz";

func encode(str string) string {
	var b strings.Builder
	for i := 0; i < len(str); i++{
	  if (str[i] >= pos[0] && str[i] < pos[1]) ||
	     (str[i] >= pos[2] && str[i] < pos[3]){
	    b.WriteString(string(str[i] + 1))
	  } else if (str[i] == pos[1]){
	    b.WriteString(string(pos[0]))
	  } else if (str[i] == pos[3]){
	    b.WriteString(string(pos[2]))
	  } else {
  	    b.WriteString(string(str[i]))
	  }
	}

	return b.String()
}

func decode(str string) string {
	var b strings.Builder
	for i := 0; i < len(str); i++{
	  if (str[i] > pos[0] && str[i] <= pos[1]) ||
	     (str[i] > pos[2] && str[i] <= pos[3]){
	    b.WriteString(string(str[i] - 1))
	  } else if (str[i] == pos[0]){
	    b.WriteString(string(pos[1]))
	  } else if (str[i] == pos[2]){
	    b.WriteString(string(pos[3]))
	  } else {
  	    b.WriteString(string(str[i]))
	  }
	}

	return b.String()
}

func main() {
	str := "Decoding string by Sezar method!";
	fmt.Println(encode(str));
	fmt.Println(decode(encode(str)));
}
