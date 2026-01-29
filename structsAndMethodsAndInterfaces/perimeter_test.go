package structsandmethodsandinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct { // 匿名構造体の定義
		name string
		shape Shape
		want float64
	}{ // スライス内部のデータ=テストケース
		// struct{a, b}としているのと同じ。匿名構造体なので構造体名が省略されている。
		{name: "rectangle", shape: Rectangle{12.0, 6.0}, want: 72.0},
		{name: "circle", shape: Circle{10.0}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{12.0, 6.0}, want: 36.0},
	}

	// テストケースごとにテスト実行
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g in %v", got, tt.want, tt.name)
			}
		})
	}
}
