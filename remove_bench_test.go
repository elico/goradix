// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "testing"

// ----------------------- Benchmarks ------------------------ //

func BenchmarkRemove(b *testing.B) {
	rx := New()
	insertDataBytes(rx, sampleData3)
	sd3 := sampleData3()
	sdLen := len(sd3)
	tn := 0

	for i := 0; i < b.N; i++ {
		if tn == sdLen {
			rx = New()
			insertDataBytes(rx, sampleData3)
		}

		rx.RemoveBytes(randomBytes(sd3))

		tn++
	}
}
