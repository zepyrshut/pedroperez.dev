package datesnightmare

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_DatesNightmare(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  []People
	}{
		{"no file", "", nil},
		{"some file", "people_test.json", []People{
			{Name: "Chip", DateTime: time.Date(2022, 07, 01, 12, 22, 0, 0, time.UTC)},
			{Name: "Rosmunda", DateTime: time.Date(2022, 12, 03, 0, 0, 0, 0, time.UTC)},
			{Name: "Kimmi", DateTime: time.Date(0001, 01, 01, 12, 28, 0, 0, time.UTC)},
			{Name: "Wildon", DateTime: time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC)},
		}},
		{"with error", "people_errors_test.json", []People{
			{Name: "Chip", DateTime: time.Date(2022, 07, 01, 12, 22, 0, 0, time.UTC)},
			{Name: "Rosmunda", DateTime: time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC)},
			{Name: "Kimmi", DateTime: time.Date(0001, 01, 01, 12, 28, 0, 0, time.UTC)},
			{Name: "Wildon", DateTime: time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC)},
		}},
	}

	for _, tt := range tests {
		file, _ := os.Open(tt.given)

		got := datesNightmare(file)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. datesNightmare() = %v, want %v", tt.name, got, tt.want)
		}

	}
}
