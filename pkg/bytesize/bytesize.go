// cmpout

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytesize

import (
	"fmt"
)

// ByteSize is a float64.
type ByteSize float64

const (
	_ = iota // ignore first value by assigning to blank identifier
	// KB is kilobytes.
	KB ByteSize = 1 << (10 * iota)
	// MB is megabytes.
	MB
	// GB is gigabytes
	GB
	// TB is terabytes
	TB
	// PB is petabytes.
	PB
	// EB is exabytes.
	EB
	// ZB is zettabytes.
	ZB
	// YB is yottabytes.
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2f YB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2f ZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2f EB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2f PB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2f TB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2f GB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2f MB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2f KB", b/KB)
	}
	return fmt.Sprintf("%.2f B", b)
}
