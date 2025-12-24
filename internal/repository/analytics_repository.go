package repository

import (
	"context"
	"go-employee-analytics/internal/model"
	"time"

	"gorm.io/gorm"
)

type AnalyticsRepository struct {
	DB *gorm.DB
}

func NewAnalyticsRepository(db *gorm.DB) *AnalyticsRepository {
	return &AnalyticsRepository{DB: db}
}

func (r *AnalyticsRepository) ActiveSmithEmployees(ctx context.Context) ([]model.EmployeeName, error) {
	results := make([]model.EmployeeName, 0)
	err := r.DB.WithContext(ctx).Raw(`
        SELECT first_name, last_name
        FROM employees
        WHERE termination_date IS NULL
          AND last_name ILIKE 'Smith%'
        ORDER BY last_name, first_name
    `).Scan(&results).Error
	return results, err
}

func (r *AnalyticsRepository) EmployeesNoReviews(ctx context.Context) ([]model.EmployeeNoReview, error) {
	type row struct {
		FirstName string
		LastName  string
		HireDate  time.Time
	}

	var rows []row
	err := r.DB.WithContext(ctx).Raw(`
		SELECT e.first_name, e.last_name, e.hire_date
		FROM employees e
		LEFT JOIN annual_reviews ar ON ar.emp_id = e.id
		WHERE ar.emp_id IS NULL
		ORDER BY e.hire_date
	`).Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	results := make([]model.EmployeeNoReview, 0, len(rows))
	for _, item := range rows {
		results = append(results, model.EmployeeNoReview{
			FirstName: item.FirstName,
			LastName:  item.LastName,
			HireDate:  item.HireDate.Format("2006-01-02"),
		})
	}

	return results, nil
}

func (r *AnalyticsRepository) HireDateDiff(ctx context.Context) (model.HireDateDiff, error) {
	type row struct {
		EarliestHireDate *time.Time
		LatestHireDate   *time.Time
		DifferenceDays   int
	}

	var result row
	if err := r.DB.WithContext(ctx).Raw(`
		SELECT MIN(hire_date) AS earliest_hire_date,
		       MAX(hire_date) AS latest_hire_date,
		       MAX(hire_date)::date - MIN(hire_date)::date AS difference_days
		FROM employees
		WHERE termination_date IS NULL
	`).Scan(&result).Error; err != nil {
		return model.HireDateDiff{}, err
	}

	output := model.HireDateDiff{DifferenceDays: result.DifferenceDays}
	if result.EarliestHireDate != nil {
		output.EarliestHireDate = result.EarliestHireDate.Format("2006-01-02")
	}
	if result.LatestHireDate != nil {
		output.LatestHireDate = result.LatestHireDate.Format("2006-01-02")
	}

	return output, nil
}

func (r *AnalyticsRepository) SalaryProjections(ctx context.Context) ([]model.SalaryProjection, error) {
	var results []model.SalaryProjection
	err := r.DB.WithContext(ctx).Raw(`
		SELECT
			e.id,
			e.first_name,
			e.last_name,
			EXTRACT(YEAR FROM e.hire_date)::int AS hire_year,
			e.salary AS base_salary,
			ROUND(e.salary * POWER(1.15, (2016 - EXTRACT(YEAR FROM e.hire_date))), 2) AS salary_2016,
			COUNT(ar.id) AS review_count
		FROM employees e
		LEFT JOIN annual_reviews ar ON ar.emp_id = e.id
		GROUP BY e.id
		ORDER BY salary_2016 DESC, review_count ASC
	`).Scan(&results).Error

	return results, err
}
