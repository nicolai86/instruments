// Package instruments allows you to send signals to Instruments.app to indicate interesting points or regions
package instruments

/*
#include <sys/kdebug_signpost.h>
*/
import "C"

// https://opensource.apple.com/source/xnu/xnu-3789.21.4/bsd/sys/kdebug_signpost.h.auto.html
const (
	// ColorBlue allows Instruments to highlight a point or region in blue
	ColorBlue int32 = iota
	// ColorBlue allows Instruments to highlight a point or region in green
	ColorGreen int32 = 1
	// ColorBlue allows Instruments to highlight a point or region in purple
	ColorPurple int32 = 2
	// ColorBlue allows Instruments to highlight a point or region in orange
	ColorOrange int32 = 3
	// ColorBlue allows Instruments to highlight a point or region in red
	ColorRed int32 = 4
)

// Signpost creates a single point of interest
func Signpost(action int32) {
	SignpostWithArguments(action, 0, 0, 0, 0)
}

// SignpostWithArguments creates a single point of interest, allowing you to specifiy all arguments
func SignpostWithArguments(action, arg1, arg2, arg3, arg4 int32) {
	C.kdebug_signpost(
		C.uint32_t(action),
		C.uintptr_t(arg1),
		C.uintptr_t(arg2),
		C.uintptr_t(arg3),
		C.uintptr_t(arg4),
	)
}

// Region describes an active region
type Region struct {
	Action int32
	Arg1   int32
	Arg2   int32
	Arg3   int32
	Arg4   int32
}

// Start marks the beginning of a Region
func Start(action int32) Region {
	return StartWithArguments(action, 0, 0, 0, 0)
}

// StartWithArguments marks the beginning of a Region, allowing you to specifiy all arguments
func StartWithArguments(action, arg1, arg2, arg3, arg4 int32) Region {
	r := Region{action, arg1, arg2, arg3, arg4}
	C.kdebug_signpost_start(
		C.uint32_t(action),
		C.uintptr_t(arg1),
		C.uintptr_t(arg2),
		C.uintptr_t(arg3),
		C.uintptr_t(arg4),
	)
	return r
}

// End marks the ending of a Region
func (r Region) End() {
	C.kdebug_signpost_end(
		C.uint32_t(r.Action),
		C.uintptr_t(r.Arg1),
		C.uintptr_t(r.Arg2),
		C.uintptr_t(r.Arg3),
		C.uintptr_t(r.Arg4),
	)
}
