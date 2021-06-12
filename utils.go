package gofnd

import (
	"mime"
	"net/http"
)

func ParseContentType(header http.Header) string {
	contentTypeRaw, ok := header["Content-Type"]
	if !ok {
		return ""
	}
	contentType, _, _ := mime.ParseMediaType(contentTypeRaw[0])

	return contentType
}
