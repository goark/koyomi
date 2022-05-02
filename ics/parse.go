package ics

import (
	"bytes"
	"io"
	"regexp"

	"github.com/goark/errs"
)

const maxDataSize = 1024 * 1024 * 1024 //1GB

var (
	regEvents = regexp.MustCompile(`(BEGIN:VEVENT(.*\n)*?END:VEVENT\r?\n)`)
)

//Parse function parses ics data stream, and returns Event list.
func Parse(r io.Reader, max int64) ([]*Event, error) {
	if max <= 0 {
		max = maxDataSize
	}
	buf := &bytes.Buffer{}
	if _, err := io.CopyN(buf, r, max); err != nil {
		if !errs.Is(err, io.EOF) {
			return nil, errs.Wrap(err)
		}
	}
	return ParseString(buf.String())
}

//ParseString function parses ics data string, and returns Event list.
func ParseString(s string) ([]*Event, error) {
	events := regEvents.FindAllString(s, len(s))
	ical := parseCalendar(regEvents.ReplaceAllString(s, ""))
	evtList := []*Event{}
	for _, e := range events {
		evtList = append(evtList, parseEvent(e, ical))
	}
	return evtList, nil
}

/** These codes are forked form "github.com/PuloV/ics-golang" package. (licensed under MIT) */
