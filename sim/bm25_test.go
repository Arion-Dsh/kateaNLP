package sim

import (
	"fmt"
	"testing"
)

func TestSim(t *testing.T) {
	text := [][]string{
		[]string{"你好", "中国"},
		[]string{"你好", "世界"},
		[]string{"你好", "人们"},
	}
	bm := BM25(text)
	test := []string{"中国"}
	r := bm.SimAll(test)
	rr := []float64{
		0.5108256237659907,
		0,
		0,
	}
	fmt.Printf("return: %g \n right: %g \n", r, rr)
	if r == nil {
		t.Error()
	}
}
