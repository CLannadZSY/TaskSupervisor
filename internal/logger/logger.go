package logger

import "fmt"

// 日志库

type Level int8

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

func InitLogger() {
	fmt.Print("init logger")
}
