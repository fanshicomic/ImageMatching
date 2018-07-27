package match

import (
	"image"
	"image/draw"
)

type MatchResult struct {
	name string
	pos  int
}

type ByPos []MatchResult

func (r ByPos) Len() int {
	return len(r)
}

func (r ByPos) Less(i, j int) bool {
	return r[i].pos < r[j].pos
}

func (r ByPos) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func checkErr(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func convertToGray(src image.Image) image.Image {
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	result := image.NewGray(image.Rect(0, 0, w, h))
	draw.Draw(result, result.Bounds(), src, bounds.Min, draw.Over)

	return result
}
