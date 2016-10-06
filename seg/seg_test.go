package seg

import "testing"

func TestSeg(t *testing.T) {
	seg := &Seg{}
	// seg.Train("msr.txt")
	// seg.Save()
	seg.Load()
	// seg.Train("text.txt")
	seg.Cut("我们都是好孩子啊。测试一下吧？")
}
