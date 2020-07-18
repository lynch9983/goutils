package goutils

import (
	"fmt"
	"testing"
)

func Test_InterfaceConvert(t *testing.T) {
	var val interface{}

	val = "qqwwee"
	if res, err := Interface2String(val); err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}

	val = 1.23
	if res, err := Interface2JsonNum(val); err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}

	val = 456
	if res, err := Interface2Int(val); err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}

	val = 7.89
	if res, err := Interface2Float64(val); err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}

	val = true
	if res, err := Interface2Bool(val); err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}
