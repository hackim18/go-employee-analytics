package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"go-employee-analytics/internal/model"
	"go-employee-analytics/internal/repository"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AnalyticsUseCase struct {
	DB        *gorm.DB
	Log       *logrus.Logger
	Repo      *repository.AnalyticsRepository
	OutputDir string
}

func NewAnalyticsUseCase(db *gorm.DB, log *logrus.Logger, outputDir string, repo *repository.AnalyticsRepository) *AnalyticsUseCase {
	return &AnalyticsUseCase{
		DB:        db,
		Log:       log,
		Repo:      repo,
		OutputDir: outputDir,
	}
}

func (u *AnalyticsUseCase) ActiveSmithEmployees(ctx context.Context) ([]model.EmployeeName, error) {
	return u.Repo.ActiveSmithEmployees(ctx)
}

func (u *AnalyticsUseCase) EmployeesNoReviews(ctx context.Context) ([]model.EmployeeNoReview, error) {
	return u.Repo.EmployeesNoReviews(ctx)
}

func (u *AnalyticsUseCase) HireDateDiff(ctx context.Context) (model.HireDateDiff, error) {
	return u.Repo.HireDateDiff(ctx)
}

func (u *AnalyticsUseCase) SalaryProjections(ctx context.Context) ([]model.SalaryProjection, error) {
	return u.Repo.SalaryProjections(ctx)
}

func (u *AnalyticsUseCase) SaveOutputs(ctx context.Context) ([]model.FileResult, error) {
	q2, err := u.ActiveSmithEmployees(ctx)
	if err != nil {
		return nil, err
	}

	q3, err := u.EmployeesNoReviews(ctx)
	if err != nil {
		return nil, err
	}

	q4, err := u.HireDateDiff(ctx)
	if err != nil {
		return nil, err
	}

	q5, err := u.SalaryProjections(ctx)
	if err != nil {
		return nil, err
	}

	files := make([]model.FileResult, 0, 4)
	if path, err := u.SaveJSONFile("contoh2.txt", q2); err != nil {
		return nil, err
	} else {
		files = append(files, model.FileResult{Name: "contoh2.txt", Path: path})
	}

	if path, err := u.SaveJSONFile("contoh3.txt", q3); err != nil {
		return nil, err
	} else {
		files = append(files, model.FileResult{Name: "contoh3.txt", Path: path})
	}

	if path, err := u.SaveJSONFile("contoh4.txt", q4); err != nil {
		return nil, err
	} else {
		files = append(files, model.FileResult{Name: "contoh4.txt", Path: path})
	}

	if path, err := u.SaveJSONFile("contoh5.txt", q5); err != nil {
		return nil, err
	} else {
		files = append(files, model.FileResult{Name: "contoh5.txt", Path: path})
	}

	return files, nil
}

func (u *AnalyticsUseCase) SaveJSONFile(name string, v any) (string, error) {
	fileName, err := sanitizeFileName(name)
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(u.OutputDir, 0755); err != nil {
		return "", err
	}

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}

	path := filepath.Join(u.OutputDir, fileName)
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", err
	}

	return path, nil
}

func (u *AnalyticsUseCase) ReadJSONFile(name string) ([]byte, error) {
	fileName, err := sanitizeFileName(name)
	if err != nil {
		return nil, err
	}

	path := filepath.Join(u.OutputDir, fileName)
	return os.ReadFile(path)
}

func sanitizeFileName(name string) (string, error) {
	if name == "" {
		return "", errors.New("filename is required")
	}

	if strings.ContainsAny(name, "\\/") {
		return "", errors.New("invalid filename")
	}

	base := filepath.Base(name)
	if base != name {
		return "", errors.New("invalid filename")
	}

	if !strings.HasSuffix(strings.ToLower(base), ".txt") {
		return "", errors.New("filename must end with .txt")
	}

	return base, nil
}
