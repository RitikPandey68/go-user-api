package service

import (
	"testing"
	"time"
)

// TestCalculateAge tests the CalculateAge function for three distinct cases.
//
// Case 1: Birthday has already passed this year → full year counted
// Case 2: Birthday is exactly today → no subtraction
// Case 3: Birthday has not yet occurred this year → subtract one year
func TestCalculateAge(t *testing.T) {
	now := time.Now()

	t.Run("Birthday already passed this year", func(t *testing.T) {
		// Person born exactly 30 years ago today — birthday already happened
		dob := time.Date(now.Year()-30, now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		// Shift dob by 1 day back to ensure birthday is "in the past" even for edge dates
		dob = dob.AddDate(0, 0, -1)
		expected := 30
		got := CalculateAge(dob)
		if got != expected {
			t.Errorf("Birthday already passed: expected age %d, got %d (dob=%s)", expected, got, dob.Format("2006-01-02"))
		}
	})

	t.Run("Birthday is today", func(t *testing.T) {
		// Person born exactly 25 years ago today
		dob := time.Date(now.Year()-25, now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		expected := 25
		got := CalculateAge(dob)
		if got != expected {
			t.Errorf("Birthday today: expected age %d, got %d (dob=%s)", expected, got, dob.Format("2006-01-02"))
		}
	})

	t.Run("Birthday upcoming this year", func(t *testing.T) {
		// Person born 20 years ago but birthday falls tomorrow → still 19
		tomorrow := now.AddDate(0, 0, 1)
		dob := time.Date(now.Year()-20, tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, time.UTC)
		expected := 19
		got := CalculateAge(dob)
		if got != expected {
			t.Errorf("Birthday upcoming: expected age %d, got %d (dob=%s)", expected, got, dob.Format("2006-01-02"))
		}
	})
}
