package log

import (
	"fmt"
	"os"
	"time"
)

//日志类型
const (

)

type Log struct {
	Conent	string
	Status  bool
	Type	int
	WaitLog chan []*Log
}

var LogManager Log = Log{
	Conent:  "",
	Status:  false,
	Type:    0,
	WaitLog: nil,
}


func (l *Log) WriteLog() {
	//当前时间
	nowTime  := time.Now().Format("2006-01-02 15:04:05")
	//截取年月日
	filename :=	"log/log/"+string([]byte(nowTime)[:10])
	if l.Status {
		filename += "access.log"
	} else {
		filename += "error.log"
	}
	file,err := FileIsExist(filename)
	if err == nil {
		defer file.Close()
		_,err := fmt.Fprintln(file, l.Conent)
		if err != nil{
			fmt.Println(err)
		}
	}
}

//判断文件是否存在
//不存在则创建
func FileIsExist(filename string) (file *os.File,err error) {
	file,err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	//文件不存在则创建
	if err != nil && os.IsNotExist(err) {
		file,err = os.Create(filename)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
	return  file,err
}