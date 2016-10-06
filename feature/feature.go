package feature

import (
	"github.com/Arion-Dsh/kateaNLP/utils"
)

// F ...
type F struct {
	D       map[string]float64
	Samples [][][2]string
	Total   float64
}

func (f *F) none() float64 {
	return 0.0
}

// Exists check key in Feature.d
func (f *F) Exists(key [][2]string) bool {
	if _, ok := f.D[utils.StrList(key)]; ok {

		return true
	}
	return false
}

// GetSum return Feature.total
func (f *F) GetSum() float64 {
	return f.Total
}

// Get return number of feature
func (f *F) Get(key [][2]string) (bool, float64) {
	if !f.Exists(key) {
		return false, f.none()
	}
	return true, f.D[utils.StrList(key)]
}

// Freq return freq
func (f *F) Freq(key [][2]string) float64 {
	_, numb := f.Get(key)
	return numb / f.Total
}

// GetSamples return all keys
func (f *F) GetSamples() [][][2]string {
	// [[word, tage], [word, tage]]
	return f.Samples
}

// Add add d
func (f *F) Add(key [][2]string, value float64) {

	if !f.Exists(key) {
		if f.D == nil {
			f.D = make(map[string]float64)
			f.Samples = make([][][2]string, len(f.D))
		}
		f.D[utils.StrList(key)] = 0.0
		f.Samples = append(f.Samples, key)
	}
	f.D[utils.StrList(key)] += value
	f.Total += float64(value)
}
