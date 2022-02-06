package date_test

import (
	"github.com/heskandari/farsilibrary.go/date"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPersianDate_ParseString(t *testing.T) {
	pd, err := date.Parse("1388/02/30")
	assert.Nil(t, err)
	if assert.NotNil(t, pd) {
		assert.Equal(t, 1388, pd.Year())
		assert.Equal(t, 02, pd.Month())
		assert.Equal(t, 30, pd.Day())
	}
}

func TestPersianDate_Format(t *testing.T) {
	pd, _ := date.NewPersianDate(1388, 2, 3)
	if assert.NotNil(t, pd) {
		assert.Equal(t, "۱۳۸۸/۰۲/۰۳", pd.Format(date.GenericFormat))
		assert.Equal(t, "اردیبهشت ۱۳۸۸", pd.Format(date.MonthYearFormat))
		assert.Equal(t, "۳ اردیبهشت", pd.Format(date.MonthDayFormat))
		assert.Equal(t, "۱۳۸۸/۲/۳", pd.Format(date.GenericShortFormat))
		assert.Equal(t, "پنجشنبه ۳ اردیبهشت ۱۳۸۸", pd.Format(date.WrittenFormat))
		assert.Equal(t, "۱۳۸۸/۰۲/۰۳", pd.Format(""))
	}
}

func TestPersianDate_NumberOfMonths(t *testing.T) {

	// Converts to a leap year in Persian Date (29th Mehr 1387)
	gDate := time.Date(2008, 10, 20, 0, 0, 0, 0, time.Local)
	pd := date.ToPersianDate(gDate)

	if assert.NotNil(t, pd) {
		assert.Equal(t, 1387, pd.Year())
		assert.Equal(t, 29, pd.Day())
		assert.Equal(t, 7, pd.Month())
	}
}

func TestPersianDate_GetDayOfWeek(t *testing.T) {

	pd, _ := date.NewPersianDate(1387, 7, 7) //7 Mehr equals Yekshanbe
	wd := pd.DayOfWeek()

	if assert.NotNil(t, wd) {
		assert.NotEmpty(t, wd)
		assert.Equal(t, "یکشنبه", wd)
	}
}
