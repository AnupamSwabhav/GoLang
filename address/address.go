package address

import (
	"preloading/test/model"

	uuid "github.com/satori/go.uuid"
)

type Address struct {
	model.Model
	StudentID uuid.UUID `sql:"REFERENCES customers(id)" type:"uuid" json:"studentID" gorm:"type:VARCHAR(36)"`
	City      string    `json:"city"`
	State     *string   `json:"state"`
}

// func New(studentid uuid.UUID, city string, state *string) *Address {
// 	return &Address{
// 		StudentID: studentid,
// 		City:      city,
// 		State:     state,
// 	}
// }
