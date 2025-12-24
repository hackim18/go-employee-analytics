package entity

import "time"

type Employee struct {
	ID              int            `gorm:"column:id;primaryKey" json:"id"`
	FirstName       string         `gorm:"column:first_name" json:"first_name"`
	LastName        string         `gorm:"column:last_name" json:"last_name"`
	HireDate        time.Time      `gorm:"column:hire_date" json:"hire_date"`
	TerminationDate *time.Time     `gorm:"column:termination_date" json:"termination_date"`
	Salary          int            `gorm:"column:salary" json:"salary"`
	// AnnualReviews   []AnnualReview `gorm:"foreignKey:EmpID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"annual_reviews,omitempty"`
}

func (Employee) TableName() string {
	return "employees"
}
