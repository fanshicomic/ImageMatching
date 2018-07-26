package match

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
