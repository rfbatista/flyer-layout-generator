package geometry

import (
	"algvisual/internal/entities"
	"math"
)

func IsContainerSimilar(c1, c2 entities.Container, limit int) bool {
	return math.Abs(float64(c1.UpperLeft.X)-float64(c2.UpperLeft.X)) <= float64(limit) &&
		math.Abs(float64(c1.UpperLeft.Y)-float64(c2.UpperLeft.Y)) <= float64(limit) &&
		math.Abs(float64(c1.DownRight.X)-float64(c2.DownRight.X)) <= float64(limit) &&
		math.Abs(float64(c1.DownRight.Y)-float64(c2.DownRight.Y)) <= float64(limit)
}
