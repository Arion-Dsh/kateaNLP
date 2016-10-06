package seg

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestModel(t *testing.T) {
	file, _ := ioutil.ReadFile("Text.txt")

	var data [][][2]string

	for _, line := range strings.Split(string(file), "\n") {
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
	// fmt.Print("\n")
	// m := &Model{}
	// m.Load("train")
	// m.Status = []string{"b", "m", "e", "s"}
	// m.Train(data)
	// m.Tag("这个真的很赞你说是不是啊我们一起吧在家里")
	// m.Save()
	// testF := "1.23"

}
