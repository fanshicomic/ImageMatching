package match

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"sort"
	"sync"
)

func calcDiffV1(x uint32, y uint32) uint32 {
	if x < y {
		return y - x
	} else {
		return x - y
	}
}

func findMinDiffV1(src image.Image, target image.Image) int {
	bounds := src.Bounds()
	tBounds := target.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	tw, th := tBounds.Max.X, tBounds.Max.Y
	var threshold uint32 = 50000
	var minDiff uint32 = threshold
	minx := -1
	// miny := -1
	step := 50
	for x := 0; x < w-tw; x++ {
		for y := 0; y < h-th; y++ {
			var sumDiff uint32
			for tx := 0; tx < tw; tx = tx + step {
				for ty := 0; ty < th; ty = ty + step {
					color := src.At(x+tx, y+ty)
					r, g, b, _ := color.RGBA()
					tColor := target.At(tx, ty)
					tr, tg, tb, _ := tColor.RGBA()
					diff := calcDiffV1(r, tr) + calcDiffV1(g, tg) + calcDiffV1(b, tb)
					sumDiff += diff
				}
			}
			if minDiff > sumDiff {
				minDiff = sumDiff
				minx = x
				// miny = y
			}
		}
	}
	if minDiff < threshold {
		return minx
	} else {
		return -1
	}
}

func MatchV1(f string) []string {
	result := []MatchResult{}
	infile, err := os.Open(f)
	checkErr(err)
	defer infile.Close()

	src, err := jpeg.Decode(infile)
	checkErr(err)

	targetFolder := "../shishen/"
	dir, err := os.Open(targetFolder)
	checkErr(err)
	defer dir.Close()

	var wg sync.WaitGroup
	list, _ := dir.Readdirnames(0)
	for _, name := range list {
		wg.Add(1)
		go func(name string) {
			cursor := len(name) - 4
			if len(name) > 4 && name[cursor:] == ".png" {
				target := targetFolder + name
				tInfile, err := os.Open(target)
				checkErr(err)
				tSrc, err := png.Decode(tInfile)
				checkErr(err)

				minx := findMinDiffV1(src, tSrc)
				if minx > 0 {
					result = append(result, MatchResult{name[:cursor], minx})
				}
			}
			wg.Done()
		}(name)
	}
	wg.Wait()

	lst := make([]string, 10)
	sort.Sort(ByPos(result))
	for _, r := range result {
		lst = append(lst, r.name)
	}

	return lst
}
