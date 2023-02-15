// Filename: main.go
// "waitgroups" are just another way to syncronize your goroutines
package main

//demonstrate flags
import (
"flag"
"fmt"
"strings"
)


func main() {


//set the flags
msg := flag.String("msg", "Howdy, stranger!", "the message to display")
num := flag.Int("num",1 , "how many time to display the message")
caps := flag.Bool("U", false,"specify whether to uppercase the string(true or false)")
reverse := flag.Bool("r", false, "reverse the string(true or false)")
//caps := flag.Bool("caps", false,"should the string print in caps")
flag.Parse()
//check if it is upper case or if it should be uppercase before printing it.
if *caps {
*msg = strings.ToUpper(*msg)
}
//check if we should reverse the string
if *reverse {
	//reverse string
	var result string
	for _, value := range *msg {
result = string(value) + result
	}
	//write the reverse string to *msg
	*msg = result;
}
//print the s string


//check if the user set the flag
//fmt.Println(*msg)
//fmt.Println(*num)
for i := 0; i < *num; i++ {
fmt.Println(*msg)
}
}
