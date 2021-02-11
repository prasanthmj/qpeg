package main

import (
	"encoding/json"
	"fmt"

	qp "github.com/prasanthmj/qpeg/qp"
)

func main() {
	res, err := qp.Parse("", []byte(`item.spec.ssd=512gb`))
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
