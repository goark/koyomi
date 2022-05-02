package ics

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/goark/koyomi/ics/duration"
)

// Event is class of event in iCal calendar object.
type Event struct {
	Start     time.Time
	End       time.Time
	StartTZID string
	EndTZID   string
	Duration  time.Duration
	WholeDay  bool
	Calendar  *Calendar
}

// NewEvent returns nerw Event instance.
func NewEvent(ical *Calendar) *Event {
	return &Event{
		Calendar: ical,
	}
}

func parseEvent(data string, ical *Calendar) *Event {
	event := NewEvent(ical)
	event.Start, event.StartTZID, event.WholeDay = getEventStart(data)
	event.End, event.EndTZID, _ = getEventEnd(data)
	event.Duration = getEventDuration(data)
	if !event.WholeDay {
		if event.Duration != 0 {
			if event.End.IsZero() && !event.Start.IsZero() {
				event.End = event.Start.Add(event.Duration)
			}
		} else {
			event.Duration = event.End.Sub(event.Start)
		}
	}

	return event
}

const (
	icsDateTimeFormat   = "20060102T150405Z"
	eventWholeDayFormat = "20060102"
)

var (
	regDuration = regexp.MustCompile(`DURATION:.*?\n`)
)

func getEventStart(data string) (time.Time, string, bool) {
	return parseTimeField("DTSTART", data)
}

func getEventEnd(data string) (time.Time, string, bool) {
	return parseTimeField("DTEND", data)
}

func parseTimeField(fieldName string, eventData string) (time.Time, string, bool) {
	resultWholeDay := regexp.MustCompile(fmt.Sprintf(`%s;VALUE=DATE:.*?\n`, fieldName)).FindString(eventData)
	if len(resultWholeDay) > 0 {
		// whole day event
		modified := trimField(resultWholeDay, fmt.Sprintf("%s;VALUE=DATE:", fieldName))
		t, err := time.Parse(eventWholeDayFormat, modified)
		if err != nil {
			t = time.Time{}
		}
		return t, "", true
	} else {
		// event that has start hour and minute
		results := regexp.MustCompile(fmt.Sprintf(`%s(;TZID=(.*?))?(;VALUE=DATE-TIME)?:(.*?)\n`, fieldName)).FindStringSubmatch(eventData)
		if len(results) < 4 {
			return time.Time{}, "", false
		}
		dt := results[4]
		if !strings.HasSuffix(dt, "Z") {
			dt += "Z"
		}
		t, err := time.Parse(icsDateTimeFormat, dt)
		if err != nil {
			t = time.Time{}
		}
		return t, results[2], false
	}
}

func getEventDuration(data string) time.Duration {
	dur := trimField(regDuration.FindString(data), "DURATION:")
	parsed, err := duration.FromString(dur)
	if err != nil {
		return 0
	}
	return parsed.ToDuration()
}

/** These codes are forked form "github.com/PuloV/ics-golang" package. (licensed under MIT) */
