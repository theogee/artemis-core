package logger

import (
	"io"
	stdlog "log"
	"os"
)

var (
	Log *stdlog.Logger
	f *os.File
)

func Setup(path string) {
	var err error

	f, err = os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		stdlog.Fatalf("[logger.Setup] error opening log file. path: %v. err: %v", path, err)
	}

	multiWriter := io.MultiWriter(os.Stdout, f)

	Log = stdlog.New(multiWriter, "", stdlog.LstdFlags)

	Log.Println("[logger.Setup] log file opened successfully")
}

func Close() {
	if f == nil {
		stdlog.Fatalln("[logger.Close] error pointer is nil")
	}
	f.Close()
}