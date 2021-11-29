package model

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Model struct {
	ID        uuid.UUID  `sql:"primary_key" json:"id" gorm:"type:VARCHAR(36)" type:"uuid" default:"uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid.String())
	return nil
}
