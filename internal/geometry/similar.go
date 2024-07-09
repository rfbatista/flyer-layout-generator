package geometry

import (
	"algvisual/internal/entities"
	"math"
)

func IsContainerSimilar(c1, c2 entities.Container, limit int) bool {
	upx := float64(c1.UpperLeft.X) - float64(c2.UpperLeft.X)
	upy := float64(c1.UpperLeft.Y) - float64(c2.UpperLeft.Y)
	dwx := float64(c1.DownRight.X) - float64(c2.DownRight.X)
	dwy := float64(c1.DownRight.Y) - float64(c2.DownRight.Y)
	return math.Abs(upx) <= float64(limit) &&
		math.Abs(upy) <= float64(limit) &&
		math.Abs(dwx) <= float64(limit) &&
		math.Abs(dwy) <= float64(limit)
}
