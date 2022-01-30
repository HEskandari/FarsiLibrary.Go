package FarsiLibrary_test

import (
	"github.com/HEskandari/FarsiLibrary.Go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPersianDate_ParseString(t *testing.T) {
	pd, err := FarsiLibrary.Parse("1388/02/30")
	assert.Nil(t, err)
	if assert.NotNil(t, pd) {
		assert.Equal(t, 1388, pd.Year())
		assert.Equal(t, 02, pd.Month())
		assert.Equal(t, 30, pd.Day())
	}
}

func TestPersianDate_NumberOfMonths(t *testing.T) {

	// Converts to a leap year in Persian Date (29th Mehr 1387)
	date := time.Date(2008, 10, 20, 0, 0, 0, 0, time.Local)
	pd := FarsiLibrary.ToPersianDate(date)

	if assert.NotNil(t, pd) {
		assert.Equal(t, 1387, pd.Year())
		assert.Equal(t, 29, pd.Day())
		assert.Equal(t, 7, pd.Month())
	}
}

func TestPersianDate_GetDayOfWeek(t *testing.T) {

	pd, _ := FarsiLibrary.NewPersianDate(1387, 7, 7) //7 Mehr equals Yekshanbe
	wd := pd.DayOfWeek()

	if assert.NotNil(t, wd) {
		assert.NotEmpty(t, wd)
		assert.Equal(t, "یکشنبه", wd)
	}
}

func TestPersianDate_GetDayOfWeekFromWeekday(t *testing.T) {
	pd := time.Date(2008, 10, 21, 0, 0, 0, 0, time.Local) // October 30th, Tuesday
	wd := FarsiLibrary.DayOfWeek(pd)

	if assert.NotNil(t, wd) {
		assert.NotEmpty(t, wd)
		assert.Equal(t, "سه‌شنبه", wd)
		t.Logf("Weekday was: %s", wd)
	}
}