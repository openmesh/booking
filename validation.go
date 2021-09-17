package booking

import (
	"time"
)

// ValidationError is used to return parameter specific issues within an Error.
type ValidationError struct {
	// The property that the validation error has occurred for.
	Name string `json:"name"`
	// A description of the reason that the validation error has occurred.
	Reason string `json:"reason"`
}

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

// processValidation wraps errs in an Error if there is at least one error and
// returns nil otherwise.
func processValidationErrors(errs []ValidationError) error {
	if len(errs) > 0 {
		return Error{
			Code:   EINVALID,
			Detail: "One or more validation errors occurred while processing your request.",
			Title:  "Invalid request",
			Params: errs,
		}
	}
	return nil
}

func validTimezone(tz string) bool {
	validTimezones := []string{
		"UTC-12:00",
		"UTC-11:00",
		"UTC-10:00",
		"UTC-09:30",
		"UTC-09:00",
		"UTC-08:00",
		"UTC-07:00",
		"UTC-06:00",
		"UTC-05:00",
		"UTC-04:00",
		"UTC-03:30",
		"UTC-03:00",
		"UTC-02:00",
		"UTC-01:00",
		"UTC+00:00",
		"UTC±00:00",
		"UTC+01:00",
		"UTC+02:00",
		"UTC+03:00",
		"UTC+03:30",
		"UTC+04:00",
		"UTC+04:30",
		"UTC+05:00",
		"UTC+05:30",
		"UTC+05:45",
		"UTC+06:00",
		"UTC+06:30",
		"UTC+07:00",
		"UTC+08:00",
		"UTC+08:45",
		"UTC+09:00",
		"UTC+09:30",
		"UTC+10:00",
		"UTC+10:30",
		"UTC+11:00",
		"UTC+12:00",
		"UTC+12:45",
		"UTC+13:00",
		"UTC+14:00",
	}
	return Strings(validTimezones).contains(tz)
}

type Strings []string

func (strings Strings) contains(value string) bool {
	for _, s := range strings {
		if value == s {
			return true
		}
	}
	return false
}
