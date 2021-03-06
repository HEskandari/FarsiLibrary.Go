package date

import (
	"strconv"
	"strings"
	"time"
)

// jLeap returns one if the specified Persian year is a leap one, otherwise returns zero.
func jLeap(jyear int) int {
	// Is jalali year a leap year?
	_, tmp := divmod(jyear, 33)

	if (tmp == 1) || (tmp == 5) || (tmp == 9) || (tmp == 13) || (tmp == 17) || (tmp == 22) || (tmp == 26) || (tmp == 30) {
		return 1
	}

	return 0
}

func isJLeapDay(jYear int, jMonth int, jDay int) bool {
	if jDay == 30 && jMonth == 12 && isJLeapYear(jYear) {
		return true
	}
	return false
}

func jLeapYears(jYear int) int {
	var div33 = jYear / 33
	var cycle = jYear - (div33 * 33)
	var leap = div33 * 8
	var i int

	if cycle > 0 {
		for i = 1; i <= 18; i += 4 {
			if i > cycle {
				break
			}
			leap++
		}
	}

	if cycle > 21 {
		for i = 22; i <= 31; i += 4 {
			if i > cycle {
				break
			}
			leap++
		}
	}

	return leap
}

// isJLeapYear Checks if a year is a leap one.
func isJLeapYear(jyear int) bool {
	return jLeap(jyear) == 1
}

// jalaliDays calculates total days of jalali years from the base calendar
func jalaliDays(jYear int, jMonth int, jDay int) int {
	var leap = jLeap(jYear)
	for i := 0; i < jMonth-1; i++ {
		jDay += jdayTable[leap][i]
	}
	leap = jLeapYears(jYear - 1)
	return (jYear-1)*365 + leap + jDay
}

// monthDays returns number of days in a month (non-leap year)
func monthDays(monthNo int) int { //nolint
	return jdayTable[1][monthNo-1]
}

// gLeap returns one if the specified Gregorian year is a leap one, otherwise returns zero.
func gLeap(gyear int) int {
	_, mod4 := divmod(gyear, 4)
	_, mod100 := divmod(gyear, 100)
	_, mod400 := divmod(gyear, 400)

	if ((mod4 == 0) && (mod100 != 0)) || (mod400 == 0) {
		return 1
	}
	return 0
}

// gregDays calculates total days of gregorian days from calendar base
func gregDays(gYear int, gMonth int, gDay int) int {
	var div4 = (gYear - 1) / 4
	var div100 = (gYear - 1) / 100
	var div400 = (gYear - 1) / 400
	var leap = gLeap(gYear)

	for i := 0; i < gMonth-1; i++ {
		gDay += gdayTable[leap][i]
	}

	return ((gYear - 1) * 365) + gDay + div4 - div100 + div400
}

func localizeDigits(v interface{}) string {

	tostring := func() string {
		switch t := v.(type) {
		case int:
			return strconv.Itoa(t)
		case string:
			return t
		default:
			return ""
		}
	}

	s := tostring()
	sb := strings.Builder{}
	for _, r := range s {
		if d, ok := digits[r]; ok {
			sb.WriteRune(d)
		}
	}
	return sb.String()
}

func localizeDayOfWeek(t time.Time) string {
	weekday := t.Weekday()
	switch weekday {
	case time.Saturday:
		return weekdays[0]
	case time.Sunday:
		return weekdays[1]
	case time.Monday:
		return weekdays[2]
	case time.Tuesday:
		return weekdays[3]
	case time.Wednesday:
		return weekdays[4]
	case time.Thursday:
		return weekdays[5]
	case time.Friday:
		return weekdays[6]
	}
	return ""
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}
