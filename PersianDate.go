package FarsiLibrary

import (
	"fmt"
)

type PersianDate struct {
	year  int
	month int
	day   int
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