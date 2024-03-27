package duration

import (
	"testing"
)

func TestNewDuration(t *testing.T) {
	d := NewDuration()
	if d.Hours() != 0 {
		t.Errorf("Expected hours to be 0, got %d", d.Hours())
	}
	if d.Minutes() != 0 {
		t.Errorf("Expected minutes to be 0, got %d", d.Minutes())
	}
	if d.Seconds() != 1 {
		t.Errorf("Expected seconds to be 1, got %d", d.Seconds())
	}
}

func TestSetDuration(t *testing.T) {
	d := NewDuration()
	params := []string{"2", "30", "45"}
	d.SetDuration(params)
	if d.Hours() != 2 {
		t.Errorf("Expected hours to be 2, got %d", d.Hours())
	}
	if d.Minutes() != 30 {
		t.Errorf("Expected minutes to be 30, got %d", d.Minutes())
	}
	if d.Seconds() != 45 {
		t.Errorf("Expected seconds to be 45, got %d", d.Seconds())
	}
}

func TestParseTime(t *testing.T) {
	s := "10"
	expected := 10
	result := ParseTime(s)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

}
