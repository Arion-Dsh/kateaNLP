package crf

// Crf the crf struct
type Crf struct {
	corps []sentence
}

type sentence struct {
	words []string // add "<S>" to the sentence start and end.
	tags  []string // the "<S>" 's tag is "<S>"
}

// Corps format the data
func (crf *Crf) Corps(data string) {

}
