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
	"sort"
	"testing"

	"pgregory.net/rapid"
)

func TestBiggest(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.SliceOfN(rapid.Int(), 1, math.MaxInt).Draw(t, "list").([]int)
		b := biggest(s)
		sort.Ints(s)
		if b != s[len(s)-1] {
			t.Errorf("failed: %v", s)
		}
	})
}

func biggest(l []int) int {
	switch len(l) {
	case 0:
		return 0
	case 1:
		return l[0]
	default:
		tmp := l[0]
		for i := 1; i < len(l); i++ {
			if l[i] > tmp {
				tmp = l[i]
			}
		}
		return tmp
	}
}
