package distinguish

import (
	"github.com/3tnet/yzm-client/pkg/distinguish_service"
	"context"
	pb "github.com/3tnet/yzm-client/pkg/distinguish_service/protos"
	"io"
	"io/ioutil"
	"fmt"
)

func Process(category int, filename string, r io.Reader) (label Label, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	distinguishClient, err := distinguish_service.Client()
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		err = fmt.Errorf("验证码图片读取失败，图片名称：%s。error:%+v", filename, err)
		return

	}
	yzm, err := distinguishClient.Distinguish(ctx, &pb.Image{Category: pb.Image_Category(category), Data: b})
	return Label{filename, yzm.Yzm}, nil
}
