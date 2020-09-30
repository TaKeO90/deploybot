package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TaKeO90/deploybot/msg"
)

const (
	logFile                 = "bot.log"
	serviceLogPrefix string = "Service: "
	msgLogPrefix     string = "Message: "
	errorLogPrefix   string = "Error: "
)

type MessageLogs struct {
	Datetime   string
	MsgText    string
	SenderName string
	SenderId   int
	writer     *os.File
}

type ServicesLogs struct {
	ServiceIssuer struct {
		name string
		id   int
	}
	ServiceName string
	writer      *os.File
}

type ErrorLogs struct {
	ErrorMsg string
	writer   *os.File
}

// OpenLogFile
func OpenLogFile() (*os.File, error) {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0600)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// NewMsgLog
func NewMsgLog(writer *os.File, userData msg.SenderData, textData msg.TextMsgData) MessageLogs {
	senderName := fmt.Sprintf("%s %s", userData.Firstname, userData.Lastname)
	msgLogs := &MessageLogs{
		time.Unix(int64(textData.Date), 0).String(), textData.Text,
		senderName, userData.Id, writer,
	}
	return *msgLogs
}

// MsgLog
func (m MessageLogs) MsgLog() {
	msgLogger := log.New(m.writer, msgLogPrefix, log.Ldate)
	msgLogger.Printf(" [Sender ID] %d ||  [Sender Name] %s || [Text message] %s || [Sending Date/time] %s ", m.SenderId,
		m.SenderName, m.MsgText, m.Datetime)
}

// NewErrLogs
func NewErrorLog(writer *os.File, errorMsg string) ErrorLogs {
	errLogs := &ErrorLogs{errorMsg, writer}
	return *errLogs
}

// ErrLog
func (e ErrorLogs) ErrLog() {
	errLogger := log.New(e.writer, errorLogPrefix, log.Ldate)
	errLogger.Printf(" error message: %s ", e.ErrorMsg)
}

//TODO: service logs need a little bit of refactoring.

// NewServiceLogs
func NewServiceLog(writer *os.File, serviceName string, clientData msg.SenderData) ServicesLogs {
	id := clientData.Id
	name := fmt.Sprintf("%s %s", clientData.Firstname, clientData.Lastname)
	srvIssuer := struct {
		name string
		id   int
	}{
		name,
		id,
	}
	serviceLogs := &ServicesLogs{srvIssuer, serviceName, writer}
	return *serviceLogs
}

// SrvLog
func (s ServicesLogs) SrvLog() {
	serviceLogger := log.New(s.writer, serviceLogPrefix, log.Ldate)
	serviceLogger.Printf(" [service issuer] %s || [issuer id] %d || [service name] %s", s.ServiceIssuer.name,
		s.ServiceIssuer.id, s.ServiceName)
}
