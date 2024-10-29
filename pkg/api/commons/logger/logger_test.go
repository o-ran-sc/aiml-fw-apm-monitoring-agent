package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

type testInfo struct {
	enum   int
	name   string
	prefix string
	suffix string
}

var oldStdout *os.File

func setUpLogging(enum int, level string) (*os.File, *os.File) {
	oldStdout = os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	loggers[enum] = log.New(os.Stdout, fmt.Sprintf("[%s][EM]", level), logFlag)
	return r, w
}

func getPrintString(r *os.File, w *os.File) string {
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out)
}

func TestLoggerINFOLevel_ExpectINFOPrefixAndSuffix(t *testing.T) {
	msg := "test"
	testValue := testInfo{
		enum:   INFO,
		name:   "INFO",
		prefix: "[INFO][EM]",
		suffix: "[" + msg + "]\n",
	}

	r, w := setUpLogging(testValue.enum, testValue.name)
	Logging(testValue.enum, msg)
	returnString := getPrintString(r, w)

	if !strings.HasPrefix(returnString, testValue.prefix) {
		t.Error()
	}

	if !strings.HasSuffix(returnString, testValue.suffix) {
		t.Error()
	}
}

func TestLoggerDEBUGLevel_ExpectDEBUGPrefixAndSuffix(t *testing.T) {
	msg := "test"
	testValue := testInfo{
		enum:   DEBUG,
		name:   "DEBUG",
		prefix: "[DEBUG][EM]",
		suffix: "[" + msg + "]\n",
	}

	r, w := setUpLogging(testValue.enum, testValue.name)
	Logging(testValue.enum, msg)
	returnString := getPrintString(r, w)

	if !strings.HasPrefix(returnString, testValue.prefix) {
		t.Error()
	}

	if !strings.HasSuffix(returnString, testValue.suffix) {
		t.Error()
	}
}

func TestLoggerERRORLevel_ExpectERRORPrefixAndSuffix(t *testing.T) {
	msg := "test"
	testValue := testInfo{
		enum:   ERROR,
		name:   "error",
		prefix: "[error][EM]",
		suffix: "[" + msg + "]\n",
	}

	r, w := setUpLogging(testValue.enum, testValue.name)
	Logging(testValue.enum, msg)
	returnString := getPrintString(r, w)

	if !strings.HasPrefix(returnString, testValue.prefix) {
		t.Error()
	}

	if !strings.HasSuffix(returnString, testValue.suffix) {
		t.Error()
	}
}
