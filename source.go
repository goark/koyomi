package koyomi

import (
	ics "github.com/PuloV/ics-golang"
	"github.com/goark/errs"
)

//Source is information of data source for koyomi
type Source struct {
	cids    []CalendarID
	start   DateJp
	end     DateJp
	tempDir string //temporary directory for github.com/PuloV/ics-golang package
}

//optFunc is self-referential function for functional options pattern
type optFunc func(*Source)

// NewSource creates a new Source instance
func NewSource(opts ...optFunc) *Source {
	s := &Source{
		cids: []CalendarID{},
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

//WithCalendarID returns function for setting Reader
func WithCalendarID(cid ...CalendarID) optFunc {
	return func(s *Source) {
		s.cids = append(s.cids, cid...)
	}
}

//WithStartDate returns function for setting Reader
func WithStartDate(start DateJp) optFunc {
	return func(s *Source) {
		s.start = start
	}
}

//WithEndDate returns function for setting Reader
func WithEndDate(end DateJp) optFunc {
	return func(s *Source) {
		s.end = end
	}
}

//WithTempDir returns function for setting Reader
func WithTempDir(dir string) optFunc {
	return func(s *Source) {
		s.tempDir = dir
	}
}

//Get returns koyomi data from calendar dources
func (s *Source) Get() (*Koyomi, error) {
	if s == nil {
		return nil, errs.Wrap(ErrNullPointer)
	}
	if len(s.cids) == 0 {
		return nil, errs.Wrap(ErrNoData)
	}
	k := newKoyomi()
	if len(s.tempDir) > 0 {
		ics.FilePath = s.tempDir
	}
	for _, cid := range s.cids {
		es, err := getFrom(cid, s.start, s.end)
		if err != nil {
			return nil, errs.Wrap(err)
		}
		k.append(es...)
	}
	k.SortByDate()
	return k, nil
}

func getFrom(cid CalendarID, start, end DateJp) ([]Event, error) {
	url := cid.URL()
	if len(url) == 0 {
		return nil, errs.Wrap(ErrNoData, errs.WithContext("cid", int(cid)), errs.WithContext("start", start.String()), errs.WithContext("end", end.String()))
	}
	parser := ics.New()
	pch := parser.GetInputChan()
	pch <- url
	parser.Wait()

	calendars, err := parser.GetCalendars()
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("cid", int(cid)), errs.WithContext("start", start.String()), errs.WithContext("end", end.String()))
	}
	kevts := []Event{}
	for _, calendar := range calendars {
		for _, evt := range calendar.GetEvents() {
			e := Event{Date: NewDate(evt.GetStart()), Title: evt.GetSummary()}
			if boundaryIn(e, start, end) {
				kevts = append(kevts, e)
			}
		}
	}
	return kevts, nil
}

func boundaryIn(e Event, start, end DateJp) bool {
	if e.Date.IsZero() {
		return false
	}
	if !start.IsZero() && e.Date.Before(start) {
		return false
	}
	if !end.IsZero() && e.Date.After(end) {
		return false
	}
	return true
}

/* Copyright 2020-2022 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
