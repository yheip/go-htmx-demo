package http

import (
	"fmt"
	"net/http"
	"sse-demo/internal/transport/http/sse"
	"time"
)

func (h *Handler) HandleSSE(w http.ResponseWriter, r *http.Request) {
	// prepare the header
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, _ := w.(http.Flusher)

	keepaliveTicker := time.NewTicker(30 * time.Second)
	defer keepaliveTicker.Stop()

	event := make(chan sse.Event)

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		i := 1
		for {
			select {
			case <-ticker.C:
				event <- sse.NewEvent("counter", fmt.Sprintf("<div>%d</div", i))
				if i%5 == 0 {
					event <- sse.NewUnamedEvent(fmt.Sprintf("<div>%d five</div", i/5))
				}
				if i%10 == 0 {
					event <- sse.NewEvent("say_hello", fmt.Sprint("helloword"))
				}
				i++
			case <-r.Context().Done():
				return
			}
		}
	}()

	for {
		select {
		case e := <-event:
			sse.WriteEvent(w, e)
			flusher.Flush()
		case <-keepaliveTicker.C:
			sse.WriteKeepAlive(w)
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
