package pages

import (
	_ "embed"
	"os"
)

//go:embed templates/upload.html
var uploadTemplate string

//go:embed templates/done.html
var doneTemplate string

// Upload returns the upload page template.
// During development, it attempts to read the file from disk to allow for quick changes.
func Upload() string {
	if data, err := os.ReadFile("pages/templates/upload.html"); err == nil {
		return string(data)
	}
	return uploadTemplate
}

// Done returns the done page template.
// During development, it attempts to read the file from disk to allow for quick changes.
func Done() string {
	if data, err := os.ReadFile("pages/templates/done.html"); err == nil {
		return string(data)
	}
	return doneTemplate
}
