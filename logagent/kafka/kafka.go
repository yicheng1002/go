package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var client sarama.SyncProducer //声明一个全局的连接kafka的生产者client

//Init 函数初始化一个连向kafka的连接
func Init(addr []string) (err error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition
	config.Producer.Return.Successes = true                   //成功交付的消息将在success channel返回

	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("producer err:", err)
		return
	}
	fmt.Println("连接kafka成功")
	return
}

func SendToKafka(topic string,data string) {
	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	//发送到kafka
	pid,offset,err := client.SendMessage(msg) //因为SendMessage函数的参数是*ProducerMessage类型，所以上面msg要声明成这种类型
	if err != nil{
		fmt.Println("send msg failed,err:",err)
		return
	}
	fmt.Printf("pid:%v offset:%v msg:%v\n",pid,offset,msg)

}
