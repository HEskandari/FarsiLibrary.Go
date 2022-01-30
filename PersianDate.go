package FarsiLibrary

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	GenericFormat     = "yyyy/mm/dd"
	MonthDayFormat    = "MMMM dd"
	MonthYearFormat   = "MMMM, yyyy"
	WrittenFormat     = "W"
)

type PersianDate struct {
	year  int
	month int
	day   int
}

type parseResult struct {
	error error
	year int
	month int
	day int
}

func (pd *PersianDate) Year() int {
	return pd.year
}

func (pd *PersianDate) Month() int {
	return pd.month
}

func (pd *PersianDate) Day() int {
	return pd.day
}

// NewPersianDate creates a valid new instance of PersianDate
func NewPersianDate(year int, month int, day int) (*PersianDate, error) {
	err := checkYear(year)
	if err != nil {
		return nil, err
	}

	err = checkMonth(month)
	if err != nil {
		return nil, err
	}

	err = checkDay(year, month, day)
	if err != nil {
		return nil, err
	}

	return &PersianDate{year, month, day}, nil
}

func Parse(value string) (*PersianDate, error) {
	parseResult := parse(value)
	if parseResult.error != nil {
		return nil, parseResult.error
	}

	return NewPersianDate(parseResult.year, parseResult.month, parseResult.day)
}

func (pd *PersianDate) Format(layout string) string

}

func parse(value string) parseResult {
	if len(value) == 0 || len(value) > 10 {
		return parseResult{error: fmt.Errorf("invalid date string")}
	}

	parts := strings.Split(value, "/")
	if len(parts) != 3 {
		return parseResult{error: fmt.Errorf("invalid date string")}
	}

	partYear := parts[0]
	partMonth := parts[1]
	partDay := parts[2]

	if len(partYear) != 4 {
		return parseResult{error: fmt.Errorf("invalid year in the value string: %s", partYear)}
	}

	if len(partMonth) == 0 || len(partMonth) > 2 {
		return parseResult{error: fmt.Errorf("invalid month in the value string: %s", partMonth)}
	}

	if len(partDay) == 0 || len(partDay) > 2 {
		return parseResult{error: fmt.Errorf("invalid day in the value string: %s", partDay)}
	}

	year, err := strconv.Atoi(partYear)
	if err != nil {
		return parseResult{error: fmt.Errorf("year value cannot be parsed: %d", year)}
	}

	month, err := strconv.Atoi(partMonth)
	if err != nil {
		return parseResult{error: fmt.Errorf("year value cannot be parsed: %d", month)}
	}

	day, err := strconv.Atoi(partDay)
	if err != nil {
		return parseResult{error: fmt.Errorf("year value cannot be parsed: %d", day)}
	}

	return parseResult{year: year, month: month, day: day}
}

func (pd *PersianDate) DayOfWeek() string {
	var dt = ToGregorianDate(pd)
	return DayOfWeek(dt)
}

func checkYear(year int) error {
	if year < 1 || year > 9999 {
		return fmt.Errorf("%d is an invalid year value", year)
	}
	return nil
}

func checkMonth(month int) error {
	if month > 12 || month < 1 {
		return fmt.Errorf("%d is an invaluid month value", month)
	}
	return nil
}

func checkDay(year int, month int, day int) error {
	if month < 6 && day > 31 {
		return fmt.Errorf("%d is an invaluid day value", day)
	}

	if month > 6 && day > 30 {
		return fmt.Errorf("%d is an invaluid day value", day)
	}

	if month == 12 && day > 29 {
		if !IsJLeapDay(year, month, day) || day > 30 {
			return fmt.Errorf("%d is an invaluid day value", day)
		}
	}

	return nil
}