package kateaNLP

import (
	"github.com/Arion-Dsh/kateaNLP/seg"
	"github.com/Arion-Dsh/kateaNLP/sim"
	"github.com/Arion-Dsh/kateaNLP/summary"
)

//KateaNLP init the NLP struct
func KateaNLP(docs string) *NLP {
	nlp := &NLP{
		Docs: docs,
		Seg:  &seg.Seg{},
	}
	nlp.Seg.Load()
	return nlp
}

//NLP the NLP struct
type NLP struct {
	Docs string
	Seg  *seg.Seg
}

//Cut cut the docs
func (nlp *NLP) Cut() []string {
	return nlp.Seg.Cut(nlp.Docs)
}

// Summary get the summary. if the limit < 0 or limit > len(summary) will be return all the summary
func (nlp *NLP) Summary(limit int) []string {
	summary := summary.Summary(nlp.Docs, nlp.Seg)[:limit]
	if len(summary) > limit && limit >= 0 {
		summary = summary[:limit]
	}
	return summary
}

// KeyWords get the keyords of docs, if the limit <0 or limit > len(keyords) will be return all the
// Keywords.
func (nlp *NLP) KeyWords(limit int) []string {
	kws := summary.KeyWords(nlp.Docs, nlp.Seg)
	if len(kws) > limit && limit >= 0 {
		kws = kws[:limit]
	}
	return kws
}

// SimAll return all similarity between text and each original doc
func (nlp *NLP) SimAll(original []string, text string) []float64 {
	originalList := [][]string{}
	for _, o := range original {
		seg := nlp.Seg.Cut(o)
		originalList = append(originalList, seg)
	}
	sim := sim.BM25(originalList)
	return sim.SimAll(nlp.Seg.Cut(text))
}
