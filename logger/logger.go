package logger

import (
	"fmt"
	"github.com/gookit/color"
	"sync/atomic"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/9/9
   File: logger.go
*/

type RAltLogger struct {
	successChan chan string
	infoChan    chan string
	warningChan chan string
	errorChan   chan string
	printlnChan chan string
}

var (
	initState   = &atomic.Bool{}
	successChan chan string
	infoChan    chan string
	warningChan chan string
	errorChan   chan string
	printlnChan chan string
)

func Logger() *RAltLogger {
	if !initState.Load() {
		successChan = make(chan string)
		infoChan = make(chan string)
		warningChan = make(chan string)
		errorChan = make(chan string)
		printlnChan = make(chan string)
		initState.Store(true)
		go func() {
			for {
				select {
				case msg := <-successChan:
					currentTime := time.Now().Format("15:04:05")
					color.Print(fmt.Sprintf("[%v]  ", currentTime))
					color.Rgb(40, 225, 119, false).Println(msg)
				case msg := <-infoChan:
					currentTime := time.Now().Format("15:04:05")
					color.Print(fmt.Sprintf("[%v]  ", currentTime))
					color.Rgb(49, 122, 221, false).Println(msg)
				case msg := <-warningChan:
					currentTime := time.Now().Format("15:04:05")
					color.Print(fmt.Sprintf("[%v]  ", currentTime))
					color.Rgb(255, 153, 0, false).Println(msg)
				case msg := <-errorChan:
					currentTime := time.Now().Format("15:04:05")
					color.Print(fmt.Sprintf("[%v]  ", currentTime))
					color.Rgb(227, 80, 13, false).Println(msg)
				case msg := <-printlnChan:
					currentTime := time.Now().Format("15:04:05")
					color.Println(fmt.Sprintf("[%v]  %v", currentTime, msg))
				}
			}
		}()
		return &RAltLogger{successChan, infoChan, warningChan, errorChan, printlnChan}
	}
	return &RAltLogger{successChan, infoChan, warningChan, errorChan, printlnChan}
}

func (logger *RAltLogger) LogSuccess(msg string) { logger.successChan <- msg }
func (logger *RAltLogger) LogInfo(msg string)    { logger.infoChan <- msg }
func (logger *RAltLogger) LogWarn(msg string)    { logger.warningChan <- msg }
func (logger *RAltLogger) LogError(msg string)   { logger.errorChan <- msg }
func (logger *RAltLogger) LogPrintln(msg string) { logger.printlnChan <- msg }
func (logger *RAltLogger) LogSuccessf(msg string, args ...any) {
	logger.successChan <- fmt.Sprintf(msg, args...)
}
func (logger *RAltLogger) LogInfof(msg string, args ...any) {
	logger.infoChan <- fmt.Sprintf(msg, args...)
}
func (logger *RAltLogger) LogWarnf(msg string, args ...any) {
	logger.warningChan <- fmt.Sprintf(msg, args...)
}
func (logger *RAltLogger) LogErrorf(msg string, args ...any) {
	logger.errorChan <- fmt.Sprintf(msg, args...)
}
func (logger *RAltLogger) LogPrintlnf(msg string, args ...any) {
	logger.printlnChan <- fmt.Sprintf(msg, args...)
}
