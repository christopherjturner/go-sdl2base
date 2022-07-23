package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

func ClampI32(i, min, max int32) int32 {
	if i < min {
		return min
	}
	if i > max {
		return max
	}
	return i
}

func DistSdlPoint(p1, p2 sdl.Point) float64 {
	return math.Sqrt(float64(((p2.X - p1.X) * (p2.X - p1.X)) +
		((p2.Y - p1.Y) * (p2.Y - p1.Y))))
}
