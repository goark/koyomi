package jdn

import (
	"math/big"
	"testing"
	"time"
)

func TestGetJDN(t *testing.T) {
	jst := time.FixedZone("JST", int((9 * time.Hour).Seconds())) // Japan standard Time
	testCases := []struct {
		inp    time.Time
		outp1  *big.Rat
		outp2  int64
		outpDt time.Time
		outp3  *big.Rat
		outp4  int64
	}{
		{inp: time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC), outp1: floatRat(2457023.5), outp2: 2457023, outpDt: time.Date(2015, time.January, 0, 12, 0, 0, 0, time.UTC), outp3: floatRat(57023.0), outp4: 57023},
		{inp: time.Date(2022, time.January, 1, 0, 0, 0, 0, jst), outp1: floatRat(2459580.125), outp2: 2459580, outpDt: time.Date(2022, time.January, 0, 12, 0, 0, 0, time.UTC), outp3: floatRat(59579.625), outp4: 59579},
		{inp: time.Date(2023, time.February, 24, 12, 0, 0, 0, time.UTC), outp1: floatRat(2460000.0), outp2: 2460000, outpDt: time.Date(2023, time.February, 24, 12, 0, 0, 0, time.UTC), outp3: floatRat(59999.5), outp4: 59999},
	}
	for _, tc := range testCases {
		jd := GetJD(tc.inp)
		if jd.Cmp(tc.outp1) != 0 {
			t.Errorf("GetJD(%v) is %v, want %v.", tc.inp, jd.FloatString(5), tc.outp1.FloatString(5))
		}
		fjd, _ := jd.Float64()
		dt := FromJD(fjd)
		if !dt.Equal(tc.inp) {
			t.Errorf("FromJD(%v) is %v, want %v.", fjd, dt, tc.inp)
		}
		jn := GetJDN(tc.inp)
		if jn != tc.outp2 {
			t.Errorf("GetJDN(%v) is %v, want %v.", tc.inp, jn, tc.outp2)
		}
		dt = FromJDN(jn)
		if !dt.Equal(tc.outpDt) {
			t.Errorf("FromJDN(%v) is %v, want %v.", jn, dt, tc.outpDt)
		}
		mjd := GetMJD(tc.inp)
		if mjd.Cmp(tc.outp3) != 0 {
			t.Errorf("GetMJD(%v) is %v, want %v.", tc.inp, mjd.FloatString(5), tc.outp3.FloatString(5))
		}
		mjdn := GetMJDN(tc.inp)
		if mjdn != tc.outp4 {
			t.Errorf("GetMJDN(%v) is %v, want %v.", tc.inp, mjdn, tc.outp4)
		}
		fmjd, _ := mjd.Float64()
		dt = FromMJD(fmjd)
		if !dt.Equal(tc.inp) {
			t.Errorf("FromMJD(%v) is %v, want %v.", fjd, dt, tc.inp)
		}
	}
}

func TestFloorRat(t *testing.T) {
	testCases := []struct {
		inp  float64
		outp float64
	}{
		{inp: 1.1, outp: 1},
		{inp: 1.0, outp: 1},
		{inp: 0.9, outp: 0},
		{inp: 0.1, outp: 0},
		{inp: 0.0, outp: 0},
		{inp: -0.1, outp: -1},
		{inp: -0.9, outp: -1},
		{inp: -1.0, outp: -1},
		{inp: -1.1, outp: -2},
	}
	for _, tc := range testCases {
		f := floorRat(floatRat(tc.inp))
		if ff, _ := f.Float64(); ff != tc.outp {
			t.Errorf("floorRat(%v) is %v, want %v.", tc.inp, f, tc.outp)
		}
	}
}

/* Copyright 2022 Spiegel
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
