package page25

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i Image) At(x, y int) color.Color {
	// 예시: 간단한 패턴 생성
	v := uint8((x + y) / 2) // v의 값을 적절하게 계산
	return color.RGBA{v, v, 255, 255}
}

func Images() {
	m := Image{Width: 256, Height: 256} // 너비와 높이를 설정
	pic.ShowImage(m)
}
