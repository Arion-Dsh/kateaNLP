package crf

import "testing"

func TestCrf(t *testing.T) {
	crf := Crf{}
	crf.Corps = []sentence{
		sentence{
			words: []string{"<S>", "I", "am", "test", "<S>"},
			tags:  []string{"<S>", "b", "m", "e", "<S>"},
		},
		sentence{
			words: []string{"<S>", "you", "are", "test", "<S>"},
			tags:  []string{"<S>", "b", "m", "e", "<S>"},
		},
	}

}
