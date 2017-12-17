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

import "testing"

// 	"2006-01-02 15:04:05.000",

var formatPicks = []struct {
	d, t, m bool
	format  string
}{
	{false, false, false, ""},
	{false, false, true, ""},
	{false, true, false, "15:04:05"},
	{false, true, true, "15:04:05.000"},
	{true, false, false, "2006-01-02"},
	{true, false, true, "2006-01-02"},
	{true, true, false, "2006-01-02 15:04:05"},
	{true, true, true, "2006-01-02 15:04:05.000"},
}

func TestPickFormat(t *testing.T) {
	for i, fp := range formatPicks {
		if f := PickFormat(fp.d, fp.t, fp.m); f != fp.format {
			t.Errorf("Case %d produced the wrong format line: %q", i, f)
		}
	}
}
