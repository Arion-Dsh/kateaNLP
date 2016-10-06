package summary

import (
	"fmt"
	"testing"
)

func TestTextRank(t *testing.T) {
	test := [][]string{
		[]string{"自然语言", "处理", "计算机", "科学", "领域", "人工智能", "领域", "一个", "重要", "方向"},
		[]string{"研究", "实现", "人", "计算机", "自然语言", "进行", "有效通信", "理论", "方法"},
		[]string{"自然语言", "处理", "一门", "融", "语言学", "计算机", "科学", "数学", "一体", "科学"},
		[]string{"因此", "领域", "研究", "涉及", "自然语言", "人们", "日常", "使用", "语言"},
		[]string{"自然语言", "处理", "研究", "自然语言"},
	}
	tr := textRank(test)
	fmt.Print(tr)
}
