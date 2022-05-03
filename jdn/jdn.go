package jdn

import (
	"math/big"
	"time"
)

// GetJD returns Julian Date from time.Time.
func GetJD(dt time.Time) *big.Rat {
	dt = dt.In(time.UTC)
	y := intRat(int64(dt.Year()))
	m := int64(dt.Month())
	d := int64(dt.Day())
	k := floorRat(quoInt(intRat(14-m), 12))
	j := floorRat(mulRat(addInt(subRat(y, k), 4800), fracInt(1461, 4)))
	j = addRat(j, floorRat(mulRat(addInt(mulInt(k, 12), m-2), fracInt(367, 12))))
	j = subRat(j, floorRat(mulRat(floorRat(quoInt(addInt(subRat(y, k), 4900), 100)), fracInt(3, 4))))
	j = addInt(j, d-32075)
	j = addRat(j, subRat(quoRat(addRat(intRat(int64(dt.Second()+dt.Minute()*60+dt.Hour()*3600)), fracInt(int64(dt.Nanosecond()), 999999999)), floatRat((24*time.Hour).Seconds())), floatRat(0.5)))
	return j
}

// GetJDN returns Julian Day Number from time.Time.
func GetJDN(dt time.Time) int64 {
	return floorRat(GetJD(dt)).Num().Int64()
}

// GetMJD returns Modified Julian Date from time.Time.
func GetMJD(dt time.Time) *big.Rat {
	return subRat(GetJD(dt), floatRat(2400000.5))
}

// GetMJDN returns Modified Julian Day Number from time.Time.
func GetMJDN(dt time.Time) int64 {
	return floorRat(GetMJD(dt)).Num().Int64()
}

// FromJDN returns time.Time instance form Julian Day Number.
func FromJDN(jdnum int64) time.Time {
	l := intRat(jdnum + 68569)
	n := floorRat(mulInt(quoInt(l, 146097), 4))
	l = subRat(l, floorRat(quoInt(addInt(mulInt(n, 146097), 3), 4)))
	i := floorRat(quoInt(mulInt(addInt(l, 1), 4000), 1461001))
	l = addInt(subRat(l, floorRat(quoInt(mulInt(i, 1461), 4))), 31)
	j := floorRat(quoInt(mulInt(l, 80), 2447))
	day := subRat(l, floorRat(quoInt(mulInt(j, 2447), 80)))
	l = floorRat(quoInt(j, 11))
	month := subRat(addInt(j, 2), mulInt(l, 12))
	year := addRat(mulInt(addInt(n, -49), 100), addRat(i, l))
	return time.Date(int(year.Num().Int64()), time.Month(int(month.Num().Int64())), int(day.Num().Int64()), 12, 0, 0, 0, time.UTC)
}

// FromJD returns time.Time instance form Julian Date.
func FromJD(jd float64) time.Time {
	jdnum := int64(jd)
	dt := FromJDN(jdnum)
	return dt.Add(time.Duration((jd - float64(jdnum)) * float64(24*time.Hour)))
}

// FromJD returns time.Time instance form Julian Date.
func FromMJD(mjd float64) time.Time {
	return FromJD(mjd + 2400000.5)
}

func intRat(x int64) *big.Rat {
	return fracInt(x, 1)
}

func floatRat(x float64) *big.Rat {
	return (&big.Rat{}).SetFloat64(x)
}

func fracInt(x, y int64) *big.Rat {
	return big.NewRat(x, y)
}

func addInt(x *big.Rat, y int64) *big.Rat {
	return addRat(x, intRat(y))
}

// func subInt(x *big.Rat, y int64) *big.Rat {
// 	return subRat(x, intRat(y))
// }

func mulInt(x *big.Rat, y int64) *big.Rat {
	return mulRat(x, intRat(y))
}

func quoInt(x *big.Rat, y int64) *big.Rat {
	return quoRat(x, intRat(y))
}

func addRat(x, y *big.Rat) *big.Rat {
	return (&big.Rat{}).Add(x, y)
}

func subRat(x, y *big.Rat) *big.Rat {
	return (&big.Rat{}).Sub(x, y)
}

func mulRat(x, y *big.Rat) *big.Rat {
	return (&big.Rat{}).Mul(x, y)
}

func quoRat(x, y *big.Rat) *big.Rat {
	return (&big.Rat{}).Quo(x, y)
}

func floorRat(n *big.Rat) *big.Rat {
	return (&big.Rat{}).SetInt((&big.Int{}).Div(n.Num(), n.Denom()))
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
