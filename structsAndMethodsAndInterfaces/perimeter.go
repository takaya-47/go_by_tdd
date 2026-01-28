package structsandmethodsandinterfaces

type Rectangle struct {
	width, height float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}

func Area(r Rectangle) float64 {
	return r.width * r.height
}