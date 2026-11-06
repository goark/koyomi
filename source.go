package koyomi

import (
	"context"
	"sort"
	"time"

	"github.com/goark/concal"
	"github.com/goark/errs"
	"github.com/goark/koyomi/value"
)

// Source is information of data source for koyomi
type Source struct {
	cids  []CalendarID
	start value.DateJp
	end   value.DateJp
}

// optFunc is self-referential function for functional options pattern
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

// WithCalendarID returns function for setting Reader
func WithCalendarID(cid ...CalendarID) optFunc {
	return func(s *Source) {
		s.cids = append(s.cids, cid...)
	}
}

// WithStartDate returns function for setting Reader
func WithStartDate(start value.DateJp) optFunc {
	return func(s *Source) {
		s.start = start
	}
}

// WithEndDate returns function for setting Reader
func WithEndDate(end value.DateJp) optFunc {
	return func(s *Source) {
		s.end = end
	}
}

// Get returns koyomi data from calendar dources
func (s *Source) Get(ctx context.Context) (*Koyomi, error) {
	if s == nil {
		return nil, errs.Wrap(ErrNullPointer)
	}
	if len(s.cids) == 0 {
		return nil, errs.Wrap(ErrNoData)
	}
	k := newKoyomi()
	es, err := getFrom(ctx, s.cids, s.start, s.end)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	k.append(es...)
	k.SortByDate()
	return k, nil
}

func getFrom(ctx context.Context, cids []CalendarID, start, end value.DateJp) (evts []Event, retErr error) {
	parser := concal.NewParser(
		concal.WithStartTime((start.Time).Add(1*time.Second)), // add 1 second to start time to avoid boundary issues
		concal.WithEndTime((end.Time).Add(1*time.Second)),     // add 1 second to end time to avoid boundary issues
		concal.WithLocationAllDayEvent(value.JST),
	)
	cparser := concal.NewConcurrent(ctx, parser)
	defer func() {
		if cerr := cparser.Close(); cerr != nil {
			retErr = errs.Join(retErr, errs.Wrap(cerr))
		}
	}()

	for _, cid := range cids {
		url := cid.URL()
		if len(url) == 0 {
			return nil, errs.Wrap(ErrNoData, errs.WithContext("cid", int(cid)), errs.WithContext("start", start.String()), errs.WithContext("end", end.String()))
		}
		if err := cparser.Submit(url); err != nil {
			return nil, errs.Wrap(err, errs.WithContext("url", url))
		}
	}
	go cparser.Wait() // start waiting in a separate goroutine

	kevts := []Event{}
	for evt := range cparser.Events() {
		kevts = append(kevts, Event{Date: value.NewDate(*evt.Start), Title: evt.Summary})
	}
	sort.Slice(kevts, func(i, j int) bool {
		return kevts[i].Date.Before(kevts[j].Date) && kevts[i].Title < kevts[j].Title
	})
	return kevts, nil
}

/* Copyright 2020-2026 Spiegel
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
