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
	"bytes"
	"sync"
	"path"
	"strings"
	"fmt"
)

const (
	imageSuffix  = ".jpg"
	ignorePrefix = "__MACOSX/"
)

type mappingsReader struct {
	buf           *bytes.Buffer
	lock          sync.Mutex
	LabelRecvChan <-chan Label
	Label2StrFunc Label2StrFunc
}

func (m *mappingsReader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	for m.buf.Len() < len(p) {
		label, ok := <-m.LabelRecvChan
		if !ok {
			break
		}
		m.buf.WriteString(m.Label2StrFunc(label))
	}
	n, err = m.buf.Read(p)
	m.buf.Reset()
	m.buf.Write(m.buf.Bytes())
	return

}

func newMappingsReader(labelRecvChan <-chan Label, label2StrFunc Label2StrFunc) *mappingsReader {
	return &mappingsReader{buf: &bytes.Buffer{}, LabelRecvChan: labelRecvChan, Label2StrFunc: label2StrFunc}
}

type batchDistinguishFunc func(ctx context.Context, imageChan <-chan Image, labelChan chan<- Label) error

func BatchProcess(category int, zipFile io.ReaderAt, size int64, bdfunc batchDistinguishFunc) (mappings io.Reader, err error) {

	reader, err := zip.NewReader(zipFile, size)

	if err != nil {
		return nil, err
	}

	label2StrFunc, err := GetLabel2StrFunc(category)

	if err != nil {
		return nil, err
	}

	imageChan := make(chan Image)
	labelChan := make(chan Label)
	ctx := context.Background()
	go bdfunc(ctx, imageChan, labelChan)
	// go bdfunc(ctx, imageChan, labelChan)

	go func() {
		for _, image := range reader.File {
			func(image *zip.File) {
				if path.Ext(image.Name) != imageSuffix || strings.HasPrefix(image.Name, ignorePrefix) {
					return
				}
				r, err := image.Open()
				defer r.Close()
				if err != nil {

					log.Warnf("验证码图片打开失败，图片名称：%s。error:%+v", image.Name, err)
					return
				}
				b, err := ioutil.ReadAll(r)
				if err != nil {
					log.Warnf("验证码图片读取失败，图片名称：%s。error:%+v", image.Name, err)
					return
				}

				select {
				case imageChan <- Image{Filename: image.Name, Category: category, Data: b}:
				case <-ctx.Done():
					break
				}
			}(image)
		}
		close(imageChan)
	}()

	return newMappingsReader(labelChan, label2StrFunc), nil
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

func BatchDistinguish(ctx context.Context, imageChan <-chan Image, labelChan chan<- Label) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	distinguishClient, err := distinguish_service.Client()
	if err != nil {
		return err
	}

	ok := true
	var image Image

	for ok {
		select {
		case image, ok = <-imageChan:
			if ok {
				fmt.Println(image.Filename)
				response, err := distinguishClient.Distinguish(ctx, &pb.Image{Category: pb.Image_Category(image.Category), Data: image.Data})
				if err != nil {
					//f, _ := os.Create(image.Filename)
					//f.Write(image.Data)
					//f.Close()
					log.Warnf("验证码识别出错：%+v", err)
				} else {
					labelChan <- Label{ImageFilename: image.Filename, Yzm: response.Yzm}
				}
			}
		case <-ctx.Done():
			break
		}
	}
	// todo 这里在多个协程里面关闭可能会有问题
	close(labelChan)
	return nil
}
