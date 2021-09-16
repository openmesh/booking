package booking

import (
	"time"
)

// validateSlotTime returns an error if the provided string is not a valid time in the format HH:MM.
// Primarily used for validating slot times.
func validateSlotTime(t string) error {
	_, err := time.Parse("15:04", t)
	return err
}

// timePrecedes returns a boolean indicating whether the parsed startTime is chronologically
// before the parsed endTime. Returns an error if parsing fails for any reason.
func timePrecedes(startTime, endTime string) (bool, error) {
	parsedStartTime, err := time.Parse("15:04", startTime)
	if err != nil {
		return false, err
	}
	parsedEndTime, err := time.Parse("15:04", endTime)
	if err != nil {
		return false, err
	}
	return parsedStartTime.Before(parsedEndTime), nil
}

// timeEqual returns a boolean indicating whether two parsed times are equal.
// Returns an error if parsing fails.
func timeEqual(time1, time2 string) (bool, error) {
	parsedTime1, err := time.Parse("15:04", time1)
	if err != nil {
		return false, err
	}
	parsedTime2, err := time.Parse("15:04", time2)
	if err != nil {
		return false, err
	}
	return parsedTime1.Equal(parsedTime2), nil
}
