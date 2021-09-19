package booking

import (
	"errors"
	"time"
)

var validTimeFormats = []string{
	time.RFC3339,
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
}

func ParseTime(input string) (time.Time, error) {
	for _, format := range validTimeFormats {
		t, err := time.Parse(format, input)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("unrecognized time format")
}
