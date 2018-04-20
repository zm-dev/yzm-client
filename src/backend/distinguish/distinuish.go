package distinguish

import (
	"github.com/zm-dev/yzm-client/src/backend/pkg/distinguish_service"
	"context"
	pb "github.com/zm-dev/yzm-client/src/backend/pkg/distinguish_service/protos"
	"io"
	"io/ioutil"
	"fmt"
)

func Process(category int, r io.Reader) (yzmStr string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	distinguishClient, err := distinguish_service.Client()
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		err = fmt.Errorf("验证码图片读取失败。error:%+v", err)
		return

	}
	yzm, err := distinguishClient.Distinguish(ctx, &pb.Image{Category: pb.Image_Category(category), Data: b})

	if err != nil {
		return "", err
	}

	label2StrFunc, err := GetLabel2StrFunc(category)

	if err != nil {
		return "", err
	}
	return label2StrFunc(Label{"", yzm.Yzm}), nil
}
