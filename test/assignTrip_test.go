package test

import (
	"testing"

	"example.com/Ride_simulator/models"
)

type AssignTest struct {
	OnlineDriver []int64
	InTrip       []int64
	Expected     int64
}

func TestAssign(t *testing.T) {

	testCases := []AssignTest{
		{
			OnlineDriver: []int64{1, 2, 3, 4, 5},
			InTrip:       []int64{3, 4, 5},
			Expected:     1, // Eligible driver is 1 (online but not in trip)
		},
		{
			OnlineDriver: []int64{10, 20, 30},
			InTrip:       []int64{10, 20, 30},
			Expected:     0, // No eligible driver (all online drivers are in trip)
		},
		{
			OnlineDriver: []int64{7, 8, 9},
			InTrip:       []int64{},
			Expected:     7, // First eligible driver is 7 (no drivers are in trip)
		},
		{
			OnlineDriver: []int64{},
			InTrip:       []int64{1, 2, 3},
			Expected:     0, // No eligible driver (no online drivers available)
		},
	}

	for _, test := range testCases {
		got := models.EligableDriverID(test.OnlineDriver, test.InTrip)

		if got != test.Expected {
			t.Errorf("For OnlineDriver %v and InTrip %v, got %v, wanted %v",
				test.OnlineDriver, test.InTrip, got, test.Expected)
		} else {
			t.Logf("For OnlineDriver %v and InTrip %v, got %v as expected.",
				test.OnlineDriver, test.InTrip, got)
		}
	}
}
