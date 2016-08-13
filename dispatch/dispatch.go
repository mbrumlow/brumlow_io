package dispatch

import (
	"net/http"

	"github.com/mbrumlow/brumlow_io/logger"
	"github.com/mbrumlow/brumlow_io/statswriter"
)

type Dispatch struct {
}

// NewDispatch returns a new Dispatch
// In favor of the service returning 404s over being down this will not error
// if it is not able to load the configuration file. A log message will be
// generated and all request will 404.
func NewDispatch(configFile string) *Dispatch {

	return nil
}

func (d *Dispatch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.route(statswriter.NewStatsWriter(w), r)
}

func (d *Dispatch) route(w *statswriter.StatsWriter, r *http.Request) {
	defer logger.Access(w, r)

	// TODO -- Find route.
	// If not in hash, check for / route.
	// If not in / route return http.NotFound

	http.NotFound(w, r)

}
