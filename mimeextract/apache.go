package main

import (
	"net/http"
	"strings"
	"io"
	"bufio"
)

const apacheDefaultUrl = "http://svn.apache.org/repos/asf/httpd/httpd/trunk/docs/conf/mime.types"

func init() {
	sources["apache"] = apache
}

func apache() ([]MimeInfo, error) {
	if *url == "" {
		*url = apacheDefaultUrl
	}

	resp, err := http.Get(*url)
	if err != nil {
		return []MimeInfo{}, err
	}
	defer resp.Body.Close()

	mimes, err := parseApache(resp.Body)
	if err != nil {
		return []MimeInfo{}, err
	}

	return mimes, err
}

// parses an apache mime type definition in the format
//
// # comment
// mime/type	extension [extension, ...]
func parseApache(data io.Reader) ([]MimeInfo, error) {
	mimes := make([]MimeInfo, 0)

	s := bufio.NewScanner(data)

	for s.Scan() {
		text := s.Text()
		if strings.IndexRune(text, '#') == 0 {
			continue
		}

		if !strings.ContainsRune(text, '/') {
			continue	
		}

		f := strings.Fields(text)
		if len(f) < 2 {
			continue
		}

		for _, ext := range f[1:] {
			mimes = append(mimes, MimeInfo{Extension: "." + ext, Type: f[0]})
		}
	}

	return mimes, s.Err()
}

