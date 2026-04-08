package main

import (
	"github.com/brunohelius/migrai-code/cmd"
	"github.com/brunohelius/migrai-code/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
