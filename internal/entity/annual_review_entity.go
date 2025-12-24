package entity

import "time"

type AnnualReview struct {
	ID         int       `gorm:"column:id;primaryKey" json:"id"`
	EmpID      int       `gorm:"column:emp_id" json:"emp_id"`
	ReviewDate time.Time `gorm:"column:review_date" json:"review_date"`
	// Employee   *Employee `gorm:"foreignKey:EmpID;references:ID" json:"employee,omitempty"`
}

func (AnnualReview) TableName() string {
	return "annual_reviews"
}
