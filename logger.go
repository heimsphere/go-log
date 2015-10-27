package log

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
	debug bool
}

var Log *Logger

func init() {
	Log = Stdout()
}

func (l *Logger) EnableDebug() {
	l.debug = true
	l.Println("Debug logging enabled.")
}

func (l *Logger) DisableDebug() {
	l.debug = false
	l.Println("Debugging logging disabled.")
}

// FIXME add DEBUG prefix
func (l *Logger) Debug(v ...interface{}) {
	if l.debug {
		l.Print(v...)
	}
}
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.debug {
		l.Printf(format, v...)
	}
}
func (l *Logger) Debugln(v ...interface{}) {
	if l.debug {
		l.Println(v...)
	}
}

// FIXME add ERROR prefix
func (l *Logger) Error(v ...interface{}) {
	l.Print(v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Printf("ERROR: "+format, v...)
}

// FIXME add ERROR prefix
func (l *Logger) Errorln(v ...interface{}) {
	l.Println(v...)
}

func Stdout() *Logger {
	return &Logger{Logger: log.New(os.Stdout, "", log.LstdFlags), debug: false}
}

func Stderr() *Logger {
	return &Logger{Logger: log.New(os.Stderr, "", log.LstdFlags), debug: false}
}

type null struct{}

func (w *null) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (l *Logger) Disable() {
	l.Logger = log.New(&null{}, "", log.LstdFlags)
	l.debug = false
}
