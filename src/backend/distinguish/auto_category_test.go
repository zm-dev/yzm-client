package distinguish

import (
	"testing"
	"os"
)

func TestAutoCategory(t *testing.T) {
	tests := []struct {
		filename       string
		expectCategory int
	}{
		{"./testdata/9999.jpg", 0},
		{"./testdata/9998.jpg", 0},
		{"./testdata/9998.jpg", 0},
		{"./testdata/auto_category_test/1.jpg", 0},
		{"./testdata/auto_category_test/2.jpg", 1},
		{"./testdata/auto_category_test/3.jpg", 2},
		{"./testdata/auto_category_test/4.jpg", 3},
	}

	for _, test := range tests {
		imageFile, err := os.Open(test.filename)
		if err != nil {
			t.Error(err)
		}
		defer imageFile.Close()
		category, err := AutoCategory(imageFile)

		if err != nil {
			t.Error(err)
		}

		if category != test.expectCategory {
			t.Errorf("AutoCategory()=%d, expected %d", category, test.expectCategory)
		}

	}
}

func TestZipAutoCategory(t *testing.T) {
	tests := []struct {
		filename       string
		expectCategory int
	}{
		{"./testdata/test.zip", 0},
	}

	for _, test := range tests {
		zipFile, err := os.Open(test.filename)
		if err != nil {
			t.Error(err)
		}
		defer zipFile.Close()

		fileInfo, err := zipFile.Stat()
		if err != nil {
			t.Error(err)
		}

		category, err := ZipAutoCategory(zipFile, fileInfo.Size())
		if err != nil {
			t.Error(err)
		}

		if category != test.expectCategory {
			t.Errorf("ZipAutoCategory()=%d, expected %d", category, test.expectCategory)
		}
	}
}
