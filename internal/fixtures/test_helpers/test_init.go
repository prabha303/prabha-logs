package test_helpers

import (
	"prabha-logs/log-wrapper/internal/config/globals"
	"sync"

	"github.com/FenixAra/go-util/log"
)

var once sync.Once

func TestInit() *log.Logger {
	l := globals.Logger
	return l
}
