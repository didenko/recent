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
	"testing"
	"time"
)

func TestToString(t *testing.T) {
	ts := time.Date(2016, time.February, 29, 23, 59, 59, 999, time.UTC)
	p := New("/does/not/exist", ts)

	got := p.ToString("2006-01-02 15:04:05")
	want := "2016-02-29 23:59:59: /does/not/exist"

	if got != want {
		t.Errorf(
			"Failed to convert path with time into a string.\nExpected:\t%q\nReceived:\t%q",
			want,
			got,
		)
	}

	got = p.ToString("")
	want = "/does/not/exist"

	if got != want {
		t.Errorf(
			"Failed to convert path with time into a string.\nExpected:\t%q\nReceived:\t%q",
			want,
			got,
		)
	}

}

func ptsDump(pts PTslice) string {
	sig := ""
	for _, p := range pts {
		sig += " " + p.string
	}
	return sig
}

type ptUseCase = struct {
	pt   *PT
	dump string
}

var tn = time.Now()

var ptucs = []ptUseCase{
	{New("0", tn), " 0"},
	{New("1", tn), " 1 0"},
	{New("2", tn.Add(-2*time.Second)), " 1 0 2"},
	{New("3", tn.Add(-time.Second)), " 1 0 3 2"},
	{New("4", tn.Add(-2*time.Second)), " 1 0 3 4 2"},
	{New("5", tn.Add(-3*time.Second)), " 1 0 3 4 2 5"},
	{New("6", tn.Add(time.Second)), " 6 1 0 3 4 2 5"},
}

func TestInsert(t *testing.T) {
	pts := make(PTslice, 0, len(ptucs)+1)

	if len(pts) != 0 {
		t.Fatal("initial PTslice length is not 0")
	}

	for i, ptuc := range ptucs {
		pts = Insert(pts, ptuc.pt)

		if len(pts) != i+1 {
			t.Errorf(
				"Wrong PTslice length %d after iteration %d",
				len(pts), i)
		}

		got := ptsDump(pts)
		if got != ptuc.dump {
			t.Errorf(
				"Wrong PTslice strings after iteration %d:\n\tExpected:\t%q\n\tReceived:\t%q",
				i, ptuc.dump, got)
		}

	}
}
