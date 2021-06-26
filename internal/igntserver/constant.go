package igntserver

import (
	"time"

	"github.com/changyoungkwon/gxample/internal/config"
)

var (
	igserverURI      = config.Get().IgServer.URI // save endpoint
	timeoutGetSingle = 1 * time.Second           // second
	timeoutGetAll    = 10 * time.Second          // second
)
