package validator

import (
	"strconv"
	"time"
)

func ValidateTotal(year string, month string, day string) (*time.Time, *time.Time, error) {
	var start_date time.Time
	var end_date time.Time
	if isAll(year) {
		start_date = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
		end_date = time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC)
	} else if isRangeYear(month) {
		year_int, err := strconv.Atoi(year)
		if err != nil {
			return nil, nil, err
		}
		start_date = time.Date(year_int, 1, 1, 0, 0, 0, 0, time.UTC)
		end_date = time.Date(year_int, 12, 31, 0, 0, 0, 0, time.UTC)
	} else if isRangeMonth(day) {
		year_int, err := strconv.Atoi(year)
		month_int, err := strconv.Atoi(month)
		if err != nil {
			return nil, nil, err
		}
		start_date = time.Date(year_int, time.Month(month_int), 1, 0, 0, 0, 0, time.UTC)
		end_date = start_date.AddDate(0, 1, -1)
	} else {
		year_int, err := strconv.Atoi(year)
		month_int, err := strconv.Atoi(month)
		day_int, err := strconv.Atoi(day)
		if err != nil {
			return nil, nil, err
		}
		start_date = time.Date(year_int, time.Month(month_int), day_int, 0, 0, 0, 0, time.UTC)
		end_date = time.Date(year_int, time.Month(month_int), day_int, 0, 0, 0, 0, time.UTC)
	}

	return &start_date, &end_date, nil
}

func isAll(year string) bool {
	return year == ""
}

func isRangeYear(month string) bool {
	return month == ""
}

func isRangeMonth(date string) bool {
	return date == ""
}
