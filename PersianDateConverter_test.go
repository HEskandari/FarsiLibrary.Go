package FarsiLibrary

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPersianDateConverter_ConvertLeapYears(t *testing.T) {

	//Converts to a leap year in Persian Date (30th Esfand 1387)
	date := time.Date(2009, 3, 20, 0, 0, 0, 0, time.Local)
	converter := PersianDateConverter{}
	pd := converter.ToPersianDate(date)

	if assert.NotNil(t, pd) {
		assert.Equal(t, 1387, pd.Year)
		assert.Equal(t, 30, pd.Day)
		assert.Equal(t, 12, pd.Month)
	}
}

func TestPersianDateConverter_ConvertNonLeapYears(t *testing.T) {

	//Converts to a leap year in Persian Date (29th Mehr 1387)
	date := time.Date(2008, 10, 20, 0, 0, 0, 0, time.Local)
	converter := PersianDateConverter{}
	pd := converter.ToPersianDate(date)

	if assert.NotNil(t, pd) {
		assert.Equal(t, 1387, pd.Year)
		assert.Equal(t, 29, pd.Day)
		assert.Equal(t, 7, pd.Month)
	}
}
