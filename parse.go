package main

import (
	"encoding/json"
	"fmt"

	qp "github.com/prasanthmj/qpeg/qp"
)

func main() {
	res, err := qp.Parse("", []byte(`item.ss=asus | item.spec.ram > 8 | item.spec.ssd>512 | ss=yy `))
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := json.MarshalIndent(res, "", "   ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("output:\n%v\n", string(result))
}
