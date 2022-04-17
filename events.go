package koyomi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/goark/errs"
)

//Event is koyomi event data
type Event struct {
	Date  DateJp
	Title string
}

//Koyomi is array of Event
type Koyomi struct {
	events []Event
}

//newKoyomi createw Koyomi instance
func newKoyomi() *Koyomi {
	return &Koyomi{events: make([]Event, 0)}
}

//Events returns event array
func (k *Koyomi) Events() []Event {
	if k == nil {
		return []Event{}
	}
	return k.events
}

//SortByDate sorts event data by date
func (k *Koyomi) SortByDate() {
	if k == nil || len(k.events) <= 1 {
		return
	}
	sort.SliceStable(k.events, func(i, j int) bool {
		return k.events[i].Date.Before(k.events[j].Date)
	})
}

//Add adds other Kyomoi instance
func (k *Koyomi) Add(kk *Koyomi) {
	if kk == nil {
		return
	}
	k.append(kk.events...)
	k.SortByDate()
}

func (k *Koyomi) append(e ...Event) {
	if k == nil {
		return
	}
	k.events = append(k.events, e...)
}

func (k *Koyomi) EncodeJSON() ([]byte, error) {
	if k == nil || len(k.events) == 0 {
		return nil, errs.Wrap(ErrNoData)
	}
	return json.Marshal(k.events)
}

func (k *Koyomi) EncodeCSV() ([]byte, error) {
	if k == nil || len(k.events) == 0 {
		return nil, errs.Wrap(ErrNoData)
	}
	buf := &bytes.Buffer{}
	_, err := buf.WriteString(`"Date","Title"` + "\n")
	if err != nil {
		return nil, errs.Wrap(err)
	}
	for _, e := range k.events {
		_, err := buf.WriteString(fmt.Sprintf("%s,%s\n", strconv.Quote(e.Date.String()), strconv.Quote(e.Title)))
		if err != nil {
			return nil, errs.Wrap(err)
		}
	}
	return buf.Bytes(), nil
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
