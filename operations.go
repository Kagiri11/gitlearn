package operations

func TriArea(b, h float32) float32 {

	var area float32
	area = 0.5 * (b * h)
	return area
}

func RectArea(l, w float32) float32 {
	area := l * w
	return area
}

func CircleArea(r float32) float32 {
	area := 3.14 * (r * r)
	return area
}
