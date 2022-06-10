package app

import (
	"flag"
	"runtime"
)

type flags struct {
	logger string
	config string
	cache  bool
	debug  bool
	https  bool
}

func Init() App {

	var (
		logger = flag.String("log", "log.txt", "logfile file path (default 'log.txt')")
		config = flag.String("config", "config.toml", "config path (default 'config.toml')")
		cache  = flag.Bool("cache", false, "use built-it storage")
		debug  = flag.Bool("debug", false, "enable debug mode")
		https  = flag.Bool("https", false, "run server in https mode")
	)

	defer runtime.GC() // force garbage collector to clear unused flag pointers
	flag.Parse()

	return App{
		flags: flags{
			logger: *logger,
			config: *config,
			cache:  *cache,
			debug:  *debug,
			https:  *https,
		},
	}
}
