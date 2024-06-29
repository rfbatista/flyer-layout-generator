package project

type request struct {
	DesignID int32 `param:"design_id"`
	LayoutID int32 `param:"layout_id"`
}

func Props() {
}
