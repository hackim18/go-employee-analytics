package migrations

import (
	"go-employee-analytics/internal/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.Employee{},
		&entity.AnnualReview{},
	)
}
