package log

import (
	"errors"
	"log"
	"runtime"
	"time"
)

func ElapsedTime(tag string, msg string) func() {
	if msg != "" {
		log.Printf("[%s] %s", tag, msg)
	}

	start := time.Now()
	return func() { log.Printf("[%s] Elipsed Time: %s", tag, time.Since(start)) }
}

func MyError(msg string) error {
	if msg != "" {
		log.Printf("[%s]", msg)
	}
	return errors.New("[Error]:" + time.Now().String() + ": " + msg)
}

// func MyError(msg string) {
// 	if len(msg) == 0 {
// 		panic(myError(msg))
// 	}
// }

func MyLog(msg string) {
	if msg != "" {
		log.Printf("[%s]", msg)
	}
}

func Trace() (string, int, string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok == false {
		return "?", 0, "?"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return file, line, "?"
	}

	return file, line, fn.Name()
}

func TraceFn() string {
	pc, _, _, ok := runtime.Caller(1)
	if ok == false {
		return "?"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "?"
	}

	return fn.Name()
}
