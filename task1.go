package main

import (
	"fmt"
)

// Automor number - natural number in a given number base {\displaystyle b}b whose square "ends" in the same digits as the number itself.
func isAutomor(num int) bool {
	var quar = num * num
	var res = true
	
	for {
	  // compare the last digit
	  res = num%10 == quar%10
    	  num /= 10
	  if (!res || num == 0) {
        	break
    	  }

	  quar /= 10
	}
	
	return res;
}

func main() {
	// we will display all Automor number between 1 and a
        var a int = 1000000

	for i := 1; i <= a; i++ {
	  if isAutomor(i) {
	    fmt.Println(i);
	  }
	}
}