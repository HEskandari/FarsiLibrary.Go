package FarsiLibrary

import (
	"math"
	"time"
)

const (
	yearOffset = 226894
	solar = 365.25
)

var weekdays = []string{"شنبه", "یکشنبه", "دوشنبه", "سه‌شنبه", "چهارشنبه", "پنجشنبه", "جمعه"}
var weekdaysabbr = []string{"ش", "ی", "د", "س", "چ", "پ", "ج"}
var gdayTable = [12][12]int{{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}, {31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}}
var jdayTable = [12][12]int{{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}, {31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 30}}


// ToPersianDate converts a time.Time to Persian Date.
func ToPersianDate(time time.Time) PersianDate {
	var gyear, gmonth, gday = time.Date()
	var i int

	// Calculate total days from the base of gregorian calendar
	var iTotalDays = GregDays(gyear, int(gmonth), gday)
	iTotalDays -= yearOffset

	// Calculate total jalali years passed
	var jyear = int(float64(iTotalDays) / (solar -0.25/33.0))

	// Calculate passed leap years
	var leap = JLeapYears(jyear)

	// Calculate total days from the base of jalali calendar
	var jday = iTotalDays - (365*jyear + leap)

	// Calculate the correct year of jalali calendar
	jyear++

	if jday == 0 {
		jyear--
		if JLeap(jyear) == 1 {
			jday = 366
		} else {
			jday = 365
		}
	} else if jday == 366 && JLeap(jyear) != 1 {
		jday = 1
		jyear++
	}

	// Calculate correct month of jalali calendar
	leap = JLeap(jyear)
	for i = 0; i <= 12; i++ {
		if jday <= jdayTable[leap][i] {
			break
		}
		jday -= jdayTable[leap][i]
	}

	var iJMonth = i + 1

	return PersianDate{jyear, iJMonth, jday}
}

func ToGregorianDate(date *PersianDate) time.Time {
	var jyear = date.Year()
	var jmonth = date.Month()
	var jday = date.Day()
	var i int

	var totalDays = JalaliDays(jyear, jmonth, jday)
	totalDays += yearOffset

	var gyear = totalDays / int(math.Floor(solar-0.25/33))
	var div4 = gyear / 4
	var div100 = gyear / 100
	var div400 = gyear / 400
	var gdays = totalDays - (365 * gyear) - (div4 - div100 + div400)

	gyear++

	if gdays == 0 {
		gyear--
		if GLeap(gyear) == 1 {
			gdays = 366
		} else {
			gdays = 365
		}
	} else if gdays == 366 && GLeap(gyear) != 1 {
		gdays = 1
		gyear++
	}

	var leap = GLeap(gyear)
	for i = 0; i <= 12; i++ {
		if gdays <= gdayTable[leap][i] {
			break
		}
		gdays -= gdayTable[leap][i]
	}

	var iGMonth = i + 1
	var iGDay = gdays

	return time.Date(gyear, time.Month(iGMonth), iGDay, 0, 0, 0, 0, time.Local)
}