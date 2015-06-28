//go:generate mimeextract
package extramime

import (
	"log"
	"mime"
)

// Register all the mimes in types.go
func init() {
	for _, v := range mimes {
		err := mime.AddExtensionType(v.Extension, v.Type)
		if err != nil {
			log.Println(err)
		}
	}
} 
