package config

import (
	"time"
)

type Appconf struct {
	Kafka `ini:"kafka"`
	Taillog `ini:"taillog"`
	Etcd `ini:"etcd"` //表示ini配置文件中与Etcd字段对应的是etcd
}
type Kafka struct {
	Address string `ini:"address"`  //代码中只认识结构体变量，不认识ini文件里的配置，所以为了找到ini配置文件中具体的值，
	Topic string	`ini:"topic"`				//就把结构体变量映射到ini文件中相应的变量，从而通过结构体变量获得ini文件中的值
}
type Etcd struct {
	Address string
	Timeout time.Duration
}
type Taillog struct {
	Filename string	`ini:"filename"`
}
