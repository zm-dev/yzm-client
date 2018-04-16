package distinguish_service

import (
	"google.golang.org/grpc"
	"context"
	"time"
	"fmt"
	pb "yzm-client/pkg/distinguish_service/protos"
	"sync"
)

const (
	timeout = 5 * time.Second
)

type gRPCService struct {
	distinguishClient pb.DistinguishClient
}

var (
	gRPCServiceInstance  *gRPCService
	distinguishClientErr error
	once                 sync.Once
)

// https://github.com/kubernetes/kubernetes/blob/8fd414537b5143ab039cb910590237cabf4af783/vendor/github.com/google/cadvisor/container/rkt/client.go

func GetGRPCService(address string) (*gRPCService, error) {
	once.Do(func() {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			gRPCServiceInstance = nil
			distinguishClientErr = fmt.Errorf("connect remote Distinguish service %q failed, error: %v", address, err)
			return
		}
		//defer conn.Close()

		gRPCServiceInstance = &gRPCService{distinguishClient: pb.NewDistinguishClient(conn)}

	})

	return gRPCServiceInstance, distinguishClientErr
}

func (g *gRPCService) Distinguish(category pb.Image_Category, imageData []byte) (yzm string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	request := &pb.Image{Category: category, Data: imageData}

	response, err := g.distinguishClient.Distinguish(ctx, request)
	if err != nil {
		return "", err
	}

	return response.Yzm, nil
}

func (g *gRPCService) DistinguishData1(imageData []byte) (yzm string, err error) {
	return g.Distinguish(pb.Image_DATA1, imageData)
}

func (g *gRPCService) DistinguishData2(imageData []byte) (yzm string, err error) {
	return g.Distinguish(pb.Image_DATA2, imageData)
}

func (g *gRPCService) DistinguishData3(imageData []byte) (yzm string, err error) {
	return g.Distinguish(pb.Image_DATA3, imageData)
}

func (g *gRPCService) DistinguishData4(imageData []byte) (yzm string, err error) {
	return g.Distinguish(pb.Image_DATA4, imageData)
}

func (g *gRPCService) DistinguishData5(imageData []byte) (yzm string, err error) {
	return g.Distinguish(pb.Image_DATA5, imageData)
}
