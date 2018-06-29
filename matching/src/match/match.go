package match

import (
	"image"
	"image/png"
	"image/jpeg"
	"os"
	"sort"
	"sync"
)

type MatchResult struct {
	name string
	pos int
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

func calcDiff(x uint32, y uint32) uint32 {
	if x < y {
	   return y - x 
	} else {
		return x - y
	}
}

func findMinDiff(src image.Image, target image.Image) int {
	bounds := src.Bounds()
	tBounds := target.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	tw, th := tBounds.Max.X, tBounds.Max.Y
	var threshold uint32 = 50000
	var minDiff uint32 = threshold
	minx := -1
	// miny := -1
	step := 50
	for x := 0; x < w - tw; x++ {
		for y := 0; y < h - th; y++ {
			var sumDiff uint32
			for tx := 0; tx < tw; tx = tx + step {
				for ty := 0; ty < th; ty = ty + step {
					color := src.At(x + tx, y + ty)
					r, g, b, _ := color.RGBA()
					tColor := target.At(tx, ty)
					tr, tg, tb, _ := tColor.RGBA()
					diff := calcDiff(r, tr) + calcDiff(g, tg) + calcDiff(b, tb)
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

func Match(f string) []string {
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

				minx := findMinDiff(src, tSrc)
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
