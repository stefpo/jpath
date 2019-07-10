package main

import (
	"encoding/json"
	"fmt"

	"github.com/stefpo/jpath"
)

type testData struct {
	Xm1 string `json:"xm1"`
	Xm2 string `json:"Xm2"`
}

func main() {
	//fmt.Println("Dynamic JSON test")
	js := `{
		"s1":"abcde",
		"i1":3,
		"f1":17.3, 
		"a1":[4,5,6, { "x1":"vx1", "x2": "vx2"}	],
		"m1":{ "xm1":"vx1", "xm2": "vx2"},
		"d":"2012-04-23T18:25:43.511Z"
	}`

	obj, _ := jpath.Parse(js)
	str, _ := jpath.Stringify(obj)

	fmt.Println(js)
	fmt.Println(obj)
	fmt.Println(str)
	fmt.Println(jpath.Get(obj, "/m1/xm1"))
	fmt.Println(jpath.Get(obj, "/m1/xm2"))
	fmt.Println(jpath.Get(obj, "/m1"))
	fmt.Println(jpath.Get(obj, "/m1/xm1/toto"))
	fmt.Println(obj)
	jpath.Set(obj, "/m1/xm1", "updated value")
	fmt.Println(obj)
	jpath.Set(obj, "/am/am1", "added value")
	fmt.Println(jpath.Stringify(obj))

	td := testData{}
	m1, _ := jpath.Get(obj, "/m1")
	str2, _ := jpath.Stringify(m1)
	json.Unmarshal([]byte(str2), &td)
	fmt.Println(str2)
	fmt.Println(jpath.Stringify(td))
}
