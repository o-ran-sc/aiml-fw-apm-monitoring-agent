package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

var moduleName string
var loggers [3]*log.Logger
var logFlag int

// init initializes package global value.
func init() {
	moduleName = "mon-agent"
	logFlag = log.Ldate | log.Ltime
	loggers[0] = log.New(os.Stdout, fmt.Sprintf("[INFO][%s]", moduleName), logFlag)
	loggers[1] = log.New(os.Stdout, fmt.Sprintf("[DEBUG][%s]", moduleName), logFlag)
	loggers[2] = log.New(os.Stdout, fmt.Sprintf("[ERROR][%s]", moduleName), logFlag)
}

const (
	INFO = iota
	DEBUG
	ERROR
)

// Logging prints log stream on standard output with file name and function name, line.
func Logging(level int, msgs ...string) {
	pc, file, line, _ := runtime.Caller(1)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	funcName := parts[pl-1]

	packageName := ""
	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	loggers[level].Println(packageName, fileName, funcName, ":", strconv.Itoa(line), msgs)
}
