package summary

import (
	"github.com/Arion-Dsh/kateaNLP/seg"
	"github.com/Arion-Dsh/kateaNLP/utils"
)

//Summary return summary
func Summary(docs string, seg *seg.Seg) []string {
	tmp := map[string]string{}
	doc := [][]string{}
	sentences := utils.GetSentences(docs, false)
	for _, sent := range sentences {
		words := seg.Cut(sent)
		words = utils.FliterStopWord(words)
		tmp[utils.StrList(words)] = sent
		doc = append(doc, words)
	}
	textRank := textRank(doc)
	ret := []string{}
	for _, s := range textRank {
		ret = append(ret, tmp[utils.StrList(s)])
	}
	return ret
}

// KeyWords ...
func KeyWords(docs string, seg *seg.Seg) []string {
	words := seg.Cut(docs)
	words = utils.FliterStopWord(words)
	return keywordTR(words)
}
