package entities

type LayoutHistory struct {
	Versions []Layout
	Last     Layout
}

func (l *LayoutHistory) Push(ll Layout) {
	l.Versions = append(l.Versions, ll)
	l.Last = ll
}
