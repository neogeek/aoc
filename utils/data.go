package utils

type Number interface {
	int64 | float64
}

type Range struct {
	Start uint64
	End   uint64
}

type BoundingBox struct {
	MinX float64
	MaxX float64
	MinY float64
	MaxY float64
}

func (a BoundingBox) Overlaps(b BoundingBox) bool {
	x_overlap := a.MinX < b.MaxX && a.MaxX > b.MinX
	y_overlap := a.MinY < b.MaxY && a.MaxY > b.MinY
	return x_overlap && y_overlap
}

type Vector2 struct {
	X float64
	Y float64
}

func (vector Vector2) Add(other Vector2) Vector2 {
	return Vector2{
		X: vector.X + other.X,
		Y: vector.Y + other.Y,
	}
}

func (vector Vector2) Subtract(other Vector2) Vector2 {
	return Vector2{
		X: vector.X - other.X,
		Y: vector.Y - other.Y,
	}
}

func (vector Vector2) Dot(other Vector2) float64 {
	return vector.X*other.X + vector.Y*other.Y
}

func (vector Vector2) Cross(other Vector2) float64 {
	return vector.X*other.Y - vector.Y*other.X
}

type Vector3 struct {
	X float64
	Y float64
	Z float64
}
