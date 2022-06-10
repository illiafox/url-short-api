package app

import (
	"flag"
	"runtime"
)

type flags struct {
	logger   string
	config   string
	inmemory bool
	debug    bool
	https    bool
}

func Init() App {

	var (
		logger   = flag.String("logfile", "log.txt", "logfile file path (default 'logfile.txt')")
		config   = flag.String("config", "config.toml", "config path (default 'config.toml')")
		inmemory = flag.Bool("in-memory", false, "use in-memory storage")
		debug    = flag.Bool("debug", false, "enable debug mode")
		https    = flag.Bool("https", false, "run server in https mode")
	)

	defer runtime.GC() // force garbage collector to clear unused pointers
	flag.Parse()

	return App{
		flags: flags{
			logger:   *logger,
			config:   *config,
			inmemory: *inmemory,
			debug:    *debug,
			https:    *https,
		},
	}
}
