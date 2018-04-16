package distinguish_service

import (
	"google.golang.org/grpc"
	"fmt"
	pb "github.com/3tnet/yzm-client/pkg/distinguish_service/protos"
	"sync"
	"os"
)

var (
	distinguishClient    pb.DistinguishClient
	distinguishClientErr error
	once                 sync.Once
)

// https://github.com/kubernetes/kubernetes/blob/8fd414537b5143ab039cb910590237cabf4af783/vendor/github.com/google/cadvisor/container/rkt/client.go
func Client() (pb.DistinguishClient, error) {
	once.Do(func() {
		address := os.Getenv("distinguish_server_address")

		if address == "" {
			address = "127.0.0.1:8080"
		}

		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			distinguishClient = nil
			distinguishClientErr = fmt.Errorf("connect remote Distinguish service %q failed, error: %v", address, err)
			return
		}
		//defer conn.Close()
		distinguishClient = pb.NewDistinguishClient(conn)
	})
	return distinguishClient, distinguishClientErr
}
