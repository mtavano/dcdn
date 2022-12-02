package logger

import "log"

type Logger struct {
	namespace string
}

func New(namespace string) *Logger {
	return &Logger{namespace}
}

func (ll *Logger) Infof(message string, args ...interface{}) {
	log.Printf("[%s] %s", append([]interface{}{ll.namespace}, args...)...)
}

func (ll *Logger) Info(message string) {
	log.Printf("[%s] %s", ll.namespace, message)
}
