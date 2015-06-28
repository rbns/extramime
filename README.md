# extramime
## wtf?
extramime adds mime type definitions from nginx or apache to the go mime 
package

## usage
### just using
importing of this package should be enough:

	import _ "github.com/rbns/extramime"

the default types used are extracted from apache (because they
seem to have more definitions..).

### regenerate mime type list
to regenerate the mime type list you can use the mimeextract tool.
install it with 

	go get github.com/rbns/extramime/mimeextract
	go install github.com/rbns/extramime/mimeextract

and run it in this directory (assuming you have $GOPATH/bin included
in your path). the tool has a few options which may
help you if you want to include a mime type list directly into your
project.

## licensing
the mime types are extracted from the nginx and sources, see
the matching LICENSE- files for the terms.

