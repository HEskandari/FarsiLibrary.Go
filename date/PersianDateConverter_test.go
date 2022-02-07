package date_test

import (
	"github.com/heskandari/farsilibrary.go/date"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPersianDateConverter_ConvertLeapYears(t *testing.T) {

	// Converts to a leap year in Persian Date (30th Esfand 1387)
	gDate := time.Date(2009, 3, 20, 0, 0, 0, 0, time.Local)
	pd := date.ToPersianDate(gDate)

	if assert.NotNil(t, pd) {
		assert.Equal(t, 1387, pd.Year())
		assert.Equal(t, 30, pd.Day())
		assert.Equal(t, 12, pd.Month())
	}
}

func TestPersianDateConverter_ConvertNonLeapYears(t *testing.T) {

	// Converts to a leap year in Persian Date (29th Mehr 1387)
	gDate := time.Date(2008, 10, 20, 0, 0, 0, 0, time.Local)
	pd := date.ToPersianDate(gDate)

	if assert.NotNil(t, pd) {
		assert.Equal(t, 1387, pd.Year())
		assert.Equal(t, 29, pd.Day())
		assert.Equal(t, 7, pd.Month())
	}
}

func TestPersianDateConverter_ConvertAndConvertBack(t *testing.T) {

	// Converts to a leap year in Persian Date (29th Mehr 1387)
	gDate := time.Date(2020, 01, 01, 0, 0, 0, 0, time.Local)
	pd := date.ToPersianDate(gDate)
	gd := date.ToGregorianDate(pd)

	if assert.NotNil(t, gd) {
		assert.Equal(t, 2020, gd.Year())
		assert.Equal(t, 01, gd.Day())
		assert.Equal(t, 01, int(gd.Month()))
	}
}
