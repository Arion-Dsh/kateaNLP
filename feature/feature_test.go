package feature

import "testing"

func TestFeature(t *testing.T) {
	f := F{}
	fd3 := [][2]string{
		[2]string{"d", "s"},
		[2]string{"e", "m"},
		[2]string{"f", "e"},
	}
	f.Add(fd3, 1)

	f.Total = 3
	f.GetSamples()

}
