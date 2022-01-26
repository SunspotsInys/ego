package logs

import "fmt"

func (l *logger) Debug(msg string) {
	l.write(l.debug, msg)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.write(l.debug, msg)
}

func (l *logger) Info(msg string) {
	l.write(l.info, msg)
}

func (l *logger) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.write(l.info, msg)
}

func (l *logger) Error(msg string) {
	l.write(l.err, msg)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.write(l.err, msg)
}
