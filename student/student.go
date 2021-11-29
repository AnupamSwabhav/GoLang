package student

import (
	"preloading/test/address"
	"preloading/test/model"
)

type Student struct {
	model.Model
	FName   string            `json:"fname"`
	LName   *string           `json:"lname"`
	Age     int               `json:"age"`
	IsMale  *bool             `json:"ismale"`
	Address []address.Address `json:"address"`
}

func New(cFName string, cLName *string, cAge int, cIsMale *bool) *Student {
	return &Student{
		FName:  cFName,
		LName:  cLName,
		Age:    cAge,
		IsMale: cIsMale,
	}
}
