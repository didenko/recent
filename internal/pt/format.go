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

// PickFormat returns the corresponding timestamp format line
func PickFormat(withdate, withtime, withmill bool) string {
	var f string
	if withdate {
		f += "2006-01-02"
	}
	if withtime {
		if len(f) > 0 {
			f += " "
		}
		f += "15:04:05"

		if withmill {
			f += ".000"
		}
	}
	return f
}
