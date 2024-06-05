package entities

import (
	"math"

	"github.com/google/uuid"
)

func NewCell(x, y, xx, yy int32) *GridCell {
	return &GridCell{
		ID:  uuid.New(),
		Xi:  x,
		Yi:  y,
		Xii: xx,
		Yii: yy,
	}
}

type GridCell struct {
	ID        uuid.UUID `json:"id"`
	Xi        int32     `json:"xi"`
	Xii       int32     `json:"xii"`
	Yi        int32     `json:"yi"`
	Yii       int32     `json:"yii"`
	isOcupied bool
	whoIsIn   []int32
}

func (r *GridCell) IsOcupied() bool {
	return r.isOcupied
}

func (r *GridCell) Ocupy(id int32) {
	r.isOcupied = true
	r.whoIsIn = append(r.whoIsIn, id)
}

func (r *GridCell) Width() int32 {
	return r.Xii - r.Xi
}

func (r *GridCell) Height() int32 {
	return r.Yii - r.Yi
}

// Find if a point is inside the cell
func (r *GridCell) IsIn(p Point) bool {
	return r.Xi <= p.X && r.Xii >= p.X && r.Yi <= p.Y && r.Yii >= p.Y
}

// Find the id is in the cell
func (r *GridCell) IsIdIn(id int32) bool {
	for _, i := range r.whoIsIn {
		if i == id {
			return true
		}
	}
	return false
}

func (r *GridCell) DistanceToEdge(p Point) int32 {
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
