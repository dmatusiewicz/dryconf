package flags

import (
	"flag"
	"sync"
)

var c = flag.String("c", "./dryconf.yaml", "Load configuration from file.")
var v = flag.Bool("v", false, "Print DryConf version.")
var d = flag.Bool("d", false, "Run DryConf in debug mode.")

func oncef() {
	flag.Parse()
}

func Parse() {
	var once sync.Once
	once.Do(oncef)
}

func Config() string {
	return *c
}

func Version() bool {
	return *v
}

func Debug() bool {
	return *d
}
