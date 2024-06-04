package entities

import (
	"math"

	"github.com/google/uuid"
)

func NewCell(x, y, xx, yy int32) *Cell {
	return &Cell{
		ID:  uuid.New(),
		Xi:  x,
		Yi:  y,
		Xii: xx,
		Yii: yy,
	}
}

type Cell struct {
	ID  uuid.UUID `json:"id"`
	Xi  int32     `json:"xi"`
	Xii int32     `json:"xii"`
	Yi  int32     `json:"yi"`
	Yii int32     `json:"yii"`
}

func (r *Cell) Width() int32 {
	return r.Xii - r.Xi
}

func (r *Cell) Height() int32 {
	return r.Yii - r.Yi
}

func (r *Cell) DistanceToEdge(p Point) int32 {
	smallerDistance := int32(math.Abs(float64(r.Xi - p.X)))
	if smallerDistance < r.Xii-p.X {
		smallerDistance = r.Xii - p.X
	}
	if smallerDistance < r.Yi-p.Y {
		smallerDistance = r.Yi - p.Y
	}
	if smallerDistance < r.Yii-p.Y {
		smallerDistance = r.Yii - p.Y
	}
	return smallerDistance
}
