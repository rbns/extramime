# extramime
## wtf?
extramime adds mime type definitions from nginx to the go mime package

## usage
### just using
importing of this package should be enough:

	import _ "github.com/rbns/extramime"

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
the mime types are extracted from the nginx sources, see
LICENSE-nginx for the terms.

