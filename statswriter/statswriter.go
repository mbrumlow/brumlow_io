package statswriter

import (
	"net/http"
	"time"
)

type StatsWriter struct {
	w      http.ResponseWriter
	Status int
	Length int
	Start  time.Time
}

func NewStatsWriter(w http.ResponseWriter) *StatsWriter {
	return &StatsWriter{w: w, Start: time.Now()}
}

func (w *StatsWriter) Header() http.Header {
	return w.w.Header()
}

func (w *StatsWriter) Write(b []byte) (int, error) {
	w.Length += len(b)
	return w.w.Write(b)
}

func (w *StatsWriter) WriteHeader(code int) {
	w.Status = code
	w.w.WriteHeader(code)
}
