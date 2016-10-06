package seg

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"regexp"
	"strings"

	"github.com/Arion-Dsh/kateaNLP/feature"
	"github.com/Arion-Dsh/kateaNLP/utils"
)

type model struct {
	L1, L2, L3   float64
	Uni, Bi, Tri feature.F
	Status       []string
}

func (m *model) save(file string) {
	data, _ := json.Marshal(m)
	ioutil.WriteFile(file, data, 0777)
}

func (m *model) load(file string) {
	var err error
	var data []byte
	data, err = ioutil.ReadFile(file)
	if err != nil {
		panic("model load file err")
	}
	err = json.Unmarshal(data, m)
	if err != nil {
		panic("unmarshal josn to model get error")
	}
}

func (m *model) addSatus(t string) {
	if len(m.Status) == 0 {
		m.Status = []string{strings.ToLower(t)}
	} else {
		var in bool
		for _, s := range m.Status {
			if s == strings.ToLower(t) {
				in = true
				break
			}
		}
		if !in {
			m.Status = append(m.Status, strings.ToLower(t))
		}
	}
}

func (m *model) div(v1, v2 float64) float64 {
	var f float64
	if v2 == 0 {
		return f
	}
	return v1 / v2
}

func (m *model) train(data [][][2]string) {
	for _, sentence := range data {
		if len(sentence) == 0 {
			continue
		}

		now := [][2]string{
			[2]string{"", "BOS"},
			[2]string{"", "BOS"},
		}
		m.Uni.Add(now[:1], 2)
		m.Bi.Add(now, 1)
		for _, d := range sentence {
			// add status
			go m.addSatus(d[1])
			now = append(now, d)
			m.Uni.Add(now[2:], 1)
			m.Bi.Add(now[1:], 1)
			m.Tri.Add(now, 1)
			now = now[1:len(now)]
		}
	}
	var tl1, tl2, tl3 float64

	for _, now := range m.Tri.GetSamples() {
		_, c31 := m.Tri.Get(now)
		_, c32 := m.Bi.Get(now[:2])
		c3 := m.div(c31-1, c32-1)

		_, c21 := m.Bi.Get(now[1:])
		_, c22 := m.Uni.Get(now[1:2])
		c2 := m.div(c21-1, c22-1)

		_, c11 := m.Uni.Get(now[2:])
		c12 := m.Uni.GetSum()
		c1 := m.div(c11-1, c12-1)

		if c3 >= c1 && c3 >= c2 {
			tl3 += c31
		} else if c2 >= c1 && c2 >= c3 {
			tl2 += c31
		} else if c1 >= c2 && c1 >= c3 {
			tl1 += c31
		}

	}
	m.L1 = m.div(tl1, tl1+tl2+tl3)
	m.L2 = m.div(tl2, tl1+tl2+tl3)
	m.L3 = m.div(tl3, tl1+tl2+tl3)
}

func (m *model) logProb(s1, s2, s3 [2]string) float64 {
	k2 := [][2]string{s2}
	k3 := [][2]string{s3}
	k12 := [][2]string{s1, s2}
	k23 := [][2]string{s2, s3}
	k123 := [][2]string{s1, s2, s3}

	uni := m.L1 * m.Uni.Freq(k3)
	_, bi23 := m.Bi.Get(k23)
	_, uni2 := m.Uni.Get(k2)
	bi := m.div(m.L2*bi23, uni2)
	_, tri123 := m.Tri.Get(k123)
	_, bi12 := m.Bi.Get(k12)
	tri := m.div(m.L3*tri123, bi12)
	if uni+bi+tri == 0.0 {
		var z float64
		return -1 / z
	}
	return math.Log(uni + bi + tri)
}

func (m *model) tag(text string) [][2]string {
	re := regexp.MustCompile(`[[:digit:]^(\d+\.\d+)^(\d{3},\d+)]+|[[:alpha:]]+|\S`)
	data := re.FindAllString(text, -1)

	keyInit := [][2]string{
		[2]string{"", "BOS"},
		[2]string{"", "BOS"},
	}
	type now struct {
		k [][2]string
		p float64
		s string
	}
	nows := []*now{
		&now{
			k: keyInit,
			p: 0.0,
			s: "",
		},
	}

	for _, w := range data {
		tmp := map[string]*now{}
		tmps := []*now{}
		notFound := true
		for _, s := range m.Status {
			key := [][2]string{
				[2]string{w, s},
			}
			if m.Uni.Freq(key) != 0.0 {
				notFound = false
				break
			}
		}
		if notFound {
			for _, s := range m.Status {
				for _, n := range nows {
					sk := [][2]string{
						n.k[1],
						[2]string{w, s},
					}
					tmp[utils.StrList(sk)] = &now{
						k: sk,
						p: n.p,
						s: n.s + s,
					}
				}
			}
			for _, t := range tmp {
				tmps = append(tmps, t)
			}
			nows = tmps
			continue
		}
		for _, s := range m.Status {
			for _, n := range nows {
				p := n.p + m.logProb(n.k[0], n.k[1], [2]string{w, s})
				sk := [][2]string{
					n.k[1],
					[2]string{w, s},
				}
				_, ok := tmp[utils.StrList(sk)]
				if (!ok) || p > tmp[utils.StrList(sk)].p {
					tmp[utils.StrList(sk)] = &now{
						k: sk,
						p: p,
						s: n.s + s,
					}
				}
			}
		}
		for _, t := range tmp {
			tmps = append(tmps, t)
		}
		nows = tmps
	}

	var p float64
	var f *now
	for i, n := range nows {
		if i == 0 {
			p = -1 / p
		}
		if p < n.p {
			p = n.p
			f = n
		}
	}
	reData := make([][2]string, len(data))
	for i, tag := range f.s {
		reData[i] = [2]string{data[i], string(tag)}
	}
	return reData
}
