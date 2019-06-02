package server

import "os"

// CreateLogger : Creates the logfile, the defered close and returns the file
func CreateLogger() *os.File {
	// logging access logs
	accessLog, err := os.Create("logs/access.log")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := accessLog.Close(); err != nil {
			panic(err)
		}
	}()

	return accessLog
}