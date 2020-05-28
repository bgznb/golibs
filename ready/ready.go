package ready

import (
	"os"
	"os/signal"

	"github.com/bgznb/golibs/conf"
	"github.com/bgznb/golibs/log"
	"github.com/bgznb/golibs/ready/module"
)

// Run ...
func Run(mods ...module.Module) {
	// logger
	if conf.LogLevel != "" {
		logger, err := log.New(conf.LogLevel, conf.LogPath, conf.LogFlag)
		if err != nil {
			panic(err)
		}
		log.Export(logger)
		defer logger.Close()
	}

	log.Release("reday %v starting up", version)

	// module
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}
	module.Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Release("ready closing down (signal: %v)", sig)
	module.Destroy()
}
