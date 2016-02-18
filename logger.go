package goldshire

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

var logger = log.NewEntry(log.Log)

func init() {
	log.SetHandler(text.New(os.Stderr))
}

func SetLogger(l *log.Entry) {
	if l != nil {
		logger = l
	}
}

func GetLogger() *log.Entry {
	return logger
}
