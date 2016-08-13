package logger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mbrumlow/brumlow_io/statswriter"
)

type AccessLog struct {
	Time       time.Time
	Duration   time.Duration
	Method     string
	Status     int
	Length     int
	RemoteAddr string
	RequestURI string
}

func Error(err string) {
	fmt.Printf("[%v, \"ERROR\", \"error\", \"%v\"]\n", err)
}

func Access(w *statswriter.StatsWriter, r *http.Request) {
	a := AccessLog{
		Time:       w.Start,
		Duration:   time.Now().Sub(w.Start),
		Method:     r.Method,
		Status:     w.Status,
		Length:     w.Length,
		RemoteAddr: r.RemoteAddr,
		RequestURI: r.RequestURI,
	}

	b, err := json.Marshal(a)
	if err != nil {
		Error("Failed to Marshal AccesLog.")
		return
	}

	fmt.Printf("[\"%v\", \"INFO\", \"access_log\", %v]\n", time.Now(), string(b))
}
