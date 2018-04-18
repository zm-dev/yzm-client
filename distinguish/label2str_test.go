package distinguish

import (
	"testing"
)

func TestGetLabel2StrFunc(t *testing.T) {
	_, err := GetLabel2StrFunc(100)
	if err == nil {
		t.Errorf("_, err := GetLabel2StrFunc(%d), err = nil, excepd error", 100)
	}

	tests := []struct {
		category    int
		label       Label
		mappingLine string
	}{
		{0, Label{"0000", "10+20*3"}, "0000,10+20*3=70\n"},
		{0, Label{"0001", "1-20*3"}, "0001,1-20*3=-59\n"},
		{0, Label{"0002", "10*2*3"}, "0002,10*2*3=60\n"},
		{0, Label{"0002", "17-"}, "0002,17-\n"},
		{0, Label{"0002", "17-**"}, "0002,17-**\n"},
		{0, Label{"0002", "-11*7--99"}, "0002,-11*7--99\n"},

		{1, Label{"0000", "ez6vD"}, "0000,EZ6VD\n"},
		{1, Label{"0001", "VE3G8"}, "0001,VE3G8\n"},
		{1, Label{"0002", "12VE8"}, "0002,12VE8\n"},

		{2, Label{"0000", "VE3G"}, "0000,VE3G\n"},
		{2, Label{"0001", "ez6v"}, "0001,EZ6V\n"},
		{2, Label{"0002", "12VE"}, "0002,12VE\n"},

		{3, Label{"0000", "1000"}, "0000,0\n"},
		{3, Label{"0001", "0100"}, "0001,1\n"},
		{3, Label{"0002", "0010"}, "0002,2\n"},
		{3, Label{"0003", "0001"}, "0003,3\n"},
	}

	for _, test := range tests {
		label2StrFunc, err := GetLabel2StrFunc(test.category)
		if err != nil {
			t.Error(err)
		}
		mappingLine := label2StrFunc(test.label)
		if test.mappingLine != mappingLine {
			t.Errorf("label2StrFunc(%v) = %s, exceptd %s\n", test.label, mappingLine, test.mappingLine)
		}
	}
}
