package etcd

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"

)
var cli *clientv3.Client
//初始化一个可以从etcd里面读取数据的客户端
func Init(addr string,duration time.Duration) (err error)  {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: duration,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	return
}
