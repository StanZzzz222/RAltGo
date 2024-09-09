package logger

import (
	"fmt"
	"github.com/gookit/color"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/9/9
   File: logger.go
*/

var (
	successChan = make(chan string)
	infoChan    = make(chan string)
	warningChan = make(chan string)
	errorChan   = make(chan string)
	printlnChan = make(chan string)
)

func init() {
	go func() {
		for {
			select {
			case msg := <-successChan:
				currentTime := time.Now().Format("15:04:05")
				color.Rgb(40, 225, 119, false).Println(fmt.Sprintf("[%v]: %v", currentTime, msg))
			case msg := <-infoChan:
				currentTime := time.Now().Format("15:04:05")
				color.Rgb(49, 122, 221, false).Println(fmt.Sprintf("[%v]: %v", currentTime, msg))
			case msg := <-warningChan:
				currentTime := time.Now().Format("15:04:05")
				color.Rgb(255, 153, 0, false).Println(fmt.Sprintf("[%v]: %v", currentTime, msg))
			case msg := <-errorChan:
				currentTime := time.Now().Format("15:04:05")
				color.Rgb(227, 80, 13, false).Println(fmt.Sprintf("[%v]: %v", currentTime, msg))
			case msg := <-printlnChan:
				currentTime := time.Now().Format("15:04:05")
				color.Println(fmt.Sprintf("[%v]: %v", currentTime, msg))
			}
		}
	}()
}

func LogSuccess(msg string)               { successChan <- msg }
func LogInfo(msg string)                  { infoChan <- msg }
func LogWarn(msg string)                  { warningChan <- msg }
func LogError(msg string)                 { errorChan <- msg }
func LogPrintln(msg string)               { printlnChan <- msg }
func LogSuccessf(msg string, args ...any) { successChan <- fmt.Sprintf(msg, args...) }
func LogInfof(msg string, args ...any)    { infoChan <- fmt.Sprintf(msg, args...) }
func LogWarnf(msg string, args ...any)    { warningChan <- fmt.Sprintf(msg, args...) }
func LogErrorf(msg string, args ...any)   { errorChan <- fmt.Sprintf(msg, args...) }
func LogPrintlnf(msg string, args ...any) { printlnChan <- fmt.Sprintf(msg, args...) }
