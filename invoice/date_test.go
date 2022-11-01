package invoice

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNewDate(t *testing.T) {
	date, err := NewDate("11/22/2022")
	if err != nil {
		t.Error(err)
	}

	expectedTime, err := time.Parse(time.RFC3339, "2022-11-22T00:00:00Z")
	if err != nil {
		t.Error(err)
	}

	if date.Time != expectedTime {
		t.Fatalf("Expected %v, got %v", expectedTime, date.Time)
	}
}

func TestDate_UnmarshalJSON(t *testing.T) {
	assertDateJSON(t, `""`, "0001-01-01T00:00:00Z")
	assertDateJSON(t, `"2022-03-02"`, "2022-03-02T00:00:00Z")
	assertDateJSON(t, `"2022-3-02"`, "2022-03-02T00:00:00Z")
	assertDateJSON(t, `"2022-03-2"`, "2022-03-02T00:00:00Z")
	assertDateJSON(t, `"2022-3-2"`, "2022-03-02T00:00:00Z")
	assertDateJSON(t, `"03/02/2022"`, "2022-03-02T00:00:00Z")
	assertDateJSON(t, `"03/2/2022"`, "2022-03-02T00:00:00Z")
	assertDateJSON(t, `"3/02/2022"`, "2022-03-02T00:00:00Z")
	assertDateJSON(t, `"3/2/2022"`, "2022-03-02T00:00:00Z")
	assertDateJSON(t, `"2022-03-02T00:00:00Z"`, "2022-03-02T00:00:00Z")
	assertDateJSONErr(t, `"not a date"`)
}

func TestDate_MarshalJSON(t *testing.T) {
	d, _ := NewDate("11/22/2022")

	jsonData, err := json.Marshal(d)
	if err != nil {
		t.Error(err)
	}

	expectedJSON := `"2022-11-22"`

	if string(jsonData) != expectedJSON {
		t.Fatalf("Expected %v, got %v", expectedJSON, string(jsonData))
	}
}

func assertDateJSONErr(t *testing.T, dateStr string) {
	var date Date
	err := json.Unmarshal([]byte(dateStr), &date)

	if err == nil {
		t.Fatalf(`Expected parsing %v to return error, but did not`, dateStr)
	}
}

func assertDateJSON(t *testing.T, dateJSON string, expected string) {
	var date Date
	err := json.Unmarshal([]byte(dateJSON), &date)
	if err != nil {
		t.Error(err)
	}

	expectedTime, err := time.Parse(time.RFC3339, expected)
	if err != nil {
		t.Error(err)
	}

	if date.Time != expectedTime {
		t.Fatalf("Expected %v, got %v", expectedTime, date.Time)
	}
}
