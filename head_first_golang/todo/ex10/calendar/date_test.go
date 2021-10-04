package calendar

import (
	"testing"
)

var date Date

func TestDate_Day(t *testing.T) {
	if err := date.SetDay(0); err == nil {
		t.Fail()
	}
	if err := date.SetDay(2); err != nil {
		t.Log("Success")
	}
	if date.Day() == 1 {
		return
	}
}

func TestDate_Year(t *testing.T) {
	if err := date.SetYear(0); err == nil {
		t.Errorf(err.Error())
	}
	if err := date.SetYear(2); err != nil {
		t.Log("Success")
	}
	if date.Year() == 1 {
		return
	}
}

func TestDate_Month(t *testing.T) {
	if err := date.SetMonth(0); err == nil {
		t.Errorf(err.Error())
	}
	if err := date.SetMonth(2); err != nil {
		t.Log("Success")
	}
	if date.Month() == 1 {
		return
	}
}
