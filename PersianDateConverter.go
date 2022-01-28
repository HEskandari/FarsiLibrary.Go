package FarsiLibrary

import (
	"math"
	"time"
)

var GYearOff = 226894
var Solar = 365.25
var weekdays = []string{"شنبه", "یکشنبه", "دوشنبه", "سه شنبه", "چهارشنبه", "پنجشنبه", "جمعه"}
var weekdaysabbr = []string{"ش", "ی", "د", "س", "چ", "پ", "ج"}
var gdayTable = [12][12]int{{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}, {31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}}
var jdayTable = [12][12]int{{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}, {31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 30}}

type PersianDate struct {
	Year  int
	Month int
	Day   int
}

type PersianDateConverter struct {
}

// JLeap returns one if the specified Persian year is a leap one, otherwise returns zero.
func (c *PersianDateConverter) JLeap(jyear int) int {
	//Is jalali year a leap year?
	_, tmp := divmod(jyear, 33)

	if (tmp == 1) || (tmp == 5) || (tmp == 9) || (tmp == 13) || (tmp == 17) || (tmp == 22) || (tmp == 26) || (tmp == 30) {
		return 1
	}

	return 0
}

// IsJLeapYear Checks if a year is a leap one.
func (c *PersianDateConverter) IsJLeapYear(jyear int) bool {
	return c.JLeap(jyear) == 1
}

// GLeap returns one if the specified Gregorian year is a leap one, otherwise returns zero.
func (c *PersianDateConverter) GLeap(gyear int) int {
	_, mod4 := divmod(gyear, 4)
	_, mod100 := divmod(gyear, 100)
	_, mod400 := divmod(gyear, 400)

	if ((mod4 == 0) && (mod100 != 0)) || (mod400 == 0) {
		return 1
	}
	return 0
}

// GregDays calculates total days of gregorian days from calendar base
func (c *PersianDateConverter) GregDays(gYear int, gMonth int, gDay int) int {
	var div4 = (gYear - 1) / 4
	var div100 = (gYear - 1) / 100
	var div400 = (gYear - 1) / 400
	var leap = c.GLeap(gYear)

	for i := 0; i < gMonth-1; i++ {
		gDay = gDay + gdayTable[leap][i]
	}

	return ((gYear - 1) * 365) + gDay + div4 - div100 + div400
}

func (c *PersianDateConverter) JLeapYears(jYear int) int {
	var div33 = jYear / 33
	var cycle = jYear - (div33 * 33)
	var leap = div33 * 8
	var i int

	if cycle > 0 {
		for i = 1; i <= 18; i = i + 4 {
			if i > cycle {
				break
			}
			leap++
		}
	}

	if cycle > 21 {
		for i = 22; i <= 31; i = i + 4 {
			if i > cycle {
				break
			}
			leap++
		}
	}

	return leap
}

// JalaliDays calculates total days of jalali years from the base calendar
func (c *PersianDateConverter) JalaliDays(jYear int, jMonth int, jDay int) int {
	var leap = c.JLeap(jYear)
	for i := 0; i < jMonth-1; i++ {
		jDay = jDay + jdayTable[leap][i]
	}
	leap = c.JLeapYears(jYear - 1)
	return (jYear-1)*365 + leap + jDay
}

// ToPersianDate converts a time.Time to Persian Date.
func (c *PersianDateConverter) ToPersianDate(time time.Time) PersianDate {
	var gyear, gmonth, gday = time.Date()
	var i int

	//Calculate total days from the base of gregorian calendar
	var iTotalDays = c.GregDays(gyear, int(gmonth), gday)
	iTotalDays = iTotalDays - GYearOff

	//Calculate total jalali years passed
	var jyear = iTotalDays / int(math.Floor(Solar-0.25/33.0))

	//Calculate passed leap years
	var leap = c.JLeapYears(jyear)

	//Calculate total days from the base of jalali calendar
	var jday = iTotalDays - (365*jyear + leap)

	//Calculate the correct year of jalali calendar
	jyear++

	if jday == 0 {
		jyear--
		if c.JLeap(jyear) == 1 {
			jday = 366
		} else {
			jday = 365
		}
	} else {
		if jday == 366 && c.JLeap(jyear) != 1 {
			jday = 1
			jyear++
		}
	}

	//Calculate correct month of jalali calendar
	leap = c.JLeap(jyear)
	for i = 0; i <= 12; i++ {
		if jday <= jdayTable[leap][i] {
			break
		}
		jday = jday - jdayTable[leap][i]
	}

	var iJMonth = i + 1

	return PersianDate{jyear, iJMonth, jday}
}

func (c *PersianDateConverter) ToGregorianDate(date PersianDate) time.Time {
	var jyear = date.Year
	var jmonth = date.Month
	var jday = date.Day
	var i int

	var totalDays = c.JalaliDays(jyear, jmonth, jday)
	totalDays = totalDays + GYearOff

	var gyear = totalDays / int(math.Floor(Solar-0.25/33))
	var Div4 = gyear / 4
	var Div100 = gyear / 100
	var Div400 = gyear / 400
	var gdays = totalDays - (365 * gyear) - (Div4 - Div100 + Div400)
	gyear = gyear + 1

	if gdays == 0 {
		gyear--
		if c.GLeap(gyear) == 1 {
			gdays = 366
		} else {
			gdays = 365
		}
	} else {
		if gdays == 366 && c.GLeap(gyear) != 1 {
			gdays = 1
			gyear++
		}
	}

	var leap = c.GLeap(gyear)
	for i = 0; i <= 12; i++ {
		if gdays <= gdayTable[leap][i] {
			break
		}
		gdays = gdays - gdayTable[leap][i]
	}

	var iGMonth = i + 1
	var iGDay = gdays

	return time.Date(gyear, time.Month(iGMonth), iGDay, 0, 0, 0, 0, time.Local)
}

func (c *PersianDateConverter) MonthDays(monthNo int) int {
	return jdayTable[1][monthNo-1]
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}
