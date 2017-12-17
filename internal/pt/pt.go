// Copyright © 2017 Vlad Didenko <business@didenko.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pt

import (
	"fmt"
	"time"
)

// PT is a path with wime struct
type PT struct {
	string
	Time time.Time
}

// ToString formats the Path with Time struct for output
func (p *PT) ToString(format string) string {
	if len(format) == 0 {
		return p.string
	}
	return fmt.Sprintf("%s: %s", p.Time.Format(format), p.string)
}

// New is a factory method for Path with Time
func New(p string, t time.Time) *PT {
	return &PT{p, t}
}

// PTslice is a slice of Paths with Times, sortable by Time
type PTslice []*PT

func (pts PTslice) locate(t time.Time) int {
	for i := range pts {
		if !pts[i].Time.After(t) {
			return i
		}
	}
	return len(pts)
}

// Insert places the Path with Time struct in a first place
// suitable to keep timestamps in descending order
func Insert(pts PTslice, pt *PT) PTslice {
	i := pts.locate(pt.Time)
	npts := append(pts, &PT{})
	copy(npts[i+1:], npts[i:])
	npts[i] = pt
	return npts
}
