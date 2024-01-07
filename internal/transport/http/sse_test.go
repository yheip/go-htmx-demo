package http

import "testing"

func TestSSE(t *testing.T) {
	s := UnnamedSSE{
		SSE: SSE{
			ID:    "abc",
			Event: "eventname",
			Data:  "yo",
			Retry: 0,
		},
	}
	t.Log(s.String())
}
