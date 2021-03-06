package main

import (
	"github.com/mkideal/log"
)

func main() {
	log.Trace("%s cannot be printed", "TRACE")
	log.Debug("%s cannot be printed", "DEBUG")
	log.Info("%s printed with std log package", "INFO")

	log.SetLevel(log.LvTRACE)
	log.Trace("%s printed with std log package", "TRACE")
	log.Debug("%s printed with std log package", "DEBUG")
}
