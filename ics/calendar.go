package ics

import (
	"regexp"
	"time"
)

// Calendar is class of iCal calendar object.
type Calendar struct {
	Name        string
	Description string
	Timezone    *time.Location
	Version     string
}

// NewCalendar returns new Calendar instance.
func NewCalendar() *Calendar {
	return &Calendar{}
}

var (
	regICalName     = regexp.MustCompile(`X-WR-CALNAME:.*?\n`)
	regICalDesc     = regexp.MustCompile(`X-WR-CALDESC:.*?\n`)
	regICalTimeZone = regexp.MustCompile(`X-WR-TIMEZONE:.*?\n`)
	regICalVersion  = regexp.MustCompile(`VERSION:.*?\n`)
)

func parseCalendar(data string) *Calendar {
	ical := NewCalendar()
	ical.Name = getICalName(data)
	ical.Description = getICalDesc(data)
	ical.Timezone = getICalTimezone(data)
	ical.Version = getICalVersion(data)
	return ical
}

func getICalName(s string) string {
	return trimField(regICalName.FindString(s), "X-WR-CALNAME:")
}

func getICalDesc(s string) string {
	return trimField(regICalDesc.FindString(s), "X-WR-CALDESC:")
}

func getICalTimezone(s string) *time.Location {
	ls := trimField(regICalTimeZone.FindString(s), "X-WR-TIMEZONE:")
	loc, err := time.LoadLocation(ls)
	if err != nil {
		return time.UTC
	}
	return loc
}

func getICalVersion(s string) string {
	return trimField(regICalVersion.FindString(s), "VERSION:")
}

/** These codes are forked form "github.com/PuloV/ics-golang" package. (licensed under MIT) */
