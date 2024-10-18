package scripts

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/core/retimer/timer"
	"github.com/StanZzzz222/RAltGo/logger"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type OP string

const (
	TCP     = "[TCP]"
	UDP     = "[UDP]"
	WEBHOOK = "[WEBHOOK]"
)

func ExecuteTimerExpression(timer *timer.ITimer) {
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
		go tcpNotify(host, int64(port), timer)
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
		go udpNotify(host, int64(port), timer)
		break
	case WEBHOOK:
		go webhookNotify(exprValue, timer)
		break
	default:
		logger.Logger().LogErrorf(":: Unknow type %v, Expr: %v", exprType, timer.Expr)
		break
	}
}

func analysisExpr(expr string) (string, string) {
	idx := strings.Index(expr, ":")
	exprType := expr[:idx]
	exprValue := expr[idx+1:]
	return exprType, exprValue
}

func udpNotify(host string, port int64, timer *timer.ITimer) {
	var socket *net.UDPConn
	var err error
	socket, err = net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.ParseIP(host),
		Port: int(port),
	})
	if err != nil {
		go udpNotify(host, port, timer)
		logger.Logger().LogInfof(":: Retry UDP timer notify to %v | Key: %v", fmt.Sprintf("%v:%v", host, port), timer.Key)
		return
	}
	jsonBytes, err := json.Marshal(timer)
	if err != nil {
		go udpNotify(host, port, timer)
		logger.Logger().LogInfof(":: Retry UDP timer notify to %v | Key: %v", fmt.Sprintf("%v:%v", host, port), timer.Key)
		return
	}
	_, err = socket.Write(jsonBytes)
	if err != nil {
		go udpNotify(host, port, timer)
		logger.Logger().LogInfof(":: Retry UDP timer notify to %v | Key: %v", fmt.Sprintf("%v:%v", host, port), timer.Key)
		return
	}
}

func tcpNotify(host string, port int64, timer *timer.ITimer) {
	var socket *net.TCPConn
	var err error
	socket, err = net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.ParseIP(host),
		Port: int(port),
	})
	if err != nil {
		go tcpNotify(host, port, timer)
		logger.Logger().LogInfof(":: Retry TCP timer notify to %v | Key: %v", fmt.Sprintf("%v:%v", host, port), timer.Key)
		return
	}
	jsonBytes, err := json.Marshal(timer)
	if err != nil {
		go tcpNotify(host, port, timer)
		logger.Logger().LogInfof(":: Retry TCP timer notify to %v | Key: %v", fmt.Sprintf("%v:%v", host, port), timer.Key)
		return
	}
	_, err = socket.Write(jsonBytes)
	if err != nil {
		go tcpNotify(host, port, timer)
		logger.Logger().LogInfof(":: Retry TCP timer notify to %v | Key: %v", fmt.Sprintf("%v:%v", host, port), timer.Key)
		return
	}
}

func webhookNotify(url string, timer *timer.ITimer) {
	jsonBytes, err := json.Marshal(timer)
	if err != nil {
		go webhookNotify(url, timer)
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
		go webhookNotify(url, timer)
		logger.Logger().LogInfof(":: Retry WEBHOOK timer notify to %v | Key: %v", url, timer.Key)
		return
	}
}
