package utils

import (
	"fmt"
)

const reset uint8 = 0
const red uint8 = 31
const green uint8 = 32
const yellow uint8 = 33
const blue uint8 = 34
const magenta uint8 = 35
const cyan uint8 = 36
const gray uint8 = 37
const white uint8 = 97

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(message string) {
	println(applyColor(white, message))
}

func (l *Logger) Error(format string, args ...interface{}) {
	if len(args) > 0 {
		message := fmt.Sprintf(format, args...)
		println(applyColor(red, message))
		return
	}
	println(applyColor(red, format))
}

func (l *Logger) Warn(message string) {
	println(applyColor(yellow, message))
}

func (l *Logger) Success(format string, args ...interface{}) {
	if len(args) > 0 {
		message := fmt.Sprintf(format, args...)
		println(applyColor(green, message))
		return
	}
	println(applyColor(green, format))
}

func applyColor(color uint8, message string) string {
	return fmt.Sprintf("\033[%dm %s \033[0m", color, message)
}
