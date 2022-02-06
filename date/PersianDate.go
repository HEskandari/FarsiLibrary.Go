package date

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	GenericFormat      = "yyyy/mm/dd"
	GenericShortFormat = "yyyy/m/d"
	MonthDayFormat     = "MMMM dd"
	MonthYearFormat    = "MMMM, yyyy"
	WrittenFormat      = "W"
)

//PersianDate represents the persian date
type PersianDate struct {
	year  int
	month int
	day   int
}

type parseResult struct {
	error error
	year  int
	month int
	day   int
}

//Year returns the persian date's year value
func (pd PersianDate) Year() int {
	return pd.year
}

//Month returns the persian date's month value
func (pd PersianDate) Month() int {
	return pd.month
}

//Day returns the persian date's day value
func (pd PersianDate) Day() int {
	return pd.day
}

// NewPersianDate creates a valid new instance of PersianDate
func NewPersianDate(year int, month int, day int) (PersianDate, error) {
	err := checkYear(year)
	if err != nil {
		return PersianDate{}, err
	}

	err = checkMonth(month)
	if err != nil {
		return PersianDate{}, err
	}

	err = checkDay(year, month, day)
	if err != nil {
		return PersianDate{}, err
	}

	return PersianDate{year, month, day}, nil
}

//Parse parses a string value to a PersianDate instance. Uses the default separate '/'.
func Parse(value string) (PersianDate, error) {
	return ParseWithSeparator(value, '/')
}

//ParseWithSeparator parses a string value to a PersianDate instance with the given separator.
func ParseWithSeparator(value string, separator rune) (PersianDate, error) {
	parseResult := parse(value, separator)
	if parseResult.error != nil {
		return PersianDate{}, parseResult.error
	}

	return NewPersianDate(parseResult.year, parseResult.month, parseResult.day)
}

//Format formats a PersianDate instance to a string with the given layout.
func (pd PersianDate) Format(layout string) string {
	generic := func(pd PersianDate) string {
		return fmt.Sprintf("%s/%s/%s",
			localizeDigits(pd.Year()),
			localizeDigits(fmt.Sprintf("%02d", pd.Month())),
			localizeDigits(fmt.Sprintf("%02d", pd.Day())))
	}

	switch layout {
	case WrittenFormat:
		return fmt.Sprintf("%s %s %s %s", pd.DayOfWeek(), localizeDigits(pd.Day()), pd.MonthName(), localizeDigits(pd.Year()))
	case MonthYearFormat:
		return fmt.Sprintf("%s %s", pd.MonthName(), localizeDigits(pd.Year()))
	case MonthDayFormat:
		return fmt.Sprintf("%s %s", localizeDigits(pd.Day()), pd.MonthName())
	case GenericShortFormat:
		return fmt.Sprintf("%s/%s/%s", localizeDigits(pd.Year()), localizeDigits(pd.Month()), localizeDigits(pd.Day()))
	case GenericFormat:
		return generic(pd)
	}

	return generic(pd)
}

func parse(value string, separator rune) parseResult {
	if len(value) == 0 || len(value) > 10 {
		return parseResult{error: fmt.Errorf("invalid date string")}
	}

	parts := strings.Split(value, string(separator))
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

//DayOfWeek returns the localized day of the week of the PersianDate
func (pd PersianDate) DayOfWeek() string {
	var dt = ToGregorianDate(pd)
	return localizeDayOfWeek(dt)
}

//MonthName returns the localized month of the PersianDate
func (pd PersianDate) MonthName() string {
	return monthNames[pd.month-1]
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
		if !isJLeapDay(year, month, day) || day > 30 {
			return fmt.Errorf("%d is an invaluid day value", day)
		}
	}

	return nil
}
