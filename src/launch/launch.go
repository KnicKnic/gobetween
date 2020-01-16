package launch

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/yyyar/gobetween/api"
	"github.com/yyyar/gobetween/config"
	"github.com/yyyar/gobetween/info"
	"github.com/yyyar/gobetween/logging"
	"github.com/yyyar/gobetween/manager"
	"github.com/yyyar/gobetween/metrics"
)

/**
 * version,revision,branch should be set while build using ldflags (see Makefile)
 */
var (
	version  string
	revision string
	branch   string
)

func init() {

	// Set GOMAXPROCS if not set
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	// Init random seed
	rand.Seed(time.Now().UnixNano())

	// Save info
	info.Version = version
	info.Revision = revision
	info.Branch = branch
	info.StartTime = time.Now()
}

// Launch starts go between
func Launch(cfg config.Config) {

	info.Configuration = struct {
		Kind string `json:"kind"`
		Path string `json:"path"`
	}{"file", "in memory"}

	logging.Configure(cfg.Logging.Output, cfg.Logging.Level)

	// Start API
	go api.Start(cfg.Api)

	/* setup metrics */
	go metrics.Start(cfg.Metrics)

	// Start manager
	go manager.Initialize(cfg)
}
