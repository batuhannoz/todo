package model

import (
	"time"
)

func (Todo) TableName() string {
	return "todo"
}

type Todo struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Task       string    `gorm:"column:task;type:varchar(150)" json:"task"`
	CreateDate time.Time `gorm:"column:create_date;type:datetime" json:"create_date"`
}
