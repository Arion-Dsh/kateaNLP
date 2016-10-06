package summary

import (
	"math"
	"sort"
)

type keywordTRI struct {
	docs    []string
	words   map[string][]string
	vertex  map[string]float64
	d       float64
	maxIter int
	minDiff float64
}

func keywordTR(docs []string) []string {
	k := keywordTRI{
		docs:    docs,
		d:       0.85,
		maxIter: 200,
		minDiff: 0.01,
		words:   map[string][]string{},
		vertex:  map[string]float64{},
	}
	k.solve()
	return k.top()
}

func (k *keywordTRI) solve() {
	wordTmp := []string{}
	for _, word := range k.docs {
		if _, in := k.words[word]; !in {
			k.vertex[word] = 1.0
			k.words[word] = []string{}
		}
		wordTmp = append(wordTmp, word)
		if len(wordTmp) > 5 {
			wordTmp = wordTmp[1:len(wordTmp)]
		}
		for _, w1 := range wordTmp {
			for _, w2 := range wordTmp {
				if w1 == w2 {
					continue
				}
				k.words[w1] = append(k.words[w1], w2)
				k.words[w2] = append(k.words[w2], w1)
			}
		}
	}
	for i := 0; i < k.maxIter; i++ {
		m := map[string]float64{}
		maxDiff := 0.0
		tmp := map[string]float64{}
		for w, v := range k.vertex {
			if len(k.words[w]) > 0 {
				tmp[w] = v
			}
		}
		for w := range tmp {
			for _, j := range k.words[w] {
				if w == j {
					continue
				}
				if _, in := m[j]; !in {
					m[j] = 1.0 - k.d
				}
				m[j] += k.d / float64(len(k.words[w])) * k.vertex[w]
			}
		}
		for w := range k.vertex {
			if _, in := m[w]; in {
				if math.Abs(m[w]-k.vertex[w]) > maxDiff {
					maxDiff = math.Abs(m[w] - k.vertex[w])
				}
			}
		}
		k.vertex = m
		if maxDiff <= k.minDiff {
			break
		}
	}
}

func (k *keywordTRI) top() []string {
	vl := vList{}
	for w, v := range k.vertex {
		vs := &vertexStruct{
			key:   w,
			value: v,
		}
		vl = append(vl, vs)
	}
	sort.Sort(vl)
	ret := []string{}
	for _, w := range vl {
		ret = append(ret, w.key)
	}
	return ret
}

type vertexStruct struct {
	key    string
	value  float64
	mapkey []string
}

type vList []*vertexStruct

func (v vList) Len() int           { return len(v) }
func (v vList) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v vList) Less(i, j int) bool { return v[i].value > v[j].value }
