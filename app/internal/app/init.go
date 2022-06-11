package app

import (
	"flag"
	"log"
	"os"
	"runtime"

	"go.uber.org/zap"
	"ozon-url-shortener/app/internal/config"
	logs "ozon-url-shortener/app/pkg/logger"
)

type flags struct {
	config string
	cache  bool
	debug  bool
	https  bool
}

func Init() App {

	var (
		logPath    = flag.String("log", "log.txt", "log file path (default 'log.txt')")
		configPath = flag.String("config", "config.toml", "config path (default 'config.toml')")
		cache      = flag.Bool("cache", false, "use built-it storage")
		debug      = flag.Bool("debug", false, "enable debug mode")
		https      = flag.Bool("https", false, "run server in https mode")
	)

	flag.Parse()

	// // app
	app := App{
		flags: flags{
			cache: *cache,
			debug: *debug,
			https: *https,
		},
	}

	// // logger
	logger, closer, err := logs.New(*logPath)
	if err != nil {
		log.Fatalf("init logger: %v", err)
	}
	app.logger = logger
	app.closers.logger = closer

	// // config
	cfg, err := config.New(*configPath)
	if err != nil {
		app.logger.Error("read config",
			zap.String("path", app.flags.config),
			zap.Error(err),
		)

		// close logger
		err = closer()
		if err != nil {
			log.Fatalf("close logger: %v", err)
		}

		os.Exit(1)
	}
	app.cfg = cfg

	// //

	runtime.GC() // force garbage collector to clear unused flag pointers

	return app
}
