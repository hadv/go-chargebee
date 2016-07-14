package log4go

import (
	"fmt"
	"net/http"
	"time"
)

type gojiLogger struct {
	h http.Handler
}

func (gmLog gojiLogger) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	start := time.Now()
	Debug("Started %s '%s' from %s", req.Method, req.RequestURI, req.RemoteAddr)
	lresp := wrapWriter(resp)

	gmLog.h.ServeHTTP(lresp, req)
	lresp.maybeWriteHeader()

	latency := float64(time.Since(start)) / float64(time.Millisecond)
	Debug("Returning %d in %s", lresp.status(), fmt.Sprintf("%6.4f ms", latency))
}

func NewGojiLog() func(http.Handler) http.Handler {
	fn := func(h http.Handler) http.Handler {
		return gojiLogger{h: h}
	}
	return fn
}
