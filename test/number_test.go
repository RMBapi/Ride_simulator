package test

import (
	"testing"

	"example.com/Ride_simulator/models"
)

type NumberTest struct {
	PhoneNumber string
	want        bool
}

func TestNumber(t *testing.T) {

	testCases := []NumberTest{
		{"01234567891", false}, // Valid
		{"01345678902", true},  // Valid
		{"01987654323", true},  // Valid
		{"01123456785", false}, // Invalid (third digit is 1)
		{"01276543216", false}, // Invalid (third digit is 2)
		{"0123abc4567", false}, // Invalid (contains letters)
		{"01234567", false},    // Invalid (too short)
	}

	for _, test := range testCases {
		got := models.IsValidPhoneNumber(test.PhoneNumber)
		if got != test.want {
			t.Errorf("For phone number %q, got %v, wanted %v", test.PhoneNumber, got, test.want)
		} else {
			t.Logf("For phone number %q, got %v as expected.", test.PhoneNumber, got)
		}
	}
}
