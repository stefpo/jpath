package main


import (
	"fmt"
	"github.com/stefpo/jpath"
)

func main() {
	js:=`{"id":"12345","method":"eq2","params":{"A":"1","B":"2","C":"-3"}}
	`
	obj, _ := jpath.Parse(js)
	
	fmt.Println(jpath.Get(obj,"id"))
	fmt.Println(jpath.GetString2(obj,"id","xx"))
}
