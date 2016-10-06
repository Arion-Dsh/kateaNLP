package summary

import (
	"math"
	"sort"

	"github.com/Arion-Dsh/kateaNLP/sim"
)

type textRankI struct {
	docs    [][]string
	D       int
	d       float64
	w       [][]float64
	ws      []float64
	vertex  []float64
	maxIter int
	minDiff float64
}

func textRank(docs [][]string) [][]string {
	tr := &textRankI{
		docs:    docs,
		D:       len(docs),
		d:       0.85,
		w:       [][]float64{},
		ws:      []float64{},
		vertex:  []float64{},
		maxIter: 200,
		minDiff: 0.001,
	}
	tr.solve()
	return tr.top()
}

func (tr *textRankI) sumWeight(weights []float64) float64 {
	ws := 0.0
	for _, w := range weights {
		ws += w
	}
	return ws
}

func (tr *textRankI) solve() {
	bm25 := sim.BM25(tr.docs)
	for i, doc := range tr.docs {
		weights := bm25.SimAll(doc)
		tr.w = append(tr.w, weights)
		tr.ws = append(tr.ws, tr.sumWeight(weights)-weights[i])
		tr.vertex = append(tr.vertex, 1.0)
	}
	for index := 0; index < tr.maxIter; index++ {
		m := []float64{}
		maxDiff := 0.0
		for i := 0; i < tr.D; i++ {
			m = append(m, 1.0-tr.d)
			for j := 0; j < tr.D; j++ {
				if i == j || tr.ws[j] == 0.0 {
					continue
				}
				m[len(m)-1] += tr.d * tr.w[j][i] / tr.ws[j] * tr.vertex[j]
			}
			if math.Abs(m[len(m)-1])-tr.vertex[i] > maxDiff {
				maxDiff = math.Abs(m[len(m)-1] - tr.vertex[i])
			}
		}
		tr.vertex = m
		if maxDiff <= tr.minDiff {
			break
		}
	}
}

func (tr *textRankI) top() [][]string {
	vl := vList{}
	for i, v := range tr.vertex {
		vs := &vertexStruct{
			mapkey: tr.docs[i],
			value:  v,
		}
		vl = append(vl, vs)
	}
	sort.Sort(vl)
	ret := [][]string{}
	for _, v := range vl {
		ret = append(ret, v.mapkey)
	}
	return ret
}
