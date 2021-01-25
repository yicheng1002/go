package taillog

import (
	"fmt"

	"github.com/hpcloud/tail"
)

var tailobj *tail.Tail
//Init 函数用于初始化
func Init(filename string) (err error){
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},//从文件的哪个地方读
		MustExist: false,
		Poll:      true,
	}
	tailobj, err = tail.TailFile(filename,config)
	if err != nil {
		fmt.Println("tail file failed,err:",err)
		return
	}
	return
}

func ReadChan() (chan *tail.Line){
	return tailobj.Lines
}
