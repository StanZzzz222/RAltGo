package scripts

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/core/retimer/internal/entities"
	"github.com/StanZzzz222/RAltGo/logger"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type OP string

const (
	TCP     = "[TCP]"
	UDP     = "[UDP]"
	WEBHOOK = "WEBHOOK"
)

var sendChan = make(chan struct{}, 3000)

func analysisExpr(expr string) (string, string) {
	idx := strings.Index(expr, ":")
	exprType := expr[:idx]
	exprValue := expr[idx+1:]
	return exprType, exprValue
}

func ExprOperation(timer *entities.Timer) {
	exprType, exprValue := analysisExpr(timer.Expr)
	switch exprType {
	case TCP:
		idx := strings.Index(exprValue, ":")
		host := exprValue[:idx]
		portStr := exprValue[idx+1:]
		port, err := strconv.Atoi(portStr)
		if err != nil {
			logger.Logger().LogErrorf(":: TCP Port is not a valid type, Timer: %v", timer.Key)
			return
		}
		tcpNotify(host, int64(port), timer)
		break
	case UDP:
		idx := strings.Index(exprValue, ":")
		host := exprValue[:idx]
		portStr := exprValue[idx+1:]
		port, err := strconv.Atoi(portStr)
		if err != nil {
			logger.Logger().LogErrorf(":: UDP Port is not a valid type, Timer: %v", timer.Key)
			return
		}
		udpNotify(host, int64(port), timer)
		break
	case WEBHOOK:
		webhookNotify(exprValue, timer)
		break
	default:
		logger.Logger().LogErrorf(":: Unknow type %v, Expr: %v", exprType, timer.Expr)
		break
	}
}

func udpNotify(host string, port int64, timer *entities.Timer) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		sendChan <- struct{}{}
		var socket *net.UDPConn
		var err error
		socket, err = net.DialUDP("udp4", nil, &net.UDPAddr{
			IP:   net.ParseIP(host),
			Port: int(port),
		})
		if err != nil {
			udpNotify(host, port, timer)
			logger.Logger().LogInfof(":: Retry UDP timer notify to %v | Key: %v", fmt.Sprintf("%v:%v", host, port), timer.Key)
			return
		}
		jsonBytes, err := json.Marshal(timer)
		if err != nil {
			udpNotify(host, port, timer)
			logger.Logger().LogInfof(":: Retry UDP timer notify to %v | Key: %v", fmt.Sprintf("%v:%v", host, port), timer.Key)
			return
		}
		_, err = socket.Write(jsonBytes)
		if err != nil {
			udpNotify(host, port, timer)
			logger.Logger().LogInfof(":: Retry UDP timer notify to %v | Key: %v", fmt.Sprintf("%v:%v", host, port), timer.Key)
			return
		}
		<-sendChan
		wg.Done()
	}()
	wg.Wait()
}

func tcpNotify(host string, port int64, timer *entities.Timer) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		sendChan <- struct{}{}
		var socket *net.TCPConn
		var err error
		socket, err = net.DialTCP("tcp", nil, &net.TCPAddr{
			IP:   net.ParseIP(host),
			Port: int(port),
		})
		if err != nil {
			tcpNotify(host, port, timer)
			logger.Logger().LogInfof(":: Retry TCP timer notify to %v | Key: %v", fmt.Sprintf("%v:%v", host, port), timer.Key)
			return
		}
		jsonBytes, err := json.Marshal(timer)
		if err != nil {
			tcpNotify(host, port, timer)
			logger.Logger().LogInfof(":: Retry TCP timer notify to %v | Key: %v", fmt.Sprintf("%v:%v", host, port), timer.Key)
			return
		}
		_, err = socket.Write(jsonBytes)
		if err != nil {
			tcpNotify(host, port, timer)
			logger.Logger().LogInfof(":: Retry TCP timer notify to %v | Key: %v", fmt.Sprintf("%v:%v", host, port), timer.Key)
			return
		}
		<-sendChan
		wg.Done()
	}()
	wg.Wait()
}

func webhookNotify(url string, timer *entities.Timer) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		sendChan <- struct{}{}
		jsonBytes, err := json.Marshal(timer)
		if err != nil {
			webhookNotify(url, timer)
			logger.Logger().LogInfof(":: Retry WEBHOOK timer notify to %v | Key: %v", url, timer.Key)
			return
		}
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: time.Second * 8,
		}
		_, err = client.Post(url, "application/json", strings.NewReader(string(jsonBytes)))
		if err != nil {
			webhookNotify(url, timer)
			logger.Logger().LogInfof(":: Retry WEBHOOK timer notify to %v | Key: %v", url, timer.Key)
			return
		}
		<-sendChan
		wg.Done()
	}()
	wg.Wait()
}
