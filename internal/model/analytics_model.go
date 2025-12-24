package model

type EmployeeName struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type EmployeeNoReview struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HireDate  string `json:"hire_date"`
}

type HireDateDiff struct {
	EarliestHireDate string `json:"earliest_hire_date"`
	LatestHireDate   string `json:"latest_hire_date"`
	DifferenceDays   int    `json:"difference_days"`
}

type SalaryProjection struct {
	ID          int     `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	HireYear    int     `json:"hire_year"`
	BaseSalary  int     `json:"base_salary"`
	Salary2016  float64 `json:"salary_2016"`
	ReviewCount int     `json:"review_count"`
}

type FileResult struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type CityResult struct {
	City        string   `json:"city"`
	Exists      bool     `json:"exists"`
	Suggestions []string `json:"suggestions,omitempty"`
}

type DuplicateCount struct {
	Value int `json:"value"`
	Count int `json:"count"`
}

type RandomStats struct {
	TotalLetters     int `json:"total_letters"`
	TotalVowels      int `json:"total_vowels"`
	TotalNumbers     int `json:"total_numbers"`
	TotalEvenNumbers int `json:"total_even_numbers"`
}

type RandomReport struct {
	Generated            string      `json:"generated"`
	Stats                RandomStats `json:"stats"`
	SortedUnique         []string    `json:"sorted_unique"`
	SortedWithDuplicates []string    `json:"sorted_with_duplicates"`
}

type RemoveNumbersRequest struct {
	Remove []int `json:"remove" validate:"required,min=1,dive"`
}

type AddNumbersRequest struct {
	Add int `json:"add" validate:"required"`
}

type FileQueryRequest struct {
	Name string `form:"name" validate:"required"`
}

type CityQueryRequest struct {
	City string `form:"city" validate:"required"`
}
