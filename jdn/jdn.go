package jdn

import (
	"math/big"
	"time"
)

// GetJD returns Julian Day from time.Time.
func GetJD(dt time.Time) *big.Rat {
	dt = dt.In(time.UTC)
	y := intRat(int64(dt.Year()))
	m := int64(dt.Month())
	d := int64(dt.Day())
	k := floorRat(quoInt(intRat(14-m), 12))
	jdn := floorRat(mulRat(addInt(subRat(y, k), 4800), fracInt(1461, 4)))
	jdn = addRat(jdn, floorRat(mulRat(addInt(mulInt(k, 12), m-2), fracInt(367, 12))))
	jdn = subRat(jdn, floorRat(mulRat(floorRat(quoInt(addInt(subRat(y, k), 4900), 100)), fracInt(3, 4))))
	jdn = addInt(jdn, d-32075)
	jdn = addRat(jdn, subRat(quoRat(addRat(intRat(int64(dt.Second()+dt.Minute()*60+dt.Hour()*3600)), fracInt(int64(dt.Nanosecond()), 999999999)), floatRat((24*time.Hour).Seconds())), floatRat(0.5)))
	return jdn
}

// GetJDN returns Julian Day Number from time.Time.
func GetJDN(dt time.Time) int64 {
	return floorRat(GetJD(dt)).Num().Int64()
}

// GetMJD returns Modified Julian Day from time.Time.
func GetMJD(dt time.Time) *big.Rat {
	return subRat(GetJD(dt), floatRat(2400000.5))
}

// GetMJDN returns Modified Julian Day Number from time.Time.
func GetMJDN(dt time.Time) int64 {
	return floorRat(GetMJD(dt)).Num().Int64()
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
