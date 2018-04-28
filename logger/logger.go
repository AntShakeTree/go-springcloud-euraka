package logger

import (
	"github.com/robfig/cron"
	"io"
	"log"
	"os"
	"time"

)

var (
	path    string
	lcorn   string
	lformat string
)


func CreateLogFile() {
	//
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", ":", err)
	}
	multi := io.MultiWriter(file, os.Stdout)
	initLogger(multi, multi, multi, multi)
	loggerCorn(lcorn, lformat)
}

func loggerCorn(corn string, fomat string) {
	c := cron.New()
	spec := corn
	c.AddFunc(spec, func() {
		str := time.Now().Format(fomat)
		create(str)
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Failed to open log file", ":", err)
		}
		multi := io.MultiWriter(file, os.Stdout)
		initLogger(multi, multi, multi, multi)
	})
	c.Start()
}

func create(file string) {
	_, err := os.Stat(file)
	if err != nil && os.IsNotExist(err) {
		os.Create(file)
	}
}
func initLogger(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}



var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)



