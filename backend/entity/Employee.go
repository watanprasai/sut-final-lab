package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid:"required~Please enter your name"`
	Email      string
	EmployeeID string `valid:"matches(^[JMS][0-9]{8}$)~EmployeeID invalid format"`
}
