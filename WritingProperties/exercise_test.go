// Copyright 2022 Yoshi Yamaguchi
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

package writing_properties

import (
	"math"
	"testing"

	"pgregory.net/rapid"
)

// NOTE: setting upper bound to math.MaxInt will cause out of memory error.
func TestIncrements(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.IntRange(math.MinInt, 100000).Draw(t, "start").(int)
		c := rapid.IntRange(1, 100000).Draw(t, "count").(int)
		l := makeRange(s, c)
		if len(l) != c+1 || !increments(l) {
			t.Errorf("failed: %v", l)
		}
	})
}

func makeRange(start, count int) []int {
	if count <= 0 {
		return []int{}
	}
	l := make([]int, count+1)
	for i := 0; i < count+1; i++ {
		l[i] = start + i
	}
	return l
}

func increments(l []int) bool {
	if len(l) == 1 {
		return true
	}
	for i, j := 0, 1; j < len(l); i, j = i+1, j+1 {
		if l[i] >= l[j] {
			return false
		}
	}
	return true
}
