package kateaNLP

import (
	"fmt"
	"testing"
)

func TestKateaNLP(t *testing.T) {
	docs := `中国是改革开放的受益者。38年来，中国逐步成长为世界经济的重要“动力源”和“稳定锚”。（数据来源：政府工作报告、世界银行等
	这些标签是中国实力增强的证明：1978年，中国经济总量仅占世界经济的份额1.8%。如今，中国已发展成全球第二大经济体、第一大贸易国、第一大外汇储备国、220多种工业产品产量全球第一、120多个国家和地区最大贸易伙伴……
	中国对世界经济增长平均贡献率已经在30%左右，居全球第一。“经济上的成功使中国走到世界舞台中心。”法国前总理拉法兰见证了中国的巨变。美国前国务卿基辛格也慨叹，他第一次访华时，绝对不曾想到中国的实力地位能在国际体系中有如此巨大的跃升。
	与此同时，中国人民也尝到了甜头。1978年，中国人均国民总收入不足200美元。如今，该项收入已增至约7880美元。按世界银行的划分标准，中国已经由低收入国家跃升至中上等收入国家。期间，共计7亿多人口脱困，中国对全球减贫贡献率逾70%……`
	nlp := KateaNLP(docs)
	words := nlp.Cut()
	fmt.Printf("\nwords are : %s\n", words)
	fmt.Printf("\nkey words are: %s\n", nlp.KeyWords(5))
}
