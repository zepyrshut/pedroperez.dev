package datesnightmare

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type People struct {
	Name     string    `json:"name"`
	DateTime time.Time `json:"date"`
}

func datesNightmare(file *os.File) []People {
	var people []People
	_ = json.NewDecoder(file).Decode(&people)
	return people
}

func (p *People) UnmarshalJSON(data []byte) error {
	var err error
	type Alias People
	aux := &struct {
		Date string `json:"date"`
		Time string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	p.DateTime, err = DatesParser(aux.Date, aux.Time)
	return nil
}

func DatesParser(dateStr string, timeStr string) (time.Time, error) {
	var dateString string
	if dateStr == "" {
		dateString = "01/01/0001"
	} else {
		dateString = dateStr
	}

	var timeString string
	if timeStr == "" {
		timeString = "00:00"
	} else {
		timeString = timeStr
	}

	_, err := time.Parse("02/01/2006", dateString)
	if err != nil {
		return time.Time{}, err
	}

	dateTimeString := fmt.Sprintf("%s %s", dateString, timeString)
	return time.Parse("02/01/2006 15:04", dateTimeString)
}
