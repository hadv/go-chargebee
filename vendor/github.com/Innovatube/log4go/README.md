# Log for golang, goji log, gorm log

### Installation:
```
go get github.com/Innovatube/log4go
```

### Usage:
- Add the following import:
```
import log "github.com/Innovatube/log4go"
```

### Example
```
package main

import (
	"time"
)

import log "github.com/Innovatube/log4go"

func main() {
	defer log.Close()
	log.AddFilter("stdout", log.DEBUG, log.NewConsoleLogWriter())
	log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	log.Debug("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	log.Warn("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	log.Error("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	log.Trace("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
}

```