package distinguish

import (
	"io"
	"archive/zip"
	"github.com/3tnet/yzm-client/pkg/distinguish_service"
	pb "github.com/3tnet/yzm-client/pkg/distinguish_service/protos"
	"context"
	log "github.com/sirupsen/logrus"
	"time"
	"io/ioutil"
)

type batchDistinguishFunc func(ctx context.Context, imageChan <-chan Image, yzmChan chan<- Label) error

func BatchProcess(category int, zipFile io.ReaderAt, size int64, bdfunc batchDistinguishFunc) (yzmRecvChan <-chan Label, err error) {

	reader, err := zip.NewReader(zipFile, size)

	if err != nil {
		return nil, err
	}

	imageChan := make(chan Image)
	yzmChan := make(chan Label)

	go bdfunc(context.Background(), imageChan, yzmChan)
	go bdfunc(context.Background(), imageChan, yzmChan)

	go func() {
		for _, image := range reader.File {
			r, err := image.Open()
			if err != nil {
				log.Warnf("验证码图片打开失败，图片名称：%s。error:%+v", image.Name, err)
				continue
			}
			b, err := ioutil.ReadAll(r)
			if err != nil {
				log.Warnf("验证码图片读取失败，图片名称：%s。error:%+v", image.Name, err)
				continue
			}
			imageChan <- Image{Filename: image.Name, Category: category, Data: b}
		}
	}()

	return yzmChan, nil
}

type Image struct {
	Filename string
	Category int
	Data     []byte
}

type Label struct {
	ImageFilename string
	Yzm           string
}

const (
	timeout = 5 * time.Second
)

func BatchDistinguish(ctx context.Context, imageChan <-chan Image, yzmChan chan<- Label) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	distinguishClient, err := distinguish_service.Client()
	if err != nil {
		return err
	}

	ok := true
	var image Image

	for ok {
		image, ok = <-imageChan
		response, err := distinguishClient.Distinguish(ctx, &pb.Image{Category: pb.Image_Category(image.Category), Data: image.Data})
		if err != nil {
			log.Warnf("验证码识别出错：%+v", err)
		} else {
			yzmChan <- Label{ImageFilename: image.Filename, Yzm: response.Yzm}
		}
	}
	return nil
}
