package ChainOfResponsibility

import "fmt"

const (
	NONE               = iota
	INFO               = iota
	DEBUG              = iota
	WARNING            = iota
	ERROR              = iota
	FUNCTIONAL_MESSAGE = iota
	FUNCTIONAL_ERROR   = iota
	ALL                = iota
)

type Handler interface {
	handle(string)
}

type Logger struct {
	logLevels []int
	next      *Logger
	handler   Handler
}

func (lo *Logger) init(logLevels []int, handler Handler) {
	lo.logLevels = logLevels
	lo.next = nil
	lo.handler = handler
}

func (lo *Logger) setNext(nx *Logger) {
	lo.next = nx
}

func (lo *Logger) message(msg string, level int) {
	if lo.handler == nil {
		panic("console cannot be nil")
	}
	for _, v := range lo.logLevels {
		if v == ALL || v == level {
			lo.handler.handle(msg)
			return
		}
	}
	if lo.next != nil {
		lo.next.message(msg, level)
	}
}

type ConsoleHandler struct {
}

func (c ConsoleHandler) handle(msg string) {
	fmt.Println("handle ALL:", msg)
}

type WarningHandler struct {
}

func (f WarningHandler) handle(msg string) {
	fmt.Println("handle WARNING:", msg)
}

func Test() {
	consoleLogger := &Logger{}
	consoleHandler := ConsoleHandler{}
	consoleLogger.init([]int{WARNING, ALL, NONE}, consoleHandler)
	warningLogger := &Logger{}
	warningHandler := &WarningHandler{}
	warningLogger.init([]int{WARNING}, warningHandler)
	warningLogger.setNext(consoleLogger)
	warningLogger.message("this is warning", WARNING)
	warningLogger.message("not a warning", FUNCTIONAL_ERROR)
}

/*
refer: https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
在这里warningLogger是后续Logger的入口
*/
