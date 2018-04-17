package distinguish

import (
	"testing"
	"os"
	"fmt"
	"io/ioutil"
)

func TestBatchProcess(t *testing.T) {

	tests := []struct {
		category         int
		zipFilename      string
		mappingsFilename string
	}{
		{0, "./testdata/test.zip", "./testdata/mappings.txt"},
	}

	for _, test := range tests {

		f, err := os.Open(test.zipFilename)
		if err != nil {
			t.Error("测试文件打开失败")
		}

		fi, err := f.Stat()
		if err != nil {
			f.Close()
			t.Error("获取测试文件 FileInfo 失败")
		}

		mappings, err := BatchProcess(test.category, f, fi.Size(), BatchDistinguish)
		if err != nil {
			f.Close()
			t.Error(err)
		}
		b, err := ioutil.ReadAll(mappings)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(string(b))
	}

}
