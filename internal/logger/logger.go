//package logger

//import "log"

//type Logger struct {
//namespace string
//}

//func New(namespace string) *Logger {
//return &Logger{namespace}
//}

//func (ll *Logger) Infof(message string, args ...interface{}) {
//log.Printf("[%s] %s %s", append([]interface{}{ll.namespace, message}, args...)...)
//}

//func (ll *Logger) Info(message string) {
//log.Printf("[%s] %s", ll.namespace, message)
//}

package logger

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	logrus "github.com/sirupsen/logrus"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

type Logger struct {
	logger    *logrus.Logger
	namespace string
}

func New(namespace string) *Logger {
	// define logrus intance
	log := logrus.New()

	// configure log
	mw := io.MultiWriter(os.Stdout)

	logrus.SetOutput(mw)
	// log.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.JSONFormatter{})

	return &Logger{
		logger:    log,
		namespace: namespace,
	}
}

func (l *Logger) WithData(message string, data map[string]interface{}) {
	fields := logrus.Fields{}

	for key, value := range data {
		k := toSnakeCase(key)
		fields[k] = value
	}

	l.logger.WithFields(fields).Info(message)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(fmt.Sprintf("[%s] ", l.namespace)+format, args...)
}

func (l *Logger) Info(message string) {
	l.logger.Info(message)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l *Logger) WithHTTPResponse(res *http.Response) {
	method := res.Request.Method
	url := res.Request.URL.String()
	statusCode := res.StatusCode
	b := res.Body
	body, err := ioutil.ReadAll(b)

	if err != nil {
		l.logger.Errorf("logger: Logger.WithHTTPResponse ioutil.ReadAll error: %s", err.Error())
		return
	}
	defer b.Close()
	bodyString := string(body)
	res.Body = ioutil.NopCloser(bytes.NewReader(body))

	l.logger.WithFields(logrus.Fields{
		"method":      method,
		"url":         url,
		"status_code": statusCode,
		"body":        bodyString,
	}).Info("HTTP RESPONSE")
}
