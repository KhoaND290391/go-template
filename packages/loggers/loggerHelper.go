package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Formatter = new(logrus.TextFormatter) //default
	log.Level = logrus.InfoLevel
	file, err := os.OpenFile(getLogFilenameByTime(false), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0775)
	if err == nil {
		log.SetOutput(file)
	} else {
		file, err = os.OpenFile(getLogFilenameByTime(false), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0775)
		if err == nil {
			log.SetOutput(file)
		}
	}
}

//WriteInfo write log to file
func WriteInfo(pack string, method string, title string, mess string) {
	file, err := os.OpenFile(getLogFilenameByTime(false), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0775)
	if err == nil {
		log.SetOutput(file)
	} else {
		file, err = os.OpenFile(getLogFilenameByTimeBackup(false), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0775)
		if err == nil {
			log.SetOutput(file)
		}
	}
	log.Level = logrus.InfoLevel
	log.WithFields(logrus.Fields{
		"A.package": pack,
		"B.method":  method,
		"C.message": mess,
	}).Info()
}

//WriteError write error to file
func WriteError(pack string, method string, title string, mess string) {
	file, err := os.OpenFile(getLogFilenameByTime(true), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0775)
	if err == nil {
		log.SetOutput(file)
	} else {
		file, err = os.OpenFile(getLogFilenameByTimeBackup(true), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0775)
		if err == nil {
			log.SetOutput(file)
		}
	}
	log.Level = logrus.ErrorLevel
	log.WithFields(logrus.Fields{
		"A.package": pack,
		"B.method":  method,
		"C.message": mess,
	}).Error()
}

/* Private region */

func getLogFilenameByTime(isErr bool) string {
	if isErr {
		return "/var/log/app/" + time.Now().Format("2006-01-02") + "-err.log"
	}
	return "/var/log/app/" + time.Now().Format("2006-01-02") + ".log"
}

func getLogFilenameByTimeBackup(isErr bool) string {
	if isErr {
		return time.Now().Format("2006-01-02") + "-err.log"
	}
	return time.Now().Format("2006-01-02") + ".log"
}
