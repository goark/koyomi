package koyomi

import "net/url"

type CalendarID int

const (
	Holiday   CalendarID = iota + 1 //国民の祝日および休日
	MoonPhase                       //朔弦望
	SolarTerm                       //二十四節気・雑節
	Eclipse                         //日食・月食・日面経過
	Planet                          //惑星現象
)

var cidMap = map[CalendarID]string{
	Holiday:   "2bk907eqjut8imoorgq1qa4olc@group.calendar.google.com", //国民の祝日および休日
	MoonPhase: "mr1q70hu2iacu62adntahc69q0@group.calendar.google.com", //朔弦望
	SolarTerm: "2i7smciu430uh0mv3i0qmd8iuk@group.calendar.google.com", //二十四節気・雑節
	Eclipse:   "9lpmd80aki4edordf25nqjnln4@group.calendar.google.com", //日食・月食・日面経過
	Planet:    "fsj78svf2km2stokku3r2ajuts@group.calendar.google.com", //惑星現象
}

//String is Stringer of CalendarID
func (cid CalendarID) String() string {
	if s, ok := cidMap[cid]; ok {
		return s
	}
	return ""
}

//URL returns URL string from CalendarID
func (cid CalendarID) URL() string {
	id := cid.String()
	if len(id) == 0 {
		return ""
	}
	return "https://calendar.google.com/calendar/ical/" + url.PathEscape(id) + "/public/basic.ics"
}

/* Copyright 2020 Spiegel
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
