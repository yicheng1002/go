package main

import (
	"fmt"
	"logagent/config"
	"logagent/etcd" //导入的包名是从$GOPATH/src/后开始计算的，在编译的时候，报此包不在goroot中，说明是把goroot当成gopath来用了
	"logagent/kafka"
	"logagent/taillog" //这个导入路径要相对于main.go文件的路径，
	"time"

	_ "github.com/hpcloud/tail" //记住这个地方的用法！！！！！
	"gopkg.in/ini.v1"
)

//var cfg config.Appconf
//cfg := new(config.Appconf)  不能这种写法，得是下面这种写法
var cfg = new(config.Appconf)

func run() {
	for {
		select {
		case line := <-taillog.ReadChan(): //读取日志,taillog.ReadChan()返回值是*tail.Line的通道，需要把通道里的值输送给line
			//发送到kafka
			kafka.SendToKafka(cfg.Kafka.Topic, line.Text)
			//		kafka.SendToKafka("web_log",line.Text)
			fmt.Println("line.Text:", line.Text)
			if line.Text == "exit" {
				return
			}
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	//cfg,err := ini.Load("./config/config.ini") //配置文件是一个整体，包含了key=kafka和key=taillog，
	//映射的时候也需要往整体上映射，所以声明一个包含kafka和taillog结构体的结构体，然后配置文件映射到这个结构体上

	err := ini.MapTo(cfg, "./config/config.ini") //把配置文件的值加载出来映射到cfg这个结构体指针变量上
	if err != nil {                              //cfg表示指向结构体的指针，是不是说明把配置文件的值赋值给这块地址中存的变量
		fmt.Println("load ini failed,err:", err)
	}
	//0.加载配置文件
	//cfg,err := ini.Load("./config/config.ini")
	//cfg.Section("kafka").Key("address")
	//cfg.Section("kafka").Key("topic")
	//cfg.Section("taillog").Key("filename")
	//1.初始化kafka连接
	err = kafka.Init([]string{cfg.Kafka.Address})
	if err != nil {
		fmt.Println("kafka init failed,err:", err)
	}
	fmt.Println("初始化kafka连接成功")

	//2.初始化etcd连接
	err = etcd.Init(cfg.Etcd.Address, time.Duration(cfg.Etcd.Timeout)*time.Second)
	if err != nil {
		fmt.Println("etcd init failed,err:", err)
	}
	fmt.Println("初始化etcd成功")

	//使用etcd代替下面用tailf模块从日志文件中读日志的过程
	//2.打开日志文件准备收集日志
	//err = taillog.Init(cfg.Taillog.Filename)
	//if err != nil {
	//	fmt.Printf("taillog field,err:%v\n",err)
	//}
	//fmt.Println("顺利打开日志文件")
	//run()
}
