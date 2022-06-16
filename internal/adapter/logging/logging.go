package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	LogsDirpath = "logs"
)

type LogDir struct {
	LogDirectory string
}

func New() *LogDir {
	err := os.Mkdir(LogsDirpath, 0666)
	if err != nil {
		return nil
	}
	return &LogDir{
		LogDirectory: LogsDirpath,
	}
}

func SetLogFile() *os.File {
	year, month, day := time.Now().Date()
	fileName := fmt.Sprintf("%v-%v-%v.log", day, month.String(), year)
	filePath, _ := os.OpenFile(LogsDirpath+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	return filePath
}

func (l *LogDir) Info() *log.Logger {
	getFilePath := SetLogFile()
	return log.New(getFilePath, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *LogDir) Warning() *log.Logger {
	getFilePath := SetLogFile()
	return log.New(getFilePath, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *LogDir) Error() *log.Logger {
	getFilePath := SetLogFile()
	return log.New(getFilePath, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *LogDir) Fatal() *log.Logger {
	getFilePath := SetLogFile()
	return log.New(getFilePath, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}
