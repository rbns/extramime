package main

import (
	"net/http"
	"strings"
	"io"
	"bufio"
)

const nginxDefaultUrl = "http://trac.nginx.org/nginx/export/6345822f0abb70807f635989b6c2df7852a55bd9/nginx/conf/mime.types"

func init() {
	sources["nginx"] = nginx
}

func nginx() ([]MimeInfo, error) {
	if *url == "" {
		*url = nginxDefaultUrl
	}

	resp, err := http.Get(*url)
	if err != nil {
		return []MimeInfo{}, err
	}
	defer resp.Body.Close()

	mimes, err := parseNginx(resp.Body)
	if err != nil {
		return []MimeInfo{}, err
	}

	return mimes, err
}

// parses an nginx mime type definition in the format
//
// types {
//		text/plain		txt text foo;
//		text/html		htm html bar;
//		...
// }
//
// It returns a slice of MimeInfo
func parseNginx(data io.Reader) ([]MimeInfo, error) {
	mimes := make([]MimeInfo, 0)

	s := bufio.NewScanner(data)

	for s.Scan() {
		text := s.Text()
		text = strings.TrimRight(text, ";")

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
