package utils

import (
	"time"
	"github.com/astaxie/beego/logs"
	"fmt"
)

func Log(message error)  {
	//日志处理
	filename := time.Now().Format("20060102")
	logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"storage/logs/%s.log"}`, filename))
	logs.Trace(message)//跟踪
	logs.Info(message)//信息
	logs.Warn(message)//警告
	logs.Debug(message)//调试
	logs.Critical(message)//严重
	logs.SetLogFuncCall(true) //输出文件名和行号
	logs.Async()              //异步输出

}