package distinguish

import (
	"image/jpeg"
	"io"
	"errors"
	"archive/zip"
	"path"
	"strings"
)

var ErrCannotAutoCategory = errors.New("无法自动分类")

func NeedAutoCategory(category int) bool {
	return category == -1
}

func AutoCategory(imageFile io.Reader) (category int, err error) {
	config, err := jpeg.DecodeConfig(imageFile)

	// defer imageFile.Seek(io.SeekStart, io.SeekStart)

	if err != nil {
		return -1, err
	}
	if config.Height == 80 && config.Width == 350 {
		return 0, nil
	} else if config.Height == 60 && config.Width == 200 {
		return 1, nil
	} else if config.Height == 80 && config.Width == 200 {
		return 2, nil
	} else if config.Height == 45 && config.Width == 150 {
		return 3, nil
	}

	return -1, ErrCannotAutoCategory
}

func ZipAutoCategory(zipFile io.ReaderAt, size int64) (category int, err error) {
	reader, err := zip.NewReader(zipFile, size)
	for _, image := range reader.File {
		if path.Ext(image.Name) == imageSuffix && !strings.HasPrefix(image.Name, ignorePrefix) {

			imageFile, err := image.Open()
			if err != nil {
				return -1, err
			}
			defer imageFile.Close()
			return AutoCategory(imageFile)
		}
	}
	return -1, ErrCannotAutoCategory
}
