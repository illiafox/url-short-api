package app

import (
	"fmt"
	"log"
	"os"

	logs "ozon-url-shortener/app/pkg/zap"
)

func (app *App) Logger() {
	file, err := os.OpenFile(app.flags.logger, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(fmt.Errorf("create/open logfile file (%s): %w", app.flags.logger, err))
	}

	info, err := file.Stat()
	if err != nil {
		log.Fatalln(fmt.Errorf("get file stats: %w", err))
	}

	if info.Size() > 0 {
		_, err = file.Write([]byte("\n\n"))
		if err != nil {
			log.Fatalln(fmt.Errorf("write to file: %w", err))
		}
	}

	logger := logs.NewLogger(os.Stdout, file)
	app.logger = logger
	// //
	app.closers.logfile = func() {
		if err := file.Close(); err != nil {
			log.Println("close log file: ", err)
		}
	}
	//
	app.closers.logger = func() {
		_ = logger.Sync()
	}

}
