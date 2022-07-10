package internal

import (
	"sync"
)

var Cleanup = new(sync.WaitGroup)
