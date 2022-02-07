package date_test

import (
	"encoding/json"
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
		assert.Equal(t, "1388-02-03", pd.Format(date.Serialized))
		assert.Equal(t, "1388/02/03", pd.Format(""))
		assert.Equal(t, "1388/02/03", pd.String())
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

func TestPersianDate_Today(t *testing.T) {

	pd := date.Today()

	if assert.NotNil(t, pd) {
		assert.NotEmpty(t, pd.String())
		assert.Equal(t, 10, len(pd.String()))
	}
}

type Person struct {
	Name string
	DoB  *date.PersianDate
}

func TestPersianDate_JsonMarshalling(t *testing.T) {

	dob, _ := date.NewPersianDate(1400, 01, 01)
	p := Person{
		Name: "John",
		DoB:  &dob,
	}

	serialized, err := json.Marshal(p)

	assert.NoError(t, err)
	assert.NotEmpty(t, serialized)
	assert.JSONEq(t, `{"Name": "John", "DoB": "1400-01-01"}`, string(serialized))
}

func TestPersianDate_JsonUnMarshalling(t *testing.T) {

	js := `{"Name": "John", "DoB": "1400-02-01"}`
	var p = Person{}

	err := json.Unmarshal([]byte(js), &p)

	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.DoB.Day(), 1)
	assert.Equal(t, p.DoB.Month(), 2)
	assert.Equal(t, p.DoB.Year(), 1400)
}
