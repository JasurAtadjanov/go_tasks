package main

import (
	"fmt"
	"strings"
	"math/rand"
)
var pos string = "AZaz";
var deltaUpper byte = pos[1] - pos[0]
var deltaLower byte = pos[3] - pos[2]

func encode(str string, step byte) string {
	var b strings.Builder
	var ch byte;
	for i := 0; i < len(str); i++{
	  ch = str[i] + step;
	  // if uppercase
	  if (str[i] >= pos[0] && str[i] <= pos[1]){
	    if ch > pos[1] {
 	       b.WriteString(string(ch - deltaUpper))
	    }else if ch < pos[0] {
 	       b.WriteString(string(ch + deltaUpper))
	    } else {
 	       b.WriteString(string(ch))
	    }
	  // if lower case  
  	  } else if (str[i] >= pos[2] && str[i] <= pos[3]){
	    if ch > pos[3] {
 	       b.WriteString(string(ch - deltaLower))
	    } else if ch < pos[2] {
 	       b.WriteString(string(ch + deltaLower))
	    } else {
 	       b.WriteString(string(ch))
	    }	  
	  
	  } else {
  	    b.WriteString(string(str[i]))
	  }
	}

	return b.String()
}

func decode(str string, step byte) string {
	return encode(str, 0 - step);
}

func main() {
        var key byte = byte(rand.Intn(1000)%26)
  	str := "The default number generator is deterministic, so it will produce the same sequence of numbers each time by default. To produce varying sequences, give it a seed that changes. Note that this is not safe to use for random numbers you intend to be secret, use crypto/rand for those.";
	
	test := [5]string{"is", "so", "Will", "Same", "Numbers"}
	// make test strings lower case
	vTest := ",";
	for i := 0; i < 5; i++{
	  vTest += strings.ToLower(test[i]) + ",";
	}

	dStr := encode(str, key);
	s := strings.Split(dStr, " ");

	sTmp := ""
	var j byte
	for j = 0; j < 27; j++ {
	  testCnt := 5;
	  for i := 0; i < len(s); i++{
  	        sTmp = "," + decode(s[i], j)+ ",";
		if strings.Contains(vTest, sTmp){
			testCnt--;
			if testCnt < 1 {
			 fmt.Println("Key=", j);
			 fmt.Println("Decoded=", decode(dStr, j)); 
			 return;
			}
		}
	  }
	}
}
