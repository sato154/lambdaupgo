package main

import (
	"encoding/json"
	"fmt"
)

type FunctionList struct {
	Functions []Function
}

type Function struct {
	FunctionName string
}

func NewFunctionList(region string) (FunctionList, error) {
	dt, err := run("aws", "lambda", "list-functions", "--region", region)

	if err != nil {

		fmt.Println("err ", err)

		return FunctionList{}, err
	}


	fmt.Println("function list " + "3")
	var res FunctionList
	err = json.Unmarshal(dt, &res)
	return res, err
}

func (fl FunctionList) HasFunction(fname string) bool {
	for _, v := range fl.Functions {
		if v.FunctionName == fname {
			return true
		}
	}
	return false
}
