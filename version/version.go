package version

import "fmt"

var (
	app     = "qrcp"
	version = "dev<g8up>"
	date    = "2026/07/24"
)

// String returns a string representation of the build.
func String() string {
	return fmt.Sprintf("%s %s [date: %s]", app, version, date)
}
