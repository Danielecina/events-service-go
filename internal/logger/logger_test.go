package logger

import (
	"bytes"
	"log"
	"testing"

	snaps "github.com/gkampitakis/go-snaps/snaps"
)

func TestInfo(t *testing.T) {
	var buf bytes.Buffer
	AppLogger = log.New(&buf, "[APP] 2006/01/02 15:04:05 logger.go:12: ", 0)
	Info("test info %s", "message")
	output := buf.String()
	snaps.MatchSnapshot(t, output)
}

func TestError(t *testing.T) {
	var buf bytes.Buffer
	AppLogger = log.New(&buf, "[APP] 2006/01/02 15:04:05 logger.go:19: ", 0)
	Error("test error %d", 123)
	output := buf.String()
	snaps.MatchSnapshot(t, output)
}

func TestDebug(t *testing.T) {
	var buf bytes.Buffer
	AppLogger = log.New(&buf, "[APP] 2006/01/02 15:04:05 logger.go:26: ", 0)
	Debug("test debug %v", true)
	output := buf.String()
	snaps.MatchSnapshot(t, output)
}
