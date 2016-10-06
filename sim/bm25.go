package sim

import "math"

// BM25I the bm25 struct
type BM25I struct {
	Docs  [][]string
	D     float64 //the lenth of docs
	Avgdl float64
	F     []map[string]float64
	DF    map[string]float64
	IDF   map[string]float64
	K1    float64 // should be 1.5
	B     float64 // should be 0.75
}

// BM25  will be BM25I
func BM25(docs [][]string) *BM25I {

	bm := &BM25I{
		Docs: docs,
		D:    float64(len(docs)),
		K1:   1.5,
		B:    0.75,
		F:    []map[string]float64{},
		DF:   map[string]float64{},
		IDF:  map[string]float64{},
	}

	bm.Avgdl = bm.avgdl()

	for _, doc := range docs {

		bm.initDF(doc)

	}

	for k, v := range bm.DF {
		bm.IDF[k] = math.Log(bm.D-v+0.5) - math.Log(v+0.5)
	}

	return bm
}

func (bm25 *BM25I) initDF(doc []string) {

	tmp := map[string]float64{}
	for _, w := range doc {
		if _, ok := tmp[w]; !ok {
			tmp[w] = 0.0
		}
		tmp[w]++
	}
	bm25.F = append(bm25.F, tmp)

	for k := range tmp {
		if _, ok := bm25.DF[k]; !ok {
			bm25.DF[k] = 0.0
		}
		bm25.DF[k]++
	}

}

func (bm25 *BM25I) avgdl() float64 {
	sum := 0.0
	for _, doc := range bm25.Docs {
		sum += float64(len(doc))
	}
	return sum / bm25.D
}

func (bm25 *BM25I) sim(doc []string, index int) float64 {
	score := 0.0
	for _, w := range doc {
		if _, ok := bm25.F[index][w]; !ok {
			continue
		}
		d := float64(len(bm25.Docs[index]))
		score += (bm25.IDF[w] * bm25.F[index][w] * (bm25.K1 + 1.0) /
			(bm25.F[index][w] + bm25.K1*(1.0-bm25.B+bm25.B*d/bm25.Avgdl)))
	}
	return score
}

// SimAll return all the sim
func (bm25 *BM25I) SimAll(doc []string) []float64 {
	score := []float64{}
	for i := 0; i < int(bm25.D); i++ {
		score = append(score, bm25.sim(doc, i))
	}
	return score
}
