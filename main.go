package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/mbrumlow/brumlow_io/dispatch"
)

var httpPort = flag.Int("httpPort", 8080, "HTTP bind port")

func main() {
	flag.Parse()

	d := dispatch.NewDispatch()
	d.AddNamespace("/", http.FileServer(http.Dir("webroot/")))

	http.Handle("/", d)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", "", *httpPort), nil))
}

// TODO - Setup hard coded name spaces; config, status, polymer.
// ----------------------------------------------------------------------------
//
// Modules -
// Modules will be copied (unzipped) to the disk.
// Polymer modules will be verified, and installed to the polymer directory.
//
// Component verification -
// All components must use the same polymer version.
// If not installed no version checking is required.
// If installed versions must match all modules.
// * Component and polymer version matching requirements may change, but it is
// not a goal at this time to support cross version components / polymer. *
