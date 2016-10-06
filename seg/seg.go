package seg

import (
	"io/ioutil"
	"strings"

	"github.com/Arion-Dsh/kateaNLP/utils"
)

// Seg ...
type Seg struct {
	Words []string
	m     model
}

// Train ...
func (s *Seg) Train(file string) {

	f, err := ioutil.ReadFile(file)
	if err != nil {
		panic("error, please check file")
	}
	var data [][][2]string

	for _, line := range strings.Split(string(f), "\n") {
		lists := strings.Split(strings.TrimSpace(string(line)), " ")
		var wt [][2]string
		for _, word := range lists {
			w := strings.Split(strings.TrimSpace(word), "/")
			if len(w) > 1 {
				var wa [2]string
				copy(wa[:], w)
				wt = append(wt, wa)
			}
		}
		data = append(data, wt)
	}
	s.m.train(data)
}

// Save save the train data
func (s *Seg) Save() {
	model := utils.GetCurrentFilePath("train.model")
	s.m.save(model)
}

// Load ..
func (s *Seg) Load() {
	model := utils.GetCurrentFilePath("train.model")
	s.m.load(model)
}

// Cut ...
func (s *Seg) Cut(text string) []string {
	cutList := s.m.tag(text)

	s.Words = []string{}

	tmp := ""

	for _, wt := range cutList {
		if wt[1] == "s" {
			if tmp != "" {
				s.Words = append(s.Words, tmp)
				tmp = ""
			}
			s.Words = append(s.Words, wt[0])
			continue
		}
		if wt[1] == "b" {
			if tmp != "" {
				s.Words = append(s.Words, tmp)
				tmp = ""
			}
			tmp += wt[0]
			continue
		}

		if wt[1] == "m" {
			tmp += wt[0]
			continue
		}
		if wt[1] == "e" {
			tmp += wt[0]
			s.Words = append(s.Words, tmp)
			tmp = ""
		}

	}
	if tmp != "" {
		s.Words = append(s.Words, tmp)
	}
	return s.Words
}
