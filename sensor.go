package sdl

import (
	"math"
)

type SensorID int32

type SensorType uint32

const (
	SensorUnknown SensorType = iota
	SensorAccelerometer
	SensorGyroscope
	SensorInvalid SensorType = math.MaxUint32
)
