package usecase

import (
	"go-employee-analytics/internal/model"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func (u *AnalyticsUseCase) CityLookup(city string) model.CityResult {
	cleaned := strings.TrimSpace(city)
	result := model.CityResult{City: cleaned}
	if cleaned == "" {
		return result
	}

	for _, item := range cityList {
		if strings.EqualFold(item, cleaned) {
			result.Exists = true
			return result
		}
	}

	runes := []rune(cleaned)
	if len(runes) == 0 {
		return result
	}

	first := strings.ToLower(string(runes[0]))
	last := strings.ToLower(string(runes[len(runes)-1]))

	suggestions := []string{}
	seen := map[string]bool{}
	for _, item := range cityList {
		lower := strings.ToLower(item)
		if strings.HasPrefix(lower, first) || strings.HasSuffix(lower, last) {
			if !seen[item] {
				suggestions = append(suggestions, item)
				seen[item] = true
			}
		}
	}

	result.Suggestions = suggestions
	return result
}

func (u *AnalyticsUseCase) SortedUniqueNumbers() []int {
	unique := map[int]bool{}
	for _, value := range numberList {
		unique[value] = true
	}

	results := make([]int, 0, len(unique))
	for value := range unique {
		results = append(results, value)
	}

	sort.Ints(results)
	return results
}

func (u *AnalyticsUseCase) DuplicateCounts() []model.DuplicateCount {
	counts := map[int]int{}
	for _, value := range numberList {
		counts[value]++
	}

	results := make([]model.DuplicateCount, 0, len(counts))
	for value, count := range counts {
		results = append(results, model.DuplicateCount{Value: value, Count: count})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Value < results[j].Value
	})

	return results
}

func (u *AnalyticsUseCase) RemoveNumbers(remove []int) []int {
	removeSet := map[int]bool{}
	for _, value := range remove {
		removeSet[value] = true
	}

	results := make([]int, 0, len(numberList))
	for _, value := range numberList {
		if !removeSet[value] {
			results = append(results, value)
		}
	}

	return results
}

func (u *AnalyticsUseCase) AddWithCap(add int) []int {
	results := make([]int, 0, len(numberList))
	for _, value := range numberList {
		updated := value + add
		if updated > 10 {
			updated = 10
		}
		results = append(results, updated)
	}

	return results
}

func (u *AnalyticsUseCase) GenerateRandomReport() model.RandomReport {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	letters := make([]byte, 50)
	digits := make([]byte, 50)
	for i := 0; i < 50; i++ {
		letters[i] = byte('a' + rng.Intn(26))
		digits[i] = byte('0' + rng.Intn(10))
	}

	combined := append(letters, digits...)
	rng.Shuffle(len(combined), func(i, j int) {
		combined[i], combined[j] = combined[j], combined[i]
	})

	stats := calculateStats(combined)
	uniqueSorted := sortedUnique(combined)
	sortedWithDup := sortedWithDuplicates(combined)

	return model.RandomReport{
		Generated:            string(combined),
		Stats:                stats,
		SortedUnique:         uniqueSorted,
		SortedWithDuplicates: sortedWithDup,
	}
}

func calculateStats(data []byte) model.RandomStats {
	stats := model.RandomStats{}
	for _, value := range data {
		if value >= '0' && value <= '9' {
			stats.TotalNumbers++
			if (value-'0')%2 == 0 {
				stats.TotalEvenNumbers++
			}
			continue
		}
		if value >= 'a' && value <= 'z' {
			stats.TotalLetters++
			switch value {
			case 'a', 'e', 'i', 'o', 'u':
				stats.TotalVowels++
			}
		}
	}

	return stats
}

func sortedUnique(data []byte) []string {
	digitSet := map[byte]bool{}
	letterSet := map[byte]bool{}

	for _, value := range data {
		if value >= '0' && value <= '9' {
			digitSet[value] = true
			continue
		}
		if value >= 'a' && value <= 'z' {
			letterSet[value] = true
		}
	}

	digits := make([]byte, 0, len(digitSet))
	for value := range digitSet {
		digits = append(digits, value)
	}
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] > digits[j]
	})

	letters := make([]byte, 0, len(letterSet))
	for value := range letterSet {
		letters = append(letters, value)
	}
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	results := make([]string, 0, len(digits)+len(letters))
	for _, value := range digits {
		results = append(results, string(value))
	}
	for _, value := range letters {
		results = append(results, string(value))
	}

	return results
}

func sortedWithDuplicates(data []byte) []string {
	digits := make([]byte, 0, len(data))
	letters := make([]byte, 0, len(data))

	for _, value := range data {
		if value >= '0' && value <= '9' {
			digits = append(digits, value)
			continue
		}
		if value >= 'a' && value <= 'z' {
			letters = append(letters, value)
		}
	}

	sort.Slice(digits, func(i, j int) bool {
		return digits[i] > digits[j]
	})
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	results := make([]string, 0, len(digits)+len(letters))
	for _, value := range digits {
		results = append(results, string(value))
	}
	for _, value := range letters {
		results = append(results, string(value))
	}

	return results
}
