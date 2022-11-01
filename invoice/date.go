package invoice

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

func NewDate(dateStr string) (Date, error) {
	time, err := parseDate(dateStr)
	if err != nil {
		return Date{}, err
	}

	return Date{time}, nil
}

func NewDateFromTime(t time.Time) Date {
	return Date{t}
}

func (d *Date) UnmarshalJSON(data []byte) error {
	d.Time = time.Time{}

	var raw interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	stringVal, isString := raw.(string)
	if !isString {
		return errors.New("date must be a string")
	}

	d.Time, err = parseDate(stringVal)
	return err
}

func (d Date) MarshalJSON() ([]byte, error) {
	dateStr := d.Format("2006-01-02")
	jsonStr := fmt.Sprintf(`"%s"`, dateStr)

	return []byte(jsonStr), nil
}

func parseDate(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)

	if dateStr == "" {
		return time.Time{}, nil
	}

	allowableFormats := []string{
		time.RFC3339,
		"2006-01-02", "2006-1-02", "2006-01-2", "2006-1-2",
		"01/02/2006", "1/2/2006", "01/2/2006", "1/02/2006",
	}

	for _, format := range allowableFormats {
		parsedTime, err := time.Parse(format, dateStr)
		if err == nil {
			return parsedTime, nil
		}
	}

	return time.Time{}, fmt.Errorf(`could not parse date "%v"`, dateStr)
}
