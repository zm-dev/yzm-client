package distinguish

import (
	"testing"
	"os"
	"bytes"
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

		mappingsBytes, err := ioutil.ReadAll(mappings)
		if err != nil {
			t.Error(err)
		}

		mf, err := os.Open(test.mappingsFilename)
		if err != nil {
			t.Error(err)
		}
		defer mf.Close()

		mfBytes, err := ioutil.ReadAll(mf)
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(bytes.TrimSuffix(mappingsBytes, []byte("\r\n")), bytes.TrimSuffix(mfBytes, []byte("\r\n"))) {
			t.Errorf("zipFilename: %s, 识别的内容和 mappingsFilename: %s, 不符", test.zipFilename, test.mappingsFilename)
		}
	}

}
