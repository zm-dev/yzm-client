package distinguish

import (
	"testing"
	"os"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		category int
		filename string
		yzmStr   string
	}{
		{0, "./testdata/9995.jpg", "11-13*15=-184"},
		{0, "./testdata/9996.jpg", "10-1-11=-2"},
		{0, "./testdata/9997.jpg", "13+9-12=10"},
		{0, "./testdata/9998.jpg", "12*8*14=1344"},
		{0, "./testdata/9999.jpg", "8+14*17=246"},
	}

	for _, test := range tests {
		f, err := os.Open(test.filename)
		if err != nil {
			t.Error(err)
		}

		yzmStr, err := Process(test.category, f)
		if err != nil {
			t.Error(err)
		}
		if yzmStr != test.yzmStr {
			t.Errorf("yzmStr, err=Process(%d, f), yzmStr=%s, excepd %s", test.category, yzmStr, test.yzmStr)
		}
		f.Close()
	}
}
